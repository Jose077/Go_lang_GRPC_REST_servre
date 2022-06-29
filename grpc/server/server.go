package grpc

import (
	"context"
	"log"
	"net"

	"grpcRest/grpc/pb"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func InitGRPCserver(lc fx.Lifecycle) (*grpc.Server, net.Listener) {

	grpcServer := grpc.NewServer()

	listener, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}

	return grpcServer, listener
}

func RunGRPCServer(lc fx.Lifecycle, grpcServer *grpc.Server, listener net.Listener) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// grpc adicao
			pb.RegisterMathServiceServer(grpcServer, &Math{})

			reflection.Register(grpcServer)

			go grpcServer.Serve(listener)

			return nil

		},

		OnStop: func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
	})

	return nil
}
