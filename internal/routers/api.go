package routers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/ratelimit"
	"nicetry/global"
	"nicetry/internal/controller"
	"nicetry/internal/middleware"
	"time"
)

func InitFiber(app *fiber.App){

	ctr := controller.New()

	jwt := middleware.JWT()

	limiter := middleware.Limiter{RL: ratelimit.New(200)}     // request per second


	// 图片访问
	app.Static("/i", global.AppSetting.ImageFilePath, fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "john.html",
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	api := app.Group("/api/v1")

	api.Use(limiter.Take())

	api.Post("/upload", ctr.UploadImage)

	// User
	user := api.Group("user")
	user.Post("/login", ctr.Login)
	user.Post("/register", ctr.Register)
	user.Post("/", ctr.GetUsers)

	// Nice
	nice := api.Group("nice")
	nice.Get("/:id", jwt, ctr.GetNice)
	nice.Get("/", jwt, ctr.GetNicelist)
	nice.Put("/:id", jwt, ctr.UpdateNice)
	nice.Delete("/:id", jwt, ctr.DeleteNice)
	nice.Post("/", jwt, ctr.AddNice)
	
	// Like
	like := api.Group("like")
	like.Post("/", ctr.Like)

	// Comment
	comment := api.Group("comment")
	comment.Get("/:id", ctr.GetComments)
	comment.Post("/", ctr.AddComment)

	// Tag
	tag := api.Group("tag")
	tag.Get("/:id", ctr.GetTag)
	tag.Post("/", ctr.AddTag)
	tag.Delete("/:id", ctr.DeleteTag)
}

