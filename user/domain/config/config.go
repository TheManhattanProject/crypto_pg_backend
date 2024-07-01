package config

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	GrpcHost                    string        `env:"GRPC_HOST"                        envDefault:"0.0.0.0:6005"`
	PSQLDatabaseName            string        `env:"POSTGRES_DB"                      envDefault:"crypto_pg_backend"`
	PSQLDatabaseHost            string        `env:"POSTGRES_HOST"                    envDefault:"postgres"` // cause we use docker
	PSQLDatabaseUser            string        `env:"POSTGRES_USER"                    envDefault:"postgres"`
	PSQLDatabasePassword        string        `env:"POSTGRES_PASSWORD"                envDefault:"postgres"`
	PSQLDatabasePort            int           `env:"POSTGRES_PORT"                    envDefault:"5432"`
	PSQLDatabaseSSLMode         string        `env:"POSTGRES_SSL_MOD"                 envDefault:"disable"`
	PSQLIdleConnectionCount     int           `env:"POSTGRES_IDLE_CONNECTIONS_COUNT"  envDefault:"10"`
	PSQLMaxOpenConnectionsCount int           `env:"POSTGRES_MAX_OPEN_CONNECTIONS"    envDefault:"10"`
	PSQLConnectionMaxLifeTime   time.Duration `env:"POSTGRES_CONNECTION_MAX_LIFETIME"                                envDefaut:"3600s"`
	RedisHost                   string        `env:"REDIS_HOST"                       envDefault:"redis"` // cause docker is being used
	RedisPort                   int           `env:"REDIS_PORT"                       envDefault:"6379"`
	RedisDB                     int           `env:"REDIS_DB"                         envDefault:"0"`
	HashRounds                  int           `env:"HASH_ROUNDS"                      envDefault:"10"`
	OTPExpiryTime               time.Duration `env:"OTP_EXPIRY_TIME"                  envDefault:"120s"`
}

func LoadConfig() *Config {
	// env to be loaded using docker anyways so we don't call godotenv or viper or smth
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
