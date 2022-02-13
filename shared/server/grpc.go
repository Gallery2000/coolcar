package server

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type GRPCConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFun       func(server *grpc.Server)
	Logger            *zap.Logger
}

func RunGRPCServer(c *GRPCConfig) error {
	nameFiled := zap.String("name", c.Name)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		c.Logger.Fatal("cannot listen", nameFiled, zap.Error(err))
	}

	var opts []grpc.ServerOption
	if c.AuthPublicKeyFile != "" {
		in,err :=
	}
}
