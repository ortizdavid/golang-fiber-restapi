package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func ConfigStaticFiles(app *fiber.App) {
	app.Static("/", "./static")
}

func GetTemplateEngine() *html.Engine {
	engine := html.New("./templates", ".html")
	return engine
}