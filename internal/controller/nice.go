package controller

import (
	"github.com/gofiber/fiber/v2"
	"nicetry/internal/model/dto"
	"nicetry/pkg/app"
	"nicetry/pkg/convert"
	"nicetry/pkg/e"
)

func (s *Controller) GetNice(ctx *fiber.Ctx) error {
	str := convert.StrTo(ctx.Params("id"))

	id,err := str.UInt()

	if err != nil {
		return ctx.JSON(app.NewErr(e.InvalidParams))
	}

	nice, err := s.Service.Get(id)
	if err != nil {
		return ctx.JSON(app.NewErr(e.InvalidParams))
	}

	return ctx.JSON(app.NewRes(nice))
}

func (s *Controller) AddNice(ctx *fiber.Ctx) error {

	n := dto.NiceParams{}

	err := s.BodyParse(ctx, &n)

	if err != nil {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), err.Error()))
	}

	nice, err := s.Service.AddNice(n.Title, n.Desc, n.Content, n.NiceType, n.Tags)

	if  err != nil {
		return ctx.JSON(app.NewErrRes(e.ERROR_CREATE_FAIL, e.GetMsg(e.ERROR_CREATE_FAIL), err.Error()))
	}

	return ctx.JSON(app.NewRes(nice.ID))
}

func (s *Controller) UpdateNice(ctx *fiber.Ctx)  error {
	id := ctx.Get("id")
	return ctx.JSON(app.NewRes(id + " update."))
}

func (s *Controller) DeleteNice(ctx *fiber.Ctx)  error {
	id := ctx.Get("id")

	return ctx.JSON(app.NewRes(id + " deleted."))
}