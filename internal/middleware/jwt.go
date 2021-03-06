package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"nicetry/pkg/app"
	"nicetry/pkg/e"
	"nicetry/pkg/utils"
)

func JWT() func(c *fiber.Ctx) error {

	return func(c *fiber.Ctx) error {
		var code int

		code = e.SUCCESS
		token := c.Get("Authorization")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			return c.JSON(app.NewErr(e.UnauthorizedTokenError))
		}

		c.Next()
		return nil
	}
}
