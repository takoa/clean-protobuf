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

type DatabaseConfig struct {
	Host     string `required:"true" split_words:"true"`
	Name     string `required:"true" split_words:"true"`
	Password string `required:"true" split_words:"true"`
	Port     int    `default:"5432" split_words:"true"`
	SSLMode  string `default:"disable" split_words:"true"`
	TimeZone string `default:"Etc/UTC" split_words:"true"`
	User     string `required:"true" split_words:"true"`
}

var (
	Server ServerConfig
	DB     DatabaseConfig
)

func init() {
	if err := envconfig.Process("server", &Server); err != nil {
		log.Fatalf("Failed to load server config: %+v", err)
	}
	if err := envconfig.Process("db", &DB); err != nil {
		log.Fatalf("Failed to load DB config: %+v", err)
	}
}
