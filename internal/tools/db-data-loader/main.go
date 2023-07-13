package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/takoa/clean-protobuf/internal/entity/model"
	"github.com/takoa/clean-protobuf/internal/infrastructure/repository"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dsn          string
	dataFilePath string
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "dsn",
				Usage:       "DSN for gorm.io/gen. example: \"host=localhost user=postgres password=password dbname=clean-protobuf port=5432 sslmode=disable TimeZone=Etc/UTC\"",
				Required:    true,
				Destination: &dsn,
			},
			&cli.StringFlag{
				Name:        "data-file-path",
				Aliases:     []string{"d"},
				Required:    true,
				Usage:       "data file path",
				Destination: &dataFilePath,
			},
		},
		Action: func(cCtx *cli.Context) error {
			return initializeDatabase()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func initializeDatabase() error {
	data, err := os.ReadFile(dataFilePath)
	if err != nil {
		return xerrors.Errorf("failed to load the provided feature data file: %w", err)
	}

	var features []*model.Feature
	if err := json.Unmarshal(data, &features); err != nil {
		return xerrors.Errorf("failed to load default features: %w", err)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return xerrors.Errorf("failed to connect to the DB: %w", err)
	}
	ctx := context.Background()

	repo := repository.NewFeatures(db)

	rowsAffected, err := repo.Save(ctx, features...)
	if err != nil {
		return xerrors.Errorf("failed to insert features: %w", err)
	}
	log.Printf("%d rows inserted.", rowsAffected)

	return nil
}
