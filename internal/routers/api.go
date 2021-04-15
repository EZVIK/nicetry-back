package routers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"nicetry/internal/service"
)

func InitFiber(app *fiber.App){

	handler := service.Handler{
		Validate: validator.New(),				// 参数校验
	}

	api := app.Group("/v1/api/")

	// Nice
	nice := api.Group("nice")

	nice.Get("/:id", handler.Get)
	nice.Put("/:id", handler.Update)
	nice.Delete("/:id", handler.Delete)
	nice.Post("/", handler.Add)

}