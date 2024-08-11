package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger = zerolog.Logger

func NewLogger(isJsonLog bool) Logger {
	consoleWriter := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime})

	if isJsonLog {
		consoleWriter = zerolog.New(os.Stderr)
	}

	return consoleWriter.With().Timestamp().Logger()
}
