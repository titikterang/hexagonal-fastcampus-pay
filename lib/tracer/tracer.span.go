package tracer

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func StartSpan(ctx context.Context, spanName string) (context.Context, trace.Span) {
	if ctx == nil {
		ctx = context.Background()
	}
	if spanName == "" {
		spanName = "spanName"
	}
	tr := otel.Tracer("/github.com/titikterang/hexagonal-fastcampus-pay/")
	return tr.Start(ctx, spanName)
}

func StartInitSpan(ctx context.Context, spanName string) (context.Context, trace.Span) {
	if ctx == nil {
		ctx = context.Background()
	}
	if spanName == "" {
		spanName = "spanName"
	}
	tr := otel.Tracer("init")
	return tr.Start(ctx, spanName)
}

func StartUseCaseSpan(ctx context.Context, spanName string) (context.Context, trace.Span) {
	if ctx == nil {
		ctx = context.Background()
	}
	if spanName == "" {
		spanName = "spanName"
	}
	tr := otel.Tracer("usecase")
	return tr.Start(ctx, spanName)
}

func StartTransportSpan(ctx context.Context, spanName string) (context.Context, trace.Span) {
	if ctx == nil {
		ctx = context.Background()
	}
	if spanName == "" {
		spanName = "spanName"
	}
	tr := otel.Tracer("transport")
	return tr.Start(ctx, spanName)
}

func StartRepositorySpan(ctx context.Context, spanName string) (context.Context, trace.Span) {
	if ctx == nil {
		ctx = context.Background()
	}
	if spanName == "" {
		spanName = "spanName"
	}
	tr := otel.Tracer("repository")
	return tr.Start(ctx, spanName)
}
