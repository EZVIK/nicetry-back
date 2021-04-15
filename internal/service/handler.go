package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"nicetry/pkg/app"
	"time"
)

type Handler struct {
	Validate *validator.Validate
}

func (h *Handler) HandlerException(ctx *fiber.Ctx, err error) bool {

	if err != nil {
		cause :=  errors.Cause(err)

		switch cause.(type) {
		default:
			fmt.Println(time.Now().Format("2006-01-02 15:04:05")," = ",err)
			ctx.JSON(app.NewErrRes(500, "Server Interal ERROR.", cause.Error()))
		}
		return true
	}
	return false
}

/**
	解析 request body 参数
 */
func (h *Handler) BodyParse(ctx *fiber.Ctx, dto interface{}) (error) {

	_ = ctx.BodyParser(dto)					    // 解析参数
	validateError := h.Validate.Struct(dto)		// 校验参数
	if validateError != nil {
		//h.HandlerException(ctx, validateError)
		return validateError
	}
	return nil
}