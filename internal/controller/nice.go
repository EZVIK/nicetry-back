package controller

import (
	"github.com/gofiber/fiber/v2"
	"nicetry/internal/model/dto"
	"nicetry/pkg/app"
	"nicetry/pkg/e"
	"nicetry/pkg/utils"
)

func (s *Controller) GetNice(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	nice, err := s.Service.GetNice(id)
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
	token := ctx.Get("token")
	UserId := utils.GetUserIdFromToken(token)

	if UserId == 0 {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), "GetUserIdFromToken fail"))
	}

	nice, err := s.Service.AddNice(n.Title, n.Desc, n.Content, UserId, n.NiceType, n.Tags)

	if err != nil {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), err.Error()))
	}

	return ctx.JSON(app.NewRes(nice.ID))
}

func (s *Controller) UpdateNice(ctx *fiber.Ctx) error {
	id := ctx.Get("id")
	return ctx.JSON(app.NewRes(id + " update."))
}

func (s *Controller) DeleteNice(ctx *fiber.Ctx) error {
	id := ctx.Get("id")

	return ctx.JSON(app.NewRes(id + " deleted."))
}

func (c *Controller) GetNicelist(ctx *fiber.Ctx) error {

	n := dto.NiceListParams{}

	if err := c.BodyParse(ctx, &n); err != nil {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), err.Error()))
	}

	niceList, err := c.Service.GetNiceList("", "", n.PageSize, n.PageIndex)
	if err != nil {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), err.Error()))
	}

	return ctx.JSON(app.NewRes(niceList))
}
