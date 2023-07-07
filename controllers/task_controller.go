package controllers

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-webapp/entities"
	"github.com/ortizdavid/golang-fiber-webapp/helpers"
	"github.com/ortizdavid/golang-fiber-webapp/models"
)

type TaskController struct {
}

func (task TaskController) RegisterRoutes(router *fiber.App) {
	group := router.Group("/tasks")
	group.Get("/", task.index)
	group.Get("/add", task.addForm)
	group.Post("/add", task.add)
	group.Get("/:id/details", task.details)
	group.Get("/:id/edit", task.editForm)
	group.Post("/:id/edit", task.edit)
	group.Get("/search", task.searchForm)
	group.Post("/search", task.search)
}

func (TaskController) index(ctx *fiber.Ctx) error {
	tasks := models.TaskModel{}.FindAll()
	count := len(tasks)
	return ctx.Render("task/index", fiber.Map{
		"Title": "All Tasks",
		"Tasks": tasks,
		"Count": count,
	})
}

func (TaskController) details(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	intId := helpers.ConvertToInt(id)
	task := models.TaskModel{}.FindById(intId)
	return ctx.Render("task/details", fiber.Map{
		"Title": "Task Details",
		"Task": task,
	})
}

func (TaskController) addForm(ctx *fiber.Ctx) error {
	return ctx.Render("task/add", fiber.Map{
		"Title": "Add Tasks",
	})
}

func (TaskController) add(ctx *fiber.Ctx) error {
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
	log.Printf("Task '%s' Added ", taskName)
	return ctx.Redirect("/tasks")
}


func (TaskController) editForm(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	intId := helpers.ConvertToInt(id)
	task := models.TaskModel{}.FindById(intId)
	return ctx.Render("task/edit", fiber.Map{
		"Title": "Edit Task",
		"Task": task,
	})
}

func (TaskController) edit(ctx *fiber.Ctx) error {
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
	log.Printf("Task '%s' Added ", taskName)
	return ctx.Redirect(fmt.Sprintf("/tasks/%s/details", id))
}

func (TaskController) searchForm(ctx *fiber.Ctx) error {
	return ctx.Render("task/search", fiber.Map{
		"Title": "Search Tasks",
	})
}

func (TaskController) search(ctx *fiber.Ctx) error {
	param := ctx.FormValue("search_param")
	results := models.TaskModel{}.Search(param)
	count := len(results)
	log.Printf("Search for Task '%v' and %v Results Founds", param, count)
	return ctx.Render("task/search-results", fiber.Map{
		"Title": "Results",
		"Results": results,
		"Param": param,
		"Count": count,
	})
}
