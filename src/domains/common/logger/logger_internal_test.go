package logger // Using same package to avoid messing up with the stdout and stderr

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestLogger(t *testing.T) {
	tests := []struct {
		desc     string
		action   func(ctx context.Context)
		expected []observer.LoggedEntry
	}{
		{
			desc: "pure_message",
			action: func(ctx context.Context) {
				Info(ctx, "info msg", zap.Any("infoKey", "info value"))
				Error(ctx, errors.New("err error"), "error msg", zap.Any("errorKey", "error value"))
				Debug(ctx, "debug msg", zap.Any("debugKey", "debug value"))
			},
			expected: []observer.LoggedEntry{
				{
					Entry: zapcore.Entry{Level: zap.InfoLevel, Message: "info msg"},
					Context: []zap.Field{
						zap.Any("infoKey", "info value")},
				},
				{
					Entry: zapcore.Entry{Level: zap.ErrorLevel, Message: "error msg"},
					Context: []zap.Field{
						zap.Error(errors.New("err error")),
						zap.Any("errorKey", "error value"),
					},
				},
				{
					Entry: zapcore.Entry{Level: zap.DebugLevel, Message: "debug msg"},
					Context: []zap.Field{
						zap.Any("debugKey", "debug value"),
					},
				},
			},
		},
		{
			desc: "message_with_context",
			action: func(ctx context.Context) {
				ctx = WithValue(ctx, zap.String("ctxKey", "ctx value"))
				Info(ctx, "info msg", zap.Any("infoKey", "info value"))
				Error(ctx, errors.New("err error"), "error msg", zap.Any("errorKey", "error value"))
				Debug(ctx, "debug msg", zap.Any("debugKey", "debug value"))
			},
			expected: []observer.LoggedEntry{
				{
					Entry: zapcore.Entry{Level: zap.InfoLevel, Message: "info msg"},
					Context: []zap.Field{
						zap.String("ctxKey", "ctx value"),
						zap.Any("infoKey", "info value"),
					},
				},
				{
					Entry: zapcore.Entry{Level: zap.ErrorLevel, Message: "error msg"},
					Context: []zap.Field{
						zap.String("ctxKey", "ctx value"),
						zap.Error(errors.New("err error")),
						zap.Any("errorKey", "error value"),
					},
				},
				{
					Entry: zapcore.Entry{Level: zap.DebugLevel, Message: "debug msg"},
					Context: []zap.Field{
						zap.String("ctxKey", "ctx value"),
						zap.Any("debugKey", "debug value"),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			ctx := context.Background()

			logs := setupTestLogger()
			defer Flush()

			tt.action(ctx)

			actual := logs.AllUntimed()
			require.Equal(t, tt.expected, actual)
		})
	}
}

func setupTestLogger() *observer.ObservedLogs {
	zCore, observer := observer.New(zap.DebugLevel)
	logger := zap.New(zCore, []zap.Option{}...)
	zap.ReplaceGlobals(logger)
	return observer
}
