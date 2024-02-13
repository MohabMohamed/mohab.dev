package routes

import (
	"github.com/MohabMohamed/mohab.dev/src/config"
	"github.com/MohabMohamed/mohab.dev/src/util"
	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func Init() {
	app = fiber.New()
	app.Get("/health", checkHealth)
}

func Serve() {
	util.Logger.Fatal("error while serving the app", app.Listen(":"+config.GetEnv("PORT", "8000")))
}
