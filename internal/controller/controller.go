package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"nicetry/global"
	"nicetry/internal/dao"
	"nicetry/internal/service"
	"nicetry/pkg/convert"
)

type Controller struct {
	Service  *service.Service
	Validate *validator.Validate
}

func New() *Controller {
	svc := service.Service{}
	svc.Dao = dao.New(global.DBEngine, global.CacheEngine)
	ctrl := Controller{
		Service:  &svc,
		Validate: validator.New(),
	}
	return &ctrl
}

// BodyParse
// parse request body to dto
func (s *Controller) BodyParse(ctx *fiber.Ctx, dto interface{}) error {
	_ = ctx.BodyParser(dto)                 // 解析参数
	validateError := s.Validate.Struct(dto) // 校验参数
	if validateError != nil {
		//h.HandlerException(ctx, validateError)
		return validateError
	}
	return nil
}

// GetParamUint
// string to uint
func (s *Controller) GetParamUint(ctx *fiber.Ctx, key string) (uint, error) {
	str := convert.StrTo(ctx.Params(key))
	id, err := str.UInt()
	if err != nil {
		return 0, err
	}
	return id, nil
}
