package controller

import (
	"github.com/gofiber/fiber/v2"
	"nicetry/internal/model/dto"
	"nicetry/pkg/app"
	"nicetry/pkg/e"
	"nicetry/pkg/utils"
)

func (s *Controller) Like(ctx *fiber.Ctx) error {

	n := dto.LikeParams{}

	err := s.BodyParse(ctx, &n)

	if err != nil {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), err.Error()))
	}
	token := ctx.Get("token")
	userId := utils.GetUserIdFromToken(token)

	if userId == 0 {
		return ctx.JSON(app.NewErr(e.UserIDParseError))
	}

	err = s.Service.Like(n.PostId, n.LikeType, userId)

	if  err != nil {
		return ctx.JSON(app.NewErrRes(e.ERROR_CREATE_FAIL, e.GetMsg(e.ERROR_CREATE_FAIL), err.Error()))
	}

	return ctx.JSON(app.NewRes(""))
}
