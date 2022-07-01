package network

import (
	"context"

	"github.com/mohammadne/university/cryptography/internal/models"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	config *Config
	logger zap.Logger

	server *grpc.Server
	models.UnimplementedGreeterServer
}

func NewServer(cfg *Config, log zap.Logger) *server {
	s := &server{config: cfg, logger: log}

	s.server = grpc.NewServer()
	models.RegisterGreeterServer(s.server, s)

	return s
}

func (s server) SayHello(context.Context, *models.HelloRequest) (*models.HelloReply, error) {
	return &models.HelloReply{Message: "hello"}, nil
}

func (s server) SayHelloAgain(context.Context, *models.HelloRequest) (*models.HelloReply, error) {
	return &models.HelloReply{Message: "hello again"}, nil
}
