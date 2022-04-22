package clients

import (
	"context"
	api_pb "emergence/api"
	"google.golang.org/grpc"
)

// IHospitalClient interface for hospital client
type IHospitalClient interface {
	// CreateIncident create an incident by user Id
	CreateIncident(userId string, incident string) error
}

// HospitalClient is a client for the hospital service
type HospitalClient struct {
	client *api_pb.HospitalServiceClient
}

// NewHospitalClient creates a new client for the hospital service
func NewHospitalClient() *HospitalClient {
	conn, _ := grpc.Dial(":9000", grpc.WithInsecure())
	client := api_pb.NewHospitalServiceClient(conn)
	return &HospitalClient{
		client: &client,
	}
}

// CreateIncident create an incident by user Id
func (h *HospitalClient) CreateIncident(userId string, incident string) error {
	response, _ := (*h.client).GetCityAndAddressByExternalUserId(
		context.Background(),
		&api_pb.GetCityAndAddressByExternalUserIdRequest{
			ExternalUserId: userId,
		})
	_, err := (*h.client).CreateIncident(
		context.Background(),
		&api_pb.CreateIncidentRequest{
			Title:       "инцидент по пользователю " + userId,
			Description: response.GetCity() + " " + response.GetAddress() + " - адресс пациента " + incident,
			City:        response.GetCity(),
		})

	return err
}
