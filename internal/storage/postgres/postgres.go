package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

const (
	DBName = "postgres"
)

const (
	lessonsTable           = "lessons"
	teachersChecklistTable = "teachers_checklist"
)

type StorageConfig struct {
	Host     string `yaml:"host" env_default:"192.168.99.100"`
	Port     string `yaml:"port" env_default:"5432"`
	Username string `yaml:"username" env_default:"postgres"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name" env_default:"postgres"`
	SSLMode  string `yaml:"ssl_mode" env_default:"disable"`
}

func New(cfg StorageConfig, logg *slog.Logger) (*sqlx.DB, error) {
	db, err := sqlx.Open(DBName, builderConnectionString(cfg))
	fmt.Println(builderConnectionString(cfg))
	if err != nil {
		logg.Error("failed to open db", slog.String("err", err.Error()))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logg.Error("failed to connect to db", slog.String("err", err.Error()))
		return nil, err
	}

	return db, nil
}

func builderConnectionString(cfg StorageConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
}
