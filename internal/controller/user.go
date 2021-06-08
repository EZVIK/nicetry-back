package controller

import (
	"github.com/gofiber/fiber/v2"
	"nicetry/internal/model/dto"
	"nicetry/pkg/app"
	"nicetry/pkg/e"
	"nicetry/pkg/utils"
)

func (c *Controller) Login(ctx *fiber.Ctx) error {

	n := dto.LoginParams{}
	if err := c.BodyParse(ctx, &n); err != nil {
		return ctx.JSON(app.NewErr(e.InvalidParams))
	}

	user, token, err := c.Service.Login(n.Mail, n.Password)

	if err != nil || token == "" {
		return ctx.JSON(app.NewErr(e.UnauthorizedFail))
	}

	retData := make(map[string]interface{})
	retData["userInfo"] = user
	retData["jwt"] = token
	return ctx.JSON(app.NewRes(retData))
}

func (c *Controller) Register(ctx *fiber.Ctx) error {

	n := dto.RegisterParams{}

	if err := c.BodyParse(ctx, &n); err != nil {
		return ctx.JSON(app.NewErr(e.InvalidParams))
	}

	if err := c.Service.Register(n.ReferralCode, n.Nickname, n.Password, n.Mail, "-", "-", "-"); err != nil {
		return ctx.JSON(app.NewErr(e.InvalidParams))
	}

	return ctx.JSON(app.NewRes("Success"))
}

// 获取用户信息
func (c *Controller) GetUsers(ctx *fiber.Ctx) error {

	n := dto.IUsers{}
	if err := c.BodyParse(ctx, &n); err != nil {
		return ctx.JSON(app.NewErr(e.InvalidParams))
	}

	user, err := c.Service.GetUser(n.Ids)
	if err != nil {
		return ctx.JSON(app.NewErr(e.ERROR_GET_USER_FAIL))
	}

	return ctx.JSON(app.NewRes(user))
}

// 获取用户推荐码
func (c *Controller) GetReferralCode(ctx *fiber.Ctx) error {

	userId := utils.GetUserIdFromToken(ctx)

	rf, err := c.Service.GetReferralCode(userId)
	if err != nil {
		return ctx.JSON(app.NewErr(e.ERROR_GET_USER_FAIL))
	}

	return ctx.JSON(app.NewRes(rf))
}
