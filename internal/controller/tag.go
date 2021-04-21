package controller

import (
	"github.com/gofiber/fiber/v2"
	"nicetry/internal/model/dto"
	"nicetry/pkg/app"
	"nicetry/pkg/e"
)

func (c *Controller) GetTag(ctx *fiber.Ctx) error {

	id, err := c.GetParamUint(ctx, "id")
	if err != nil {
	    return ctx.JSON(app.NewErr(e.InvalidParams))
	}

	t, err := c.Service.GetTag(id)

	return ctx.JSON(app.NewRes(t))
}


func (c *Controller) AddTag(ctx *fiber.Ctx) error {

	n := dto.TagAddParams{}

	if err := c.BodyParse(ctx, &n); err != nil {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), err.Error()))
	}

	if err := c.Service.AddTag(n.Name, n.ParentId); err != nil {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), err.Error()))
	}

	return ctx.JSON(app.NewRes(""))
}


func (c *Controller) DeleteTag(ctx *fiber.Ctx) error {

	id, err := c.GetParamUint(ctx, "id")
	if err != nil {
		return ctx.JSON(app.NewErr(e.InvalidParams))
	}

	t, err := c.Service.GetTag(id)

	return ctx.JSON(app.NewRes(t))
}