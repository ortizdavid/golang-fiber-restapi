package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-restapi/entities"
	"github.com/ortizdavid/golang-fiber-restapi/helpers"
	"github.com/ortizdavid/golang-fiber-restapi/models"
)

type TaskController struct {
}

func (task TaskController) RegisterRoutes(router *fiber.App) {
	group := router.Group("/api/tasks")
	group.Get("/", task.getAll)
	group.Post("/", task.create)
	group.Get("/:id", task.getTask)
	group.Put("/:id", task.update)
	group.Post("/search", task.search)
}

func (TaskController) getAll(ctx *fiber.Ctx) error {
	tasks := models.TaskModel{}.FindAll()
	count := len(tasks)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Error": false,
		"Count": count,
		"Message": "All Tasks Found",
		"Tasks": tasks,
	})
}

func (TaskController) getTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	intId := helpers.ConvertToInt(id)
	task := models.TaskModel{}.FindById(intId)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Error": false,
		"Message": "Task Found",
		"Task": task,
	})
}


func (TaskController) create(ctx *fiber.Ctx) error {
	taskName := ctx.FormValue("task_name")
	status := ctx.FormValue("status")
	description := ctx.FormValue("description")
	startDate := ctx.FormValue("start_date")
	endDate := ctx.FormValue("end_date")

	var taskModel models.TaskModel
	task := entities.Task{
		TaskId:      0,
		TaskName:    taskName,
		StartDate:   helpers.StringToDate(startDate),
		EndDate:     helpers.StringToDate(endDate),
		Description: description,
		Status:      status,
	}
	taskModel.Create(task)
	log.Printf("Task '%s' createed ", taskName)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Error": false,
		"Message": "Task Created Successfully",
		"Tasks": task,
	})
}

func (TaskController) update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	intId := helpers.ConvertToInt(id)
	taskName := ctx.FormValue("task_name")
	status := ctx.FormValue("status")
	description := ctx.FormValue("description")
	startDate := ctx.FormValue("start_date")
	endDate := ctx.FormValue("end_date")

	var taskModel models.TaskModel
	task := taskModel.FindById(intId)

	task.TaskName = taskName
	task.Status = status
	task.Description = description
	task.StartDate = helpers.StringToDate(startDate)
	task.EndDate = helpers.StringToDate(endDate)
	taskModel.Update(task)
	log.Printf("Task '%s' createed ", taskName)
	return ctx.JSON(task)
}

func (TaskController) search(ctx *fiber.Ctx) error {
	param := ctx.FormValue("search_param")
	results := models.TaskModel{}.Search(param)
	count := len(results)
	log.Printf("Search for Task '%v' and %v Results Founds", param, count)
	return ctx.JSON(results)
}
