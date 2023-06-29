package logger

import (
	"errors"
	"fmt"
	"os"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Setup initializes the logger.
// Use "dev" for developer friendly logs.
// Use "prod" for production suitable logs (JSON).
// Use "nop" for no log at all
// Use "example" for ExampleTests (mostly internal)
// Default is "prod".
func Setup(env string) error {
	var zapConfig zap.Config
	switch env {
	case "dev":
		zapConfig = zap.NewDevelopmentConfig()
	case "example":
		zap.ReplaceGlobals(zap.NewExample())
		return nil
	case "nop":
		zap.ReplaceGlobals(zap.NewNop())
		return nil
	default:
		zapConfig = zap.NewProductionConfig()
		zapConfig.DisableStacktrace = true
		zapConfig.Sampling = nil
	}

	zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	logger, err := zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		return fmt.Errorf("failed to build the logger: %s", err)
	}

	zap.ReplaceGlobals(logger)
	return nil
}

// Flush calls the underlying logger flush, flushing any buffered log
// entries. Applications should take care to call Flush before exiting.
func Flush() {
	err := zap.L().Sync()
	if err != nil {
		// https://github.com/uber-go/zap/issues/370
		if !errors.Is(err, syscall.EINVAL) {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}
