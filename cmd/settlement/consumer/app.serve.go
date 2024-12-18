package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func startService(cfg *config.Config) {
	consumer, dbConn, err := initHandler(cfg)
	if err != nil {
		log.Fatal("failed initiate NewHandler: %v", err)
	}

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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = dbConn.Disconnect(ctx); err != nil {
		panic(err)
	}

	log.Info("Server shutdown gracefully ...")
	time.Sleep(5 * time.Millisecond) // wait for zero log to finish log writing
}
