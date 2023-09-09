package logging

import (
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Get(logLevel string) zerolog.Logger {

	var level zerolog.Level
	switch logLevel {
	case "trace":
		level = zerolog.TraceLevel
	case "debug":
		level = zerolog.DebugLevel
	case "info":
		level = zerolog.InfoLevel
	case "warn":
		level = zerolog.WarnLevel
	case "error":
		level = zerolog.ErrorLevel
	default:
		level = zerolog.InfoLevel
	}

	log := (zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
		Level(level)).
		With().
		Timestamp().
		Logger()
	return log
}

type Logger struct {
	zerolog.Logger
}

func GetLogger(level string) *Logger {
	return &Logger{Get(level)}
}

func (l *Logger) HandlerLog(r *http.Request, status int, msg string) {
	code := strconv.Itoa(status)
	l.Info().Str("method", r.Method).
		Str("host", r.Host).
		Str("URL", r.RequestURI).
		Str("from", r.RemoteAddr).
		Str("status", code).
		Msg(msg)
}

func (l *Logger) HandlerErrorLog(r *http.Request, status int, msg string, err error) {
	code := strconv.Itoa(status)
	l.Error().Str("method", r.Method).
		Str("host", r.Host).
		Str("URL", r.RequestURI).
		Str("from", r.RemoteAddr).
		Str("status", code).
		Err(err).
		Msg(msg)
}

func (l *Logger) Fatal(msg string, err error) {
	l.Error().Err(err).Msg(msg)
	os.Exit(1)
}
