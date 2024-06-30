package database

import (
	"database/sql"
	"fmt"
	"time"
)

type PostgresConfig struct {
	// Connection Requirements
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string

	// Other postgres related options
	IdleConnectionCount     int
	MaxOpenConnectionsCount int
	ConnectionMaxLifeTime   time.Duration
}

func CreateNewPostgresClient(cfg *PostgresConfig) *sql.DB { // TODO: Add logger
	connectionString := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmod=%v",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(cfg.IdleConnectionCount)
	db.SetConnMaxIdleTime(cfg.ConnectionMaxLifeTime)
	db.SetMaxOpenConns(cfg.MaxOpenConnectionsCount)

	return db
}
