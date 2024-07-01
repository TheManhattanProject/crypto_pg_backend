package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	grpcdelivery "github.com/TheManhattanProject/crypt_pg_backend/user/delivery/grpc"
	"github.com/TheManhattanProject/crypt_pg_backend/user/domain/config"
	"github.com/TheManhattanProject/crypt_pg_backend/user/repository"
	userrepository "github.com/TheManhattanProject/crypt_pg_backend/user/repository/user"
	"github.com/TheManhattanProject/crypt_pg_backend/user/usecase"
	"github.com/TheManhattanProject/crypto_pg_backend/common/database"
)

func main() {
	config := config.LoadConfig()

	postgresConfig := &database.PostgresConfig{
		Host:                    config.PSQLDatabaseHost,
		Port:                    config.PSQLDatabasePort,
		User:                    config.PSQLDatabaseUser,
		Password:                config.PSQLDatabasePassword,
		DBName:                  config.PSQLDatabaseName,
		SSLMode:                 config.PSQLDatabaseSSLMode,
		IdleConnectionCount:     config.PSQLIdleConnectionCount,
		MaxOpenConnectionsCount: config.PSQLMaxOpenConnectionsCount,
		ConnectionMaxLifeTime:   config.PSQLConnectionMaxLifeTime,
	}

	redisConfig := &database.RedisConfig{
		Host:     config.RedisHost,
		Port:     config.RedisPort,
		DB:       config.RedisDB,
		Password: config.RedisPassword,
	}

	db := database.CreateNewPostgresClient(postgresConfig)
	cache := database.NewRedisClient(context.Background(), *redisConfig)

	repositories := repository.Repositories{
		User: userrepository.NewUserRepositoryImpl(config, db),
	}

	caches := repository.Caches{
		User: userrepository.NewUserCache(config, cache),
	}

	usecases := usecase.InitUsecases(config, repositories, caches)

	cancel := make(chan os.Signal, 1)
	signal.Notify(cancel, os.Interrupt)

	server := grpcdelivery.NewGRPCServer()
	go server.Run(config, usecases)
	fmt.Printf("server running at %v", config.GrpcHost)

	<-cancel
	server.Stop()
}
