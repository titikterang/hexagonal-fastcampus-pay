package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/http"
	gorHdl "github.com/gorilla/handlers"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/adapter/handler"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/payment"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func startService(cfg *config.Config) {
	consumer, producer, dbConn, err := initHandler(cfg)
	if err != nil {
		log.Fatal("failed initiate NewHandler: %v", err)
	}

	hdl, err := handler.NewHandler(cfg, nil)
	if err != nil {
		log.Fatal("failed initiate NewHandler: %v", err)
	}

	httpOpts := []http.ServerOption{
		http.Timeout(cfg.Http.Timeout),
		http.Address(cfg.App.HealthAddress),
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
	payment.RegisterPaymentServiceHTTPServer(httpServer, hdl)
	server := kratos.New(
		kratos.Name("consumer-"+cfg.App.Label),
		kratos.Server(
			httpServer,
		),
	)
	go func() {
		if err := server.Run(); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	go consumer.StartConsumer()
	log.Info("Starting consumer...")
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")
	consumer.CloseClient()
	producer.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = dbConn.Disconnect(ctx); err != nil {
		panic(err)
	}
	if err := server.Stop(); err != nil {
		log.Fatalf("Server forced to shutdown: %#v", err)
	}
	log.Info("Server shutdown gracefully ...")
	time.Sleep(5 * time.Millisecond) // wait for zero log to finish log writing
}
