package server

import (
	"context"
	"testing"

	api_pb "emergence/api"
)

func Test_NotyfierServiceServer_ListNotyfiers(t *testing.T) {
	svr := NewNotyfierServiceServer()

	ctx := context.Background()
	req := &api_pb.ListNotyfiersRequest{}

	resp, err := svr.ListNotyfiers(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_NotyfierServiceServer_GetNotyfier(t *testing.T) {
	svr := NewNotyfierServiceServer()

	ctx := context.Background()
	req := &api_pb.GetNotyfierRequest{}

	resp, err := svr.GetNotyfier(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_NotyfierServiceServer_CreateNotyfier(t *testing.T) {
	svr := NewNotyfierServiceServer()

	ctx := context.Background()
	req := &api_pb.CreateNotyfierRequest{}

	resp, err := svr.CreateNotyfier(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_NotyfierServiceServer_UpdateNotyfier(t *testing.T) {
	svr := NewNotyfierServiceServer()

	ctx := context.Background()
	req := &api_pb.UpdateNotyfierRequest{}

	resp, err := svr.UpdateNotyfier(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}

func Test_NotyfierServiceServer_DeleteNotyfier(t *testing.T) {
	svr := NewNotyfierServiceServer()

	ctx := context.Background()
	req := &api_pb.DeleteNotyfierRequest{}

	resp, err := svr.DeleteNotyfier(ctx, req)

	t.SkipNow()

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}
