package logging

import (
	"errors"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/bakins/php-app-server/internal/stackdriver"
)

func NewLogger() *zap.Logger {
	enc := stackdriver.Encoder()

	core := zapcore.NewCore(enc, Stdout, zap.NewAtomicLevelAt(zap.InfoLevel))

	c := stackdriver.WrapCore(core, "", "")

	logger := zap.New(
		c,
		zap.ErrorOutput(Stderr),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return logger
}

// LoggingError wraps an error with a logger
type LoggingError struct {
	err     error
	logger  *zap.Logger
	message string
}

func (l *LoggingError) Error() string {
	return fmt.Sprintf("%s %v", l.message, l.err)
}

// Create a new logging error
func NewLoggingError(logger *zap.Logger, message string, err error) *LoggingError {
	l := LoggingError{
		err:     err,
		logger:  logger,
		message: message,
	}

	return &l
}

// Exit will check if error is a logging error, then log it
// other wise will print message.
// Either way, it then exits
func Exit(err error) {
	var le *LoggingError
	if errors.As(err, &le) {
		le.logger.Fatal(le.message, zap.Error(err))
		return
	}

	fmt.Fprintf(Stderr, "%v", err)
	os.Exit(1)
}

var (
	Stdout = zapcore.AddSync(os.Stdout)
	Stderr = zapcore.AddSync(os.Stderr)
)
