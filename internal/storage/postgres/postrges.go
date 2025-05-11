package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

const (
	DBName = "postgres"
)

// It an old code, but I don't want to delete it
// I'm sorry for that
const (
	usersTable      = "users"
	teamsTable      = "teams"
	usersTeamsTable = "users_teams"
	nhlTeamsTable   = "nhl_teams"
	nflTeamsTable   = "nfl_teams"
	nhlRosterTable  = "nhl_roster"
)

type StorageConfig struct {
	Host     string `yaml:"host" env-default:"192.168.99.101"`
	Port     string `yaml:"port" env-default:"5432"`
	Username string `yaml:"username" env-default:"postgres"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname" env-default:"postgres"`
	SSLMode  string `yaml:"sslmode" env-default:"disable"`
}

func New(cfg StorageConfig, logg *slog.Logger) (*sqlx.DB, error) {
	db, err := sqlx.Open(DBName, builderConnectionString(cfg))
	fmt.Println(builderConnectionString(cfg)) // TODO: Delete
	if err != nil {
		logg.Error("failed to open db connection", slog.Any("error", err), slog.String("connection_string", builderConnectionString(cfg)))
		return nil, fmt.Errorf("sqlx.Open - %w", err)
	}

	err = db.Ping()
	if err != nil {
		logg.Error("failed to ping db", slog.Any("error", err), slog.String("connection_string", builderConnectionString(cfg)))
		return nil, fmt.Errorf("db.Ping - %w", err)
	}

	return db, nil
}

func builderConnectionString(cfg StorageConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
}
