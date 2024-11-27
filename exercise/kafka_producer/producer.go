package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
	"os"
	"os/signal"
	"strings"
)

var (
	seedBrokers = flag.String("brokers", "localhost:9094", "comma delimited list of seed brokers")
	topic       = flag.String("topic", "fastcampus_message", "topic to consume from")
	style       = flag.String("commit-style", "autocommit", "commit style (which consume & commit is chosen); autocommit|records|uncommitted")
	group       = flag.String("group", "payment-service", "group to consume within")
	logger      = flag.Bool("logger", false, "if true, enable an info level logger")
)

func die(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}

func main() {
	flag.Parse()

	opts := []kgo.Opt{
		kgo.SeedBrokers(strings.Split(*seedBrokers, ",")...),
		kgo.ConsumerGroup(*group),
		kgo.ConsumeTopics(*topic),
	}

	if *logger {
		opts = append(opts, kgo.WithLogger(kgo.BasicLogger(os.Stderr, kgo.LogLevelInfo, nil)))
	}

	cl, err := kgo.NewClient(opts...)
	if err != nil {
		die("unable to create client: %v", err)
	}

	cl.ProduceSync(context.TODO(), &kgo.Record{
		Key:   nil,
		Value: []byte("publish transfer payment"),
		Topic: *topic,
	})

	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, os.Interrupt)

	<-sigs
	fmt.Println("received interrupt signal; closing client")
	done := make(chan struct{})
	go func() {
		defer close(done)
		cl.Close()
	}()

	select {
	case <-sigs:
		fmt.Println("received second interrupt signal; quitting without waiting for graceful close")
	case <-done:
	}
}
