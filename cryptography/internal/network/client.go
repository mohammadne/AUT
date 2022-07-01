package network

import (
	"context"
	"fmt"

	"github.com/mohammadne/university/cryptography/internal/models"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type client struct {
	config *Config
	logger zap.Logger

	c models.GreeterClient
}

func NewClient(cfg *Config, log zap.Logger) *client {
	c := &client{config: cfg, logger: log}

	Address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	connection, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c.c = models.NewGreeterClient(connection)

	return c, nil
}

func (c *client) someFunc() {
	req := models.HelloRequest{}
	res, err := c.c.SayHello(context.Background(), &req)
	if err != nil {
		c.logger.Error("error", zap.Error(err))
	}
	_ = res.Message
}
