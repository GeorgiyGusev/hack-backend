package postgres

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type Config struct {
	DbName string `env:"DB_NAME" json:"db_name" envDefault:"postgres"`
	DbUser string `env:"DB_USER" json:"db_user" envDefault:"postgres"`
	DbPass string `env:"DB_PASS" json:"db_pass" envDefault:"postgres"`
	DbHost string `env:"DB_HOST" json:"db_host" envDefault:"localhost"`
	DbPort string `env:"DB_PORT" json:"db_port" envDefault:"5432"`
}

func LoadConfig() (*Config, error) {
	var cfg struct {
		Config Config `json:"postgres" env-prefix:"POSTGRES_"`
	}
	err := cleanenv.ReadConfig("config.json", &cfg)
	if err != nil {
		err := cleanenv.ReadEnv(&cfg)
		if err != nil {
			return nil, err
		}
	}
	return &cfg.Config, nil
}

func NewPostgressConn(cfg *Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPass, cfg.DbName,
	)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func TestPostgreConn(db *sqlx.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}
