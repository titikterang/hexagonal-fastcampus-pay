package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/http"
	gorHdl "github.com/gorilla/handlers"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/money"
)

func startService(cfg *config.Config) {
	handler, consumer, err := initHandler(cfg)
	if err != nil {
		log.Fatal("failed initiate NewHandler: %v", err)
	}

	go consumer.StartConsumer()

	httpOpts := []http.ServerOption{
		http.Timeout(cfg.Http.Timeout),
		http.Address(cfg.App.Address),
		http.Filter(
			gorHdl.CORS(
				gorHdl.AllowedOrigins([]string{"*"}),
				gorHdl.AllowedHeaders([]string{"Content-Type", "Authorization"}),
				gorHdl.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}),
			),
		),
		http.Middleware(
			metadata.Server(),
			logging.Server(log.GetLogger()),
		),
		http.Logger(log.GetLogger()),
	}

	httpServer := http.NewServer(
		httpOpts...,
	)

	money.RegisterMoneyServiceHTTPServer(httpServer, handler)

	server := kratos.New(
		kratos.Name(cfg.App.Label),
		kratos.Server(
			httpServer,
		),
	)

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
