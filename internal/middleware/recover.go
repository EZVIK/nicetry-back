package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"nicetry/pkg/app"
	"nicetry/pkg/e"
)

func NewRecover() fiber.Handler {

	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		// Don't execute middleware if Next returns true
		//if cfg.Next != nil && cfg.Next(c) {
		//	return c.Next()
		//}

		// Catch panics
		defer func() {
			if r := recover(); r != nil {
				var ok bool
				if err, ok = r.(error); !ok {
					// Set error that will call the global error handler
					err = fmt.Errorf("%v", r)

					c.JSON(app.NewErrRes(e.ERROR, e.GetMsg(e.ERROR), err.Error()))
				}
			}
		}()

		// Return err if exist, else move to next handler
		return c.Next()
	}
}
