package main

import (
	"emergence/app/server"
	"emergence/cmd/migrations"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/srvc/appctx"
)

func run() error {
	//auto-migrations uo
	errorFromMigrations := migrations.Up()

	if errorFromMigrations != nil {
		panic("Migrations are not successful")
	}
	// Application context
	ctx := appctx.Global()

	s := grapiserver.New(
		grapiserver.WithGrpcAddr("tcp", ":9000"),
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewNotyfierServiceServer(),
		),
	)
	return s.Serve(ctx)
}
