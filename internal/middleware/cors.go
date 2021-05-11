package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func CORS() func(c *fiber.Ctx) error {

	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set("Access-Control-Allow-Headers", "token, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Set("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Method() == "OPTIONS" {
			c.Status(204)
			return nil
		}

		c.Next()
		return nil
	}
}
