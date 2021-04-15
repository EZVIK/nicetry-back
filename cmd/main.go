package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"nicetry/global"
	"nicetry/internal/model"
	"nicetry/internal/routers"
	"nicetry/pkg/logger"
	"nicetry/pkg/setting"
	"time"
)

func init() {
	log.Println("Initializing...\n")

	fmt.Print("Setting Initializing...")
	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	fmt.Println("Finished.\n")

	fmt.Print("DBEngine Initializing...")
	if err := setupDBEngine(); err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	fmt.Println("Finished.\n")

	fmt.Print("Logger Initializing...")
	if err := setupLogger(); err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	fmt.Println("Finished.\n")

	global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")
}

func main() {
	fSetting := fiber.Config{
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
	}
	app := fiber.New(fSetting)
	routers.InitFiber(app)
	app.Listen(":3009")
}

func setupSetting()  error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger()   error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}