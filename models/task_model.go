package models

import (
	"github.com/ortizdavid/golang-fiber-restapi/config"
	"github.com/ortizdavid/golang-fiber-restapi/entities"
)

type TaskModel struct {
}

func (TaskModel) Create(task entities.Task) {
	db, _ := config.ConnectDB()
	db.Create(&task)
}

func (TaskModel) FindAll() []entities.Task {
	db, _ := config.ConnectDB()
	Tasks := []entities.Task{}
	db.Find(&Tasks)
	return Tasks
}

func (TaskModel) Update(task entities.Task) {
	db, _ := config.ConnectDB()
	db.Save(&task)
}

func (TaskModel) FindById(id int) entities.Task {
	db, _ := config.ConnectDB()
	var task entities.Task
	db.First(&task, id)
	return task
}

func (TaskModel) Search(param string) []entities.Task {
	db, _ := config.ConnectDB()
	tasks := []entities.Task{}
	db.Where("task_name like ?", "%"+param+"%").Find(&tasks)
	return tasks
}
