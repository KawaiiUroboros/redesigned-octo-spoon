package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "emergence/api"
)

// NotyfierServiceServer is a composite interface of api_pb.NotyfierServiceServer and grapiserver.Server.
type NotyfierServiceServer interface {
	api_pb.NotyfierServiceServer
	grapiserver.Server
}

// NewNotyfierServiceServer creates a new NotyfierServiceServer instance.
func NewNotyfierServiceServer() NotyfierServiceServer {
	return &notyfierServiceServerImpl{}
}

type notyfierServiceServerImpl struct {
}

func (s *notyfierServiceServerImpl) ListNotyfiers(ctx context.Context, req *api_pb.ListNotyfiersRequest) (*api_pb.ListNotyfiersResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *notyfierServiceServerImpl) GetNotyfier(ctx context.Context, req *api_pb.GetNotyfierRequest) (*api_pb.Notyfier, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *notyfierServiceServerImpl) CreateNotyfier(ctx context.Context, req *api_pb.CreateNotyfierRequest) (*api_pb.Notyfier, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *notyfierServiceServerImpl) UpdateNotyfier(ctx context.Context, req *api_pb.UpdateNotyfierRequest) (*api_pb.Notyfier, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *notyfierServiceServerImpl) DeleteNotyfier(ctx context.Context, req *api_pb.DeleteNotyfierRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
