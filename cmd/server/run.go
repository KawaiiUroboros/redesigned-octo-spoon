package main

import (
	"emergence/app/server"
	"github.com/srvc/appctx"

	"github.com/izumin5210/grapi/pkg/grapiserver"
)

func run() error {
	// Application context
	ctx := appctx.Global()

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewNotyfierServiceServer(),
		),
	)
	return s.Serve(ctx)
}
