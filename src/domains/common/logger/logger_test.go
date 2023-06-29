package logger_test

import (
	"context"
	"testing"

	"github.com/castmetal/golang-api-boilerplate/src/domains/common/logger"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func ExampleWithValue() {
	err := logger.Setup("example") // Change env to "dev" (dev friendly) or "prod" (production suitable)
	if err != nil {
		return
	}
	defer logger.Flush()

	ctx := context.Background()
	ctx = logger.WithValue(ctx, zap.String("foo", "bar"))
	logger.Debug(ctx, "my message")
	// Output: {"level":"debug","msg":"my message","foo":"bar"}
}

func TestSetup(t *testing.T) {
	tests := []struct {
		env string
	}{
		{env: "prod"},
		{env: "example"},
		{env: "nop"},
		{env: "dev"},
	}
	for _, test := range tests {
		t.Run(test.env, func(t *testing.T) {
			ctx := context.Background()
			err := logger.Setup(test.env)
			require.NoError(t, err)
			logger.Debug(ctx, "my message")
			logger.Flush()
		})
	}
}
