package postgres

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	DBName          string        `mapstructure:"dbname"`
	SSLMode         string        `mapstructure:"ssl_mode"`
	ConnectTimeout  time.Duration `mapstructure:"connect_timeout"`
	ApplicationName string        `mapstructure:"application_name"`

	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
}

type DB struct {
	Sqlx *sqlx.DB
	mu   sync.RWMutex
}

func Connect(cfg *Config) (*DB, error) {
	if cfg == nil {
		return nil, errors.New("postgres config is nil")
	}

	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s connect_timeout=%d",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DBName,
		cfg.Password,
		cfg.SSLMode,
		cfg.ConnectTimeout,
	)

	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime * time.Minute)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime * time.Minute)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	log.Println("Connected to postgres database")

	return &DB{Sqlx: db}, nil
}
