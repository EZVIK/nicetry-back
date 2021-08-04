package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/ratelimit"
	"nicetry/global"
	"nicetry/internal/controller"
	"nicetry/internal/middleware"
	"time"
)

func InitFiber(app *fiber.App) {

	ctr := controller.New()

	jwt := middleware.JWT()

	//limiter := middleware.Limiter{RL: ratelimit.New(200)}     // request per second

	uploadLimit := middleware.Limiter{RL: ratelimit.New(5)} // request per second

	// 图片访问
	app.Static("/FileOp", global.AppSetting.ImageFilePath, fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "john.html",
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	//api.Use(limiter.Take())
	app.Use(middleware.CORS())

	corsConfig := cors.ConfigDefault
	corsConfig.AllowOrigins = "*"
	corsConfig.AllowHeaders = "Authorization, token, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, accept, origin, Cache-Control, X-Requested-With"
	app.Use(cors.New(corsConfig))

	//app.Use(middleware.NewRecover())
	api := app.Group("/api/v1")

	api.Post("/upload", uploadLimit.Take(), ctr.UploadImage)

	// Group
	user := api.Group("user")
	nice := api.Group("nice", jwt)
	like := api.Group("like", jwt)
	comment := api.Group("comment", jwt)
	tag := api.Group("tag", jwt)

	// User
	user.Post("/login", ctr.Login)
	user.Post("/register", ctr.Register)
	user.Post("/", ctr.GetUsers)
	user.Get("/rf", ctr.GetReferralCode)

	// Nice
	nice.Get("/:id", ctr.GetNice)
	nice.Get("/", ctr.GetNicelist)
	nice.Put("/:id", ctr.UpdateNice)
	nice.Delete("/:id", ctr.DeleteNice)
	nice.Post("/", ctr.AddNice)

	// Like
	like.Post("/", ctr.Like)

	// Comment
	comment.Get("/:id", ctr.GetComments)
	comment.Post("/", ctr.AddComment)

	// Tag
	tag.Get("/:id", ctr.GetTag)
	tag.Post("/", ctr.AddTag)
	tag.Delete("/:id", ctr.DeleteTag)

}
