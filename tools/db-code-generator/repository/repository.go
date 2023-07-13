package repository

import (
	"bytes"
	"errors"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/takoa/clean-protobuf/tools/db-code-generator/inflection"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dsn                      string
	excluded                 cli.StringSlice
	modelPkg                 string
	interfaceOutputPath      string
	implementationOutputPath string
	templatePath             string
)

var GenerateRepositoryCmd = &cli.Command{
	Name: "repository",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "dsn",
			Usage:       "DSN for gorm.io/gen. Example: \"host=localhost user=postgres password=password dbname=clean-protobuf port=5432 sslmode=disable TimeZone=Etc/UTC\"",
			Required:    true,
			Destination: &dsn,
		},
		&cli.StringSliceFlag{
			Name:        "excluded-table",
			Aliases:     []string{"e"},
			Usage:       "Tables to exclude",
			Destination: &excluded,
		},
		&cli.StringFlag{
			Name:        "model-pkg",
			Required:    true,
			Usage:       "The model package",
			Destination: &modelPkg,
		},
		&cli.StringFlag{
			Name:        "interface-out",
			Value:       "./repository",
			Usage:       "The output path for the generated interfaces.",
			Destination: &interfaceOutputPath,
		},
		&cli.StringFlag{
			Name:        "implementation-out",
			Value:       "./infrastructure",
			Usage:       "The output path for the generated implementation.",
			Destination: &implementationOutputPath,
		},
		&cli.StringFlag{
			Name:        "template-path",
			Required:    true,
			Usage:       "The template file",
			Destination: &templatePath,
		},
	},
	Action: func(cCtx *cli.Context) error {
		return run()
	},
}

type packageInfo struct {
	FullName  string
	ShortName string
}

type modelInfo struct {
	Name    string
	Package packageInfo
}

func run() error {
	excludedTables := map[string]struct{}{}
	for _, excludedTable := range excluded.Value() {
		excludedTables[excludedTable] = struct{}{}
	}
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return xerrors.Errorf("failed to parse template %s: %w", templatePath, err)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return xerrors.Errorf("failed to connect to the DB: %w", err)
	}
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return xerrors.Errorf("failed to list all tables: %w", err)
	}

	m := &modelInfo{
		Package: packageInfo{
			FullName:  modelPkg,
			ShortName: filepath.Base(modelPkg),
		},
	}
	params := map[string]interface{}{
		"InterfacePackageName":      filepath.Base(interfaceOutputPath),
		"ImplementationPackageName": filepath.Base(implementationOutputPath),
		"Model":                     m,
	}
	for _, table := range tables {
		if _, ok := excludedTables[table]; ok {
			continue
		}

		m.Name = strcase.ToCamel(inflection.ToSingularSnake(table))
		pluralSnakeName := inflection.ToPluralSnake(table)
		fileName := pluralSnakeName + ".go"

		interfacePath := filepath.Join(interfaceOutputPath, fileName)
		if _, err := os.Stat(interfacePath); err != nil && errors.Is(err, os.ErrNotExist) {
			params["InterfaceName"] = strcase.ToCamel(pluralSnakeName)
			if err := generate("interface", interfacePath, tmpl, params); err != nil {
				return xerrors.Errorf("error on generating %s: %w", interfacePath, err)
			}
		} else {
			log.Printf("Skipping; file already exists or an error occured: %s", interfacePath)
		}

		implementationPath := filepath.Join(implementationOutputPath, fileName)
		if _, err := os.Stat(implementationPath); err != nil && errors.Is(err, os.ErrNotExist) {
			params["ImplementationName"] = strcase.ToCamel(pluralSnakeName)

			if err := generate("implementation", implementationPath, tmpl, params); err != nil {
				return xerrors.Errorf("error on generating %s: %w", implementationPath, err)
			}
		} else {
			log.Printf("Skipping; file already exists or an error occured: %s", implementationPath)
		}
	}

	return nil
}

func generate(
	templateName string,
	filePath string,
	tmpl *template.Template,
	templateParams map[string]interface{},
) error {
	buffer := &bytes.Buffer{}

	if err := tmpl.ExecuteTemplate(buffer, templateName, templateParams); err != nil {
		return xerrors.Errorf("failed to apply template: %w", err)
	}

	formatted, err := format.Source(buffer.Bytes())
	if err != nil {
		return xerrors.Errorf("failed to format: %w", err)
	}

	if err := os.WriteFile(filePath, formatted, 0644); err != nil {
		return xerrors.Errorf("failed to write %s: %w", filePath, err)
	}

	log.Printf("Generated: %s", filePath)

	return nil
}
