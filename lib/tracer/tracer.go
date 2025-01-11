package tracer

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"google.golang.org/grpc/credentials"
	"strings"

	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type ClientTracer struct {
	Insecure     string
	ServiceName  string
	CollectorURL string
}

func (t *ClientTracer) InitTracer() (func(context.Context) error, *sdktrace.TracerProvider) {
	var secureOption otlptracegrpc.Option
	if strings.ToLower(t.Insecure) == "false" || t.Insecure == "0" || strings.ToLower(t.Insecure) == "f" {
		secureOption = otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	} else {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(t.CollectorURL),
		),
	)

	if err != nil {
		log.Fatal().Msgf("Failed to create exporter: %v", err)
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", t.ServiceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Fatal().Msgf("Could not set resources: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resources),
	)
	otel.SetTracerProvider(tp)
	return exporter.Shutdown, tp
}
