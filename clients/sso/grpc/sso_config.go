package grpc

import (
	"log"
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type SSOConfig struct {
	Address      string        `yaml:"address"`
	Timeout      time.Duration `yaml:"timeout"`
	RetriesCount uint          `yaml:"retriescount"`
}

func NewSSOConfig() *SSOConfig {
	var config SSOConfig
	if err := godotenv.Load(); err != nil {
		log.Fatalf("couldn't load env variables #%v", err)
	}
	filepath := os.Getenv("SSO_CONFIG_PATH")
	if filepath == "" {
		log.Fatal("SSO_CONFIG_PATH is empty")
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
