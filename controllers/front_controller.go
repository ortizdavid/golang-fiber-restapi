package controllers

import "github.com/gofiber/fiber/v2"

type FrontController struct {
}

func (FrontController) index(ctx *fiber.Ctx) error {
	return ctx.Render("front-office/index", fiber.Map{
		"Title": "Golang Web App",
	})
}

func (FrontController) about(ctx *fiber.Ctx) error {
	return ctx.Render("front-office/about", fiber.Map{
		"Title": "About this Example",
	})
}

func (front FrontController) RegisterRoutes(router *fiber.App) {

	router.Get("/", front.index)
	router.Get("/about", front.about)
}