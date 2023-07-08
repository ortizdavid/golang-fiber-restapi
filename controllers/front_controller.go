package controllers

import "github.com/gofiber/fiber/v2"

type FrontController struct {
}

func (FrontController) index(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello API")
}

func (FrontController) about(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello API")
}

func (front FrontController) RegisterRoutes(router *fiber.App) {

	router.Get("/", front.index)
	router.Get("/about", front.about)
}