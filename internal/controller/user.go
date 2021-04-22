package controller

import (
	"github.com/gofiber/fiber/v2"
	"nicetry/internal/model/dto"
	"nicetry/pkg/app"
	"nicetry/pkg/e"
)

func (c *Controller) Login(ctx *fiber.Ctx) error {

	n := dto.LoginParams{}

	err := c.BodyParse(ctx, &n)

	if err != nil {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), err.Error()))
	}

	token, err := c.Service.Login(n.Mail, n.Password)

	return ctx.JSON(app.NewRes(token))
}


func (c *Controller) Register(ctx *fiber.Ctx) error {

	n := dto.RegisterParams{}

	err := c.BodyParse(ctx, &n)

	if err != nil {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), err.Error()))
	}

	if err := c.Service.Register(n.ReferralCode, n.Nickname, n.Password, n.Mail, "-", "-", "-"); err != nil {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), err.Error()))
	}

	return ctx.JSON(app.NewRes("Success"))
}

func (c *Controller) GetUsers(ctx *fiber.Ctx) error {

	n := dto.IUsers{}

	err := c.BodyParse(ctx, &n)

	if err != nil {
		return ctx.JSON(app.NewErrRes(e.INVALID_PARAMS, e.GetMsg(e.INVALID_PARAMS), err.Error()))
	}

	token, err := c.Service.GetUser(n)

	return ctx.JSON(app.NewRes(token))
}


