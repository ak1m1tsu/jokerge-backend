package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var (
	reqIDCtxKey = "reqID"
	reqIDHeader = "X-Request-ID"
)

func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqID := c.Get(reqIDHeader)
		if reqID == "" {
			reqID = uuid.NewString()
		}

		c.Set(reqIDHeader, reqID)
		c.Locals(reqIDCtxKey, reqID)

		return c.Next()
	}
}
