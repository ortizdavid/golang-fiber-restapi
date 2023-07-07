package models

type InterfaceModel interface {
	Create()
	FindAll()
	FindById()
	Update()
	Delete()
	Search()
}