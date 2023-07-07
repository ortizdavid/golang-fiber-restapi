package main

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-webapp/config"
	"github.com/ortizdavid/golang-fiber-webapp/entities"
	"github.com/ortizdavid/golang-fiber-webapp/controllers"
)

func main() {

	app := fiber.New(fiber.Config{
        Views: config.GetTemplateEngine(),
    })

	entities.SetupMigrations()
	config.LoadDotEnv()
	config.ConfigStaticFiles(app)
	controllers.SetupRoutes(app)

	app.Listen(os.Getenv("APP_HOST")+":"+os.Getenv("APP_PORT"))
}