package controller

import (
	"github.com/gofiber/fiber/v2"
	"nicetry/internal/model/dto"
	"nicetry/pkg/app"
	"nicetry/pkg/e"
	"nicetry/pkg/utils"
)

// 登陆
func (s *Controller) Login(ctx *fiber.Ctx) error {

	n := dto.LoginParams{}
	if err := s.BodyParse(ctx, &n); err != nil {
		return ctx.JSON(app.NewErr(e.InvalidParams))
	}

	user, token, err := s.Service.Login(n.Mail, n.Password)

	if err != nil || token == "" {
		return ctx.JSON(app.NewErr(e.UnauthorizedFail))
	}

	retData := make(map[string]interface{})
	retData["userInfo"] = user
	retData["jwt"] = token
	return ctx.JSON(app.NewRes(retData))
}

// 注册
func (s *Controller) Register(ctx *fiber.Ctx) error {

	n := dto.RegisterParams{}

	if err := s.BodyParse(ctx, &n); err != nil {
		return ctx.JSON(app.NewErr(e.InvalidParams))
	}

	if err := s.Service.Register(n.ReferralCode, n.Nickname, n.Password, n.Mail, "-", "-", "-"); err != nil {
		return ctx.JSON(app.NewErr(e.InvalidParams))
	}

	return ctx.JSON(app.NewRes("Success"))
}

// 获取用户信息
func (s *Controller) GetUsers(ctx *fiber.Ctx) error {

	n := dto.IUsers{}
	if err := s.BodyParse(ctx, &n); err != nil {
		return ctx.JSON(app.NewErr(e.InvalidParams))
	}

	user, err := s.Service.GetUser(n.Ids)
	if err != nil {
		return ctx.JSON(app.NewErr(e.ERROR_GET_USER_FAIL))
	}

	return ctx.JSON(app.NewRes(user))
}

// 获取用户推荐码
func (s *Controller) GetReferralCode(ctx *fiber.Ctx) error {

	userId := utils.GetUserIdFromToken(ctx)

	rf, err := s.Service.GetReferralCode(userId)
	if err != nil {
		return ctx.JSON(app.NewErr(e.ERROR_GET_USER_FAIL))
	}

	return ctx.JSON(app.NewRes(rf))
}
