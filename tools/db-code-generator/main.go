package main

import (
	"log"
	"os"

	"github.com/takoa/clean-protobuf/tools/db-code-generator/gorm"
	"github.com/takoa/clean-protobuf/tools/db-code-generator/model"
	"github.com/takoa/clean-protobuf/tools/db-code-generator/repository"
	"github.com/urfave/cli/v2"
)

var rootCmd = &cli.App{
	Commands: []*cli.Command{
		gorm.GenerateGORMCmd,
		model.GenerateModelCmd,
		repository.GenerateRepositoryCmd,
	},
}

func main() {
	if err := rootCmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
