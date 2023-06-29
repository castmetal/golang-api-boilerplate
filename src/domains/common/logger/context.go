package logger

import (
	"context"

	"go.uber.org/zap"
)

type contextKey struct{}

var (
	// logFieldsKey is used to get the []zap.Field from context
	logFieldsKey = contextKey{}
)

// WithValue returns a copy of parent with the log fields used in all subsequent logs
func WithValue(ctx context.Context, fields ...zap.Field) context.Context {
	values := get(ctx)
	values = append(values, fields...)

	return context.WithValue(ctx, logFieldsKey, values)
}

func get(ctx context.Context) []zap.Field {
	var allFields []zap.Field

	value := ctx.Value(logFieldsKey)
	if value == nil {
		value = allFields
	}
	fields := value.([]zap.Field)
	allFields = append(allFields, fields...)

	if allFields == nil {
		// 16 is big enough to avoid slice resizes in most cases
		allFields = make([]zap.Field, 0, 16)
	}

	return allFields
}
