package handler

import (
	"github.com/TheManhattanProject/crypt_pg_backend/user/delivery/grpc/pb"
	"github.com/TheManhattanProject/crypt_pg_backend/user/domain/config"
	"github.com/TheManhattanProject/crypt_pg_backend/user/usecase"
)

type server struct {
	pb.UnimplementedUserServiceServer

	config   *config.Config
	usecases *usecase.Usecases
}

func NewUserServer(config *config.Config, usecases *usecase.Usecases) *server {
	return &server{
		config:   config,
		usecases: usecases,
	}
}
