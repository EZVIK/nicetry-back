package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/errgo.v2/errors"
	"nicetry/internal/dao"
	"nicetry/pkg/app"
	"time"
)

type Service struct {
	Dao *dao.Dao
}

/***
异常处理
*/
func (s *Service) HandlerException(ctx *fiber.Ctx, err error) bool {

	if err != nil {
		cause := errors.Cause(err)

		switch cause.(type) {
		default:
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " = ", err)
			ctx.JSON(app.NewErrRes(500, "Server Internal ERROR.", cause.Error()))
		}
		return true
	}
	return false
}
