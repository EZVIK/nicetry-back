package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"nicetry/internal/model/dto"
	"nicetry/pkg/app"
	"nicetry/pkg/e"
	"nicetry/pkg/utils"
	"strconv"
)

func (s *Controller) GetNice(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	nice, err := s.Service.GetNice(id)
	if err != nil {
		return ctx.JSON(app.NewErr(e.InvalidParams))
	}

	ids := []uint{}
	for _, comment := range nice.Comments {
		ids = append(ids, comment.UserId)
	}

	iUsers, _ := s.Service.GetUsersAvatar(ids)

	userInfo := make(map[string]interface{})

	st := ToMap(nice)
	for _, k := range iUsers {
		userInfo[strconv.Itoa(int(k.ID))] = k
	}
	st["userInfo"] = userInfo
	return ctx.JSON(app.NewRes(st))
}

func (s *Controller) AddNice(ctx *fiber.Ctx) error {

	n := dto.NiceParams{}

	err := s.BodyParse(ctx, &n)

	if err != nil {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), err.Error()))
	}

	UserId := utils.GetUserIdFromToken(ctx)

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

func ToMap(in2 interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	inrec, _ := json.Marshal(in2)
	json.Unmarshal(inrec, &m)
	return m
}
