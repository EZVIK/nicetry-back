package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"nicetry/internal/model"
	"nicetry/internal/model/dto"
	"nicetry/pkg/app"
	"nicetry/pkg/e"
)



func (h *Handler) Get(ctx *fiber.Ctx) error {
	nice := model.Nice{"只需回答有或没有", "-",
		"这是一个每周更新的人生挑战节目。由特特、文森特动物园、realEnjolras、cbvivi 共同主持。",
		0, 0, 1, 0, nil, gorm.Model{}}

	return ctx.JSON(app.NewRes(nice))
}

func (h *Handler) Update(ctx *fiber.Ctx) error {
	id := ctx.Get("id")

	return ctx.JSON(app.NewRes(id + " update."))
}

func (h *Handler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Get("id")

	return ctx.JSON(app.NewRes(id + " deleted."))
}

/**
    新增 Nice
 */
func (h *Handler) Add(ctx *fiber.Ctx) error {

	n := dto.NiceParams{}
	err := h.BodyParse(ctx, &n)

	if err != nil {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), err.Error()))
	}

	return ctx.JSON(app.NewRes(n))
}