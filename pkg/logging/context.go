package logging

import "context"

type ctxLogger struct{}

// ContextWithLogger adds logging to context
func ContextWithLogger(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// LoggerFromContext returns logger from context
func loggerFromContext(ctx context.Context) Logger {
	if l, ok := ctx.Value(ctxLogger{}).(*logger); ok {
		return l
	}
	return NewLogger()
}
