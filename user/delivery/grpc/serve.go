package grpcdelivery

import (
	"fmt"
	"net"

	"github.com/TheManhattanProject/crypt_pg_backend/user/delivery/grpc/handler"
	"github.com/TheManhattanProject/crypt_pg_backend/user/delivery/grpc/pb"
	"github.com/TheManhattanProject/crypt_pg_backend/user/domain/config"
	"github.com/TheManhattanProject/crypt_pg_backend/user/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	server *grpc.Server
}

func (s *GRPCServer) Run(
	config *config.Config,
	usecases *usecase.Usecases,
) {
	listener, err := net.Listen("tcp", config.GrpcHost)
	if err != nil {
		panic(err)
	}

	reflection.Register(s.server)

	server := handler.NewUserServer(config, usecases)
	pb.RegisterUserServiceServer(s.server, server)

	if err := s.server.Serve(listener); err != nil && err != listener.Close() {
		panic(err)
	}
}

func (s *GRPCServer) Stop() {
	fmt.Println("stopping server")
	s.server.GracefulStop()
	fmt.Println("server stopped")
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{
		server: grpc.NewServer(),
	}
}
