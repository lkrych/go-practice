package vault

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"
)

type grpcServer struct {
	hash     grpctransport.Handler
	validate grpctransport.Handler
}

//Hash will serve requests by first decoding the requests, calling the appropriate
//endpoint function, getting the response, and encoding and sendint it to the client. This is the MAGIC
//of of the serveGRPC method of Go-kit
func (s *grpcServer) Hash(ctx context.Context, r *pb.HashRequest) (*pb.HashResponse, error) {
	_, resp, err := s.hash.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.HashResponse), nil
}

//Validate will serve requests by first decoding the requests, calling the appropriate
//endpoint function, getting the response, and encoding and sendint it to the client. This is the MAGIC
//of of the serveGRPC method of Go-kit
func (s *grpcServer) Validate(ctx context.Context, r *pb.ValidateRequest) (*ValidateResponse, error) {
	_, resp, err := s.validate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return nil, rep.(*pb.ValidateResponse), nil
}
