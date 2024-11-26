package env

import (
	"errors"
	"os"
)

const (
	dsnEnvName = "DB_DSN"
)

type DBConfig interface {
	DSN() string
}

type dbConfig struct {
	dsn string
}

func NewDBConfig() (*dbConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("pg_dsn not found")
	}

	return &dbConfig{
		dsn: dsn,
	}, nil
}

func (cfg *dbConfig) DSN() string {
	return cfg.dsn
}
