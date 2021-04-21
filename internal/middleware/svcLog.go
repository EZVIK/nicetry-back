package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func SvcLogger() func(c *fiber.Ctx) error {


	return func(c *fiber.Ctx) error {



		c.Next()
		return nil
	}
}
