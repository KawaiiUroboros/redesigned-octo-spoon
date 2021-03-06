//go:build wireinject
// +build wireinject

package main

import (
	"emergence/integration"
	"emergence/integration/clients"
	"github.com/google/wire"
)

// InitializeDependency initialize all structs and interfaces
func InitializeDependency() (integration.IRepository, error) {
	wire.Build(
		clients.NewSlackClient,
		clients.NewPostgresClient,
		clients.NewHospitalClient,
		integration.NewRepository,
		wire.Bind(new(clients.ISlackClient), new(clients.SlackClient)),
		wire.Bind(new(clients.IHospitalClient), new(clients.HospitalClient)),
		wire.Bind(new(clients.IPostgresClient), new(clients.PostgresClient)),
		wire.Bind(new(integration.IRepository), new(integration.Repository)),
	)
	return &integration.Repository{}, nil
}
