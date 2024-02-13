package middlewares

import "github.com/gofiber/fiber/v3"

func RealIP(c fiber.Ctx) error {
	clientIP := c.Get("X-Forwarded-For")
	if clientIP == "" {
		clientIP = c.IP()
	}
	return c.Next()
}
