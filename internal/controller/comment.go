package controller

import (
	"github.com/gofiber/fiber/v2"
	"nicetry/internal/model/dto"
	"nicetry/pkg/app"
	"nicetry/pkg/convert"
	"nicetry/pkg/e"
	"nicetry/pkg/utils"
)

func (c *Controller) GetComments(ctx *fiber.Ctx) error {
	str := convert.StrTo(ctx.Params("id"))

	id, err := str.UInt()

	if err != nil {
		return ctx.JSON(app.NewErr(e.InvalidParams))
	}

	comments, err := c.Service.GetComments(id)

	if err != nil {
		return err
	}

	return ctx.JSON(app.NewRes(comments))
}

func (c *Controller) AddComment(ctx *fiber.Ctx) error {
	n := dto.CommentAddParams{}

	err := c.BodyParse(ctx, &n)

	if err != nil {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), err.Error()))
	}

	userId := utils.GetUserIdFromToken(ctx)

	err = c.Service.AddComment(n.NiceId, userId, n.Content)

	if err != nil {
		return err
	}

	return ctx.JSON(app.NewRes(""))
}
