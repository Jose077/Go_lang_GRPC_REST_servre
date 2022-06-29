package main

import (
	"log"
	"os"

	grpc "grpcRest/grpc/server"
	"grpcRest/rest"

	"go.uber.org/fx"
)

func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")
	return logger
}

func main() {
	app := fx.New(

		fx.Provide(
			NewLogger,
			rest.NewHandler,
			rest.NewMux,
			grpc.InitGRPCserver,
		),

		fx.Invoke(
			rest.Register,
			grpc.RunGRPCServer,
		),
	)

	app.Run()

}
