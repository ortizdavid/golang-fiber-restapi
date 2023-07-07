package controllers

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router *fiber.App) {
	FrontController{}.RegisterRoutes(router)
	TaskController{}.RegisterRoutes(router)
}