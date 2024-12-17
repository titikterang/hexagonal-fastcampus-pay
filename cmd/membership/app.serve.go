package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	gorHdl "github.com/gorilla/handlers"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/membership"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func startService(cfg *config.Config) {
	handler, err := initHandler(cfg)
	if err != nil {
		log.Fatal("failed initiate NewHandler: %v", err)
	}

	// init grpc
	grpcOpts := []grpc.ServerOption{
		grpc.Address(cfg.Grpc.GrpcAddress),
		grpc.Timeout(cfg.Grpc.Timeout),
		grpc.Middleware(
			metadata.Server(),
			logging.Server(log.GetLogger()),
		),
		grpc.Logger(log.GetLogger()),
	}

	// init rest api
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

	grpcServer := grpc.NewServer(
		grpcOpts...,
	)

	httpServer := http.NewServer(
		httpOpts...,
	)

	membership.RegisterMembershipServiceServer(grpcServer, handler)
	membership.RegisterMembershipServiceHTTPServer(httpServer, handler)

	server := kratos.New(
		kratos.Name(cfg.App.Label),
		kratos.Server(
			httpServer,
			grpcServer,
		),
	)

	go func() {
		if err := server.Run(); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")
	if err := server.Stop(); err != nil {
		log.Fatalf("Server forced to shutdown: %#v", err)
	}
	log.Info("Server shutdown gracefully ...")
	time.Sleep(5 * time.Millisecond) // wait for zero log to finish log writing
}
