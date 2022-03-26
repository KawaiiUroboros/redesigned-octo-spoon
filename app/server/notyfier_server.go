package server

import (
	"context"
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"

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
	webhookUrl := "https://hooks.slack.com/services/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"

	attachment1 := slack.Attachment{}
	attachment1.AddField(slack.Field{Title: "Author", Value: "Чирикова Анастасия"}).AddField(slack.Field{Title: "Status", Value: "Completed"})
	attachment1.AddAction(slack.Action{Type: "button", Text: "Улететь в Армению 🛫", Url: "https://flights.example.com/book/r123456", Style: "primary"})
	payload := slack.Payload{
		Text:        "всем привет, добро пожаловать в Ереван",
		Username:    "robot",
		Channel:     "#general",
		IconEmoji:   ":monkey_face:",
		Attachments: []slack.Attachment{attachment1},
	}
	err := slack.Send(webhookUrl, "", payload)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	return &api_pb.Notyfier{NotyfierId: "mew"}, nil
}

func (s *notyfierServiceServerImpl) UpdateNotyfier(ctx context.Context, req *api_pb.UpdateNotyfierRequest) (*api_pb.Notyfier, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *notyfierServiceServerImpl) DeleteNotyfier(ctx context.Context, req *api_pb.DeleteNotyfierRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
