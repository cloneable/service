package log

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

const (
	traceIDKey = "traceId"
	spanIDKey  = "spanId"
)

type loggerContextKey struct{}

var globalLogger *zap.Logger

func init() {
	var err error
	globalLogger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(globalLogger)
	zap.RedirectStdLog(globalLogger)
}

func Inject(ctx context.Context) context.Context {
	return context.WithValue(ctx, loggerContextKey{}, globalLogger)
}

func L(ctx context.Context) *zap.Logger {
	span := trace.SpanFromContext(ctx)
	if logger, ok := ctx.Value(loggerContextKey{}).(*zap.Logger); ok {
		return logger.With(TraceID(span), SpanID(span))
	}
	return globalLogger.With(TraceID(span), SpanID(span))
}

func S(ctx context.Context) *zap.SugaredLogger {
	return L(ctx).Sugar()
}

func TraceID(span trace.Span) zap.Field {
	if sc := span.SpanContext(); sc.HasTraceID() {
		return zap.String(traceIDKey, sc.TraceID().String())
	}
	return zap.Skip()
}

func SpanID(span trace.Span) zap.Field {
	if sc := span.SpanContext(); sc.HasSpanID() {
		return zap.String(spanIDKey, sc.SpanID().String())
	}
	return zap.Skip()
}
