package logger

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"io"
	"os"
	"time"
)

var Logger = New(zerolog.ConsoleWriter{
	Out:        os.Stdout,
	TimeFormat: time.RFC3339,
})

func New(writer io.Writer) *zerolog.Logger {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339Nano

	logger := zerolog.New(writer).
		Level(zerolog.DebugLevel).
		With().
		Timestamp().
		Logger()

	return &logger
}

func Info(ctx context.Context, msg string) {
	zerolog.Ctx(ctx).Info().Msg(msg)
}

func Infof(ctx context.Context, msg string, fields ...interface{}) {
	zerolog.Ctx(ctx).Info().Msgf(msg, fields...)
}

func Warn(ctx context.Context, msg string) {
	zerolog.Ctx(ctx).Warn().Msg(msg)
}

func Warnf(ctx context.Context, msg string, fields ...interface{}) {
	zerolog.Ctx(ctx).Warn().Msgf(msg, fields...)
}

func Debug(ctx context.Context, msg string) {
	zerolog.Ctx(ctx).Info().Msg(msg)
}

func Debugf(ctx context.Context, msg string, fields ...interface{}) {
	zerolog.Ctx(ctx).Debug().Msgf(msg, fields...)
}

func Error(ctx context.Context, msg string) {
	zerolog.Ctx(ctx).Error().Msg(msg)
}

func Errorf(ctx context.Context, msg string, fields ...interface{}) {
	zerolog.Ctx(ctx).Error().Msgf(msg, fields...)
}
