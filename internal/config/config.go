package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	CertFilePath string `split_words:"true"`
	DataFilePath string `split_words:"true"`
	KeyFilePath  string `split_words:"true"`
	Port         int    `default:"50051" split_words:"true"`
	UsesTLS      bool   `default:"false" split_words:"true"`
}

var Server ServerConfig

func init() {
	if err := envconfig.Process("", &Server); err != nil {
		log.Fatalf("Failed to load config: %+v", err)
	}
}
