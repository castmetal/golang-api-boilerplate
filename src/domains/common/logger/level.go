package logger

import (
	"context"

	"go.uber.org/zap"
)

// Logs a message at ErrorLevel with Zap fields.
func Error(ctx context.Context, err error, msg string, fields ...zap.Field) {
	allFields := get(ctx)
	allFields = append(allFields, zap.Error(err))
	allFields = append(allFields, fields...)

	zap.L().With(allFields...).Error(msg)

	// scope := fieldsMap(allFields)
	// monitoring.Error(ctx, err, monitoring.Scope(scope))
}

// Logs a message at FatalLevel with Zap fields.
func Fatal(ctx context.Context, err error, msg string, fields ...zap.Field) {
	allFields := get(ctx)
	allFields = append(allFields, zap.Error(err))
	allFields = append(allFields, fields...)

	zap.L().With(allFields...).Fatal(msg)

	// scope := fieldsMap(allFields)
	// monitoring.Fatal(ctx, err, monitoring.Scope(scope))
}

// Logs a message at DebugLevel with Zap fields.
func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	zap.L().With(get(ctx)...).Debug(msg, fields...)
}

// Logs a message at InfoLevel with Zap fields.
func Info(ctx context.Context, msg string, fields ...zap.Field) {
	zap.L().With(get(ctx)...).Info(msg, fields...)
}

// Logs a message at WarnLevel with Zap fields.
func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	zap.L().With(get(ctx)...).Warn(msg, fields...)
}

// Keep it simple.
// Info, Error and Debug should be enough
// Warn should be used for things that are not errors(there is no action), but are unexpected.
// DPanic, Panic and Fatal have unexpected behavior
