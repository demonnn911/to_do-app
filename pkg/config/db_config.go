package config

import (
	"log"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

func NewDBConfig() *DBConfig {
	var config DBConfig
	if err := godotenv.Load(); err != nil {
		log.Fatalf("couldn't load env variables #%v", err)
	}
	filepath := os.Getenv("DB_CONFIG_PATH")
	if filepath == "" {
		log.Fatal("DB_CONFIG_PATH is empty")
	}
	configFile, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("couldn't load config file #%v", err)
	}

	if err := yaml.Unmarshal(configFile, &config); err != nil {
		log.Fatalf("couldn't parse config file into model #%v", err)
	}
	if err := env.Parse(&config); err != nil {
		log.Fatalf("couldn't fill `env` fields into model #%v", err)
	}
	return &config
}
