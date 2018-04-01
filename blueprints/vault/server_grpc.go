package vault

import (
	"context"
	"vault/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	xnetcontext "golang.org/x/net/context"
)

//In this file we are creating an interface for the service endpoints
//Our methods will serve requests by first decoding them, calling the appropriate endpoint, getting the response, encoding it and sending it back tot he client!
//Much of this work was done when we used the protobuf compiler
type grpcServer struct {
	hash     grpctransport.Handler
	validate grpctransport.Handler
}

func NewGRPCServer(ctx context.Context, endpoints Endpoints) pb.VaultServer {
	return &grpcServer{
		hash: grpctransport.NewServer(
			endpoints.HashEndpoint,
			DecodeGRPCHashRequest,
			EncodeGRPCHashResponse,
		),
		validate: grpctransport.NewServer(
			endpoints.ValidateEndpoint,
			DecodeGRPCValidateRequest,
			EncodeGRPCValidateResponse,
		),
	}
}

func (s *grpcServer) Hash(ctx xnetcontext.Context, r *pb.HashRequest) (*pb.HashResponse, error) {
	_, resp, err := s.hash.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.HashResponse), nil
}

func (s *grpcServer) Validate(ctx xnetcontext.Context, r *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	_, resp, err := s.validate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ValidateResponse), nil
}

//we are using the request and response objects from the pb package, but our own endpoints use the structs we defined in service.go
//we need to translate between the two.

//EncodeGRPCHashRequest is used to translate our own hashRequest type into a protocol buffer type that can be used to communicated with the client.
func EncodeGRPCHashRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(hashRequest)
	return &pb.HashRequest{Password: req.Password}, nil
}

//DecodeGRPCHashRequest is used to translate our protocol buffer type into our hashRequest type that can be used to communicated with the client.
func DecodeGRPCHashRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.HashRequest)
	return hashRequest{Password: req.Password}, nil
}

//EncodeGRPCHashResponse is used to translate our own hashResponse type into a protocol buffer type that can be used to communicated with the client.
func EncodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(hashResponse)
	return &pb.HashResponse{Hash: res.Hash, Err: res.Err}, nil
}

//DecodeGRPCHashResponse is used to translate our protocol buffer type into our hashResponse type that can be used to communicated with the client.
func DecodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.HashResponse)
	return hashResponse{Hash: res.Hash, Err: res.Err}, nil
}

//EncodeGRPCValidateRequest is used to translate our own validateRequest type into a protocol buffer type that can be used to communicated with the client.
func EncodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(validateRequest)
	return &pb.ValidateRequest{Password: req.Password, Hash: req.Hash}, nil
}

//DecodeGRPCValidateRequest is used to translate our protocol buffer type into our validateRequest type that can be used to communicated with the client.
func DecodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ValidateRequest)
	return validateRequest{Password: req.Password, Hash: req.Hash}, nil
}

//EncodeGRPCValidateResponse is used to translate our own validateResponse type into a protocol buffer type that can be used to communicated with the client.
func EncodeGRPCValidateResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(validateResponse)
	return &pb.ValidateResponse{Valid: res.Valid}, nil
}

//DecodeGRPCValidateResponse is used to translate our protocol buffer type into our validateResponse type that can be used to communicated with the client.
func DecodeGRPCValidateResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.ValidateResponse)
	return validateResponse{Valid: res.Valid}, nil
}
