package middleware

type ThumbsUp_ interface {
	Like() error
}

func ThumbsUp(t ThumbsUp_) error {
	return t.Like()
}

//func Like() func(ctx *fiber.Ctx) error {
//
//	return func(ctx *fiber.Ctx) error {
//		return nil
//	}
//}