package model

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
	dsn          string
	excluded     cli.StringSlice
	gormModelPkg string
	outputPath   string
	templatePath string
)

var GenerateModelCmd = &cli.Command{
	Name: "model",
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
			Name:        "gorm-model-pkg",
			Required:    true,
			Usage:       "The GORM model package",
			Destination: &gormModelPkg,
		},
		&cli.StringFlag{
			Name:        "out",
			Value:       "./model",
			Usage:       "The output path for the generated models.",
			Destination: &outputPath,
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
	params := map[string]interface{}{
		"GORMModelPackage": &packageInfo{
			FullName:  gormModelPkg,
			ShortName: filepath.Base(gormModelPkg),
		},
		"PackageName": filepath.Base(outputPath),
	}

	for _, table := range tables {
		if _, ok := excludedTables[table]; ok {
			continue
		}

		snakeName := inflection.ToSingularSnake(table)
		filePath := filepath.Join(outputPath, snakeName+".go")
		if _, err := os.Stat(filePath); err != nil && errors.Is(err, os.ErrNotExist) {
			params["StructName"] = strcase.ToCamel(snakeName)
			if err := generate(filePath, tmpl, params); err != nil {
				return xerrors.Errorf("error on generating %s: %w", filePath, err)
			}
		}
	}

	return nil
}

func generate(
	filePath string,
	tmpl *template.Template,
	templateParams map[string]interface{},
) error {
	buffer := &bytes.Buffer{}

	if err := tmpl.Execute(buffer, templateParams); err != nil {
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
