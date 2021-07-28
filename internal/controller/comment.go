package controller

import (
	"github.com/gofiber/fiber/v2"
	"nicetry/internal/model/dto"
	"nicetry/pkg/app"
	"nicetry/pkg/convert"
	"nicetry/pkg/e"
	"nicetry/pkg/utils"
)

func (s *Controller) GetComments(ctx *fiber.Ctx) error {
	str := convert.StrTo(ctx.Params("id"))

	id, err := str.UInt()

	if err != nil {
		return ctx.JSON(app.NewErr(e.InvalidParams))
	}

	comments, err := s.Service.GetComments(id)

	if err != nil {
		return err
	}

	return ctx.JSON(app.NewRes(comments))
}

func (s *Controller) AddComment(ctx *fiber.Ctx) error {
	n := dto.CommentAddParams{}

	err := s.BodyParse(ctx, &n)

	if err != nil {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), err.Error()))
	}

	userId := utils.GetUserIdFromToken(ctx)

	err = s.Service.AddComment(n.NiceId, userId, n.Content)

	if err != nil {
		return err
	}

	return ctx.JSON(app.NewRes(""))
}
