package gorm

import (
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	dsn        string
	outputPath string
	excluded   cli.StringSlice
)

var GenerateGORMCmd = &cli.Command{
	Name: "gorm",
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
			Name:        "out",
			Value:       "./model/generated",
			Usage:       "The output path for the generated GORM models.",
			Destination: &outputPath,
		},
	},
	Action: func(cCtx *cli.Context) error {
		return run()
	},
}

func run() error {
	excludedTables := map[string]struct{}{}
	for _, excludedTable := range excluded.Value() {
		excludedTables[excludedTable] = struct{}{}
	}

	if err := generateGORMCode(
		dsn,
		excludedTables,
		outputPath,
	); err != nil {
		return err
	}

	return nil
}

func generateGORMCode(
	dsn string,
	excludedTables map[string]struct{},
	modelOutputPath string,
) error {
	g := gen.NewGenerator(gen.Config{
		ModelPkgPath:  modelOutputPath,
		FieldNullable: true,
	})

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return xerrors.Errorf("failed to connect to the DB: %w", err)
	}
	g.UseDB(db)

	tables, err := db.Migrator().GetTables()
	if err != nil {
		return xerrors.Errorf("failed to list all tables: %w", err)
	}

	for _, table := range tables {
		if _, ok := excludedTables[table]; ok {
			continue
		}
		g.GenerateModel(table)
	}

	g.Execute()

	return nil
}
