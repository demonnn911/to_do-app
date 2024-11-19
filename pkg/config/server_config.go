package config

import (
	"log"
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type HTTPServer struct {
	Address        string        `yaml:"address"`
	MaxHeaderBytes int           `yaml:"maxheaderbytes"`
	ReadTimeout    time.Duration `yaml:"readtimeout"`
	WriteTimeout   time.Duration `yaml:"writetimeout"`
	IdleTimeout    time.Duration `yaml:"idletimeout"`
}

func NewHTTPServerConfig() *HTTPServer {
	var config HTTPServer
	if err := godotenv.Load(); err != nil {
		log.Fatalf("couldn't load env variables #%v", err)
	}
	filepath := os.Getenv("HTTPSERVER_CONFIG_PATH")
	if filepath == "" {
		log.Fatal("HTTPSERVER_CONFIG_PATH is empty")
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
