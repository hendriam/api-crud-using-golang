package lib

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logging struct {
	*zerolog.Logger
}

func LoadLogging() Logging {
	config := LoadConfig()

	// zerolog.SetGlobalLevel(zerolog.InfoLevel)

	logger := zerolog.
		New(os.Stdout).
		Level(logLevel(config.Log.Level)).
		With().
		Timestamp().
		Logger()

	output := logger.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.DateTime,
	})

	return Logging{&output}
}

func logLevel(logLevel string) zerolog.Level {
	switch logLevel {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}
