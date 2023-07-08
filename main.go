package main

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-restapi/config"
	"github.com/ortizdavid/golang-fiber-restapi/entities"
	"github.com/ortizdavid/golang-fiber-restapi/controllers"
)

func main() {

	app := fiber.New()

	entities.SetupMigrations()
	config.LoadDotEnv()
	config.ConfigStaticFiles(app)
	controllers.SetupRoutes(app)

	app.Listen(os.Getenv("APP_HOST")+":"+os.Getenv("APP_PORT"))
}