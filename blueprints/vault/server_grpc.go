package vault

import (
	"context"
	"vault/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

//In this file we are creating an interface for the service endpoints
//Our methods will serve requests by first decoding them, calling the appropriate endpoint, getting the response, encoding it and sending it back tot he client!
//Much of this work was done when we used the protobuf compiler
type grpcServer struct {
	hash     grpctransport.Handler
	validate grpctransport.Handler
}

func (s *grpcServer) Hash(ctx context.Context, r *pb.HashRequest) (*pb.HashResponse, error) {
	_, resp, err := s.hash.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.HashResponse), nil
}

func (s *grpcServer) Validate(ctx context.Context, r *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	_, resp, err := s.validate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ValidateResponse), nil
}
