package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func Logger() fiber.Handler {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339Nano

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "localhost"
	}

	baseLogger := zerolog.New(os.Stdout).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Str("hostname", hostname).
		Logger()

	return func(c *fiber.Ctx) error {
		start := time.Now()

		reqLogger := baseLogger.
			With().
			Str("method", c.Method()).
			Str("path", c.Path()).
			Str("remote_addr", c.Context().RemoteAddr().String()).
			Interface("request_id", c.Locals(reqIDCtxKey)).
			Bytes("user_agent", c.Context().UserAgent()).
			Logger()

		reqLogger.Info().Msg("start request")

		c.SetUserContext(reqLogger.WithContext(c.UserContext()))
		err := c.Next()

		reqLogger.
			Info().
			Int("status", c.Response().StatusCode()).
			Dur("latency", time.Since(start)).
			Msg("complete request")

		return err
	}
}
