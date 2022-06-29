package grpc

import (
	"context"
	"grpcRest/grpc/pb"
)

type Math struct {
	pb.UnimplementedMathServiceServer
}

func (m *Math) Sum(ctx context.Context, in *pb.NewSumRequest) (*pb.NewSumResponse, error) {
	res := in.Sum.GetA() + in.Sum.GetB()

	return &pb.NewSumResponse{
		Result: res,
	}, nil
}
