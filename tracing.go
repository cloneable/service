package service

import (
	"context"
	"os"

	"github.com/cloneable/service/log"
	"github.com/go-logr/zapr"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.uber.org/zap"
)

func initTracing(ctx context.Context) (context.Context, *sdktrace.TracerProvider) {
	l := log.L(ctx)

	otelLogger := l.WithOptions(zap.AddCallerSkip(2))
	otel.SetErrorHandler(errLogger{otelLogger})
	otel.SetLogger(zapr.NewLogger(otelLogger))

	f, err := os.Create("traces.txt")
	if err != nil {
		l.Sugar().Fatal(err)
	}

	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(f),
		stdouttrace.WithPrettyPrint(),
	)
	if err != nil {
		l.Sugar().Fatal(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("service"),
			semconv.ServiceVersionKey.String("v0.0.0"),
		)),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return ctx, tp
}

type errLogger struct {
	logger *zap.Logger
}

func (l errLogger) Handle(err error) {
	l.logger.Error("OpenTelemetry error", zap.Error(err))
}
