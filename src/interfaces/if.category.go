package interfaces

import (
	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/library"
)

type Category_RepoIF interface {
	GetById(id uint) (*models.Category, error)
	GetAll() (*models.Categories, error)
	Insert(data *models.Category) (*models.Category, error)
	Set(id uint, res *models.Category) (*models.Category, error)
	Remove(id uint) (*library.Messeage, error)
}

type Category_ServiceIF interface {
	FindById(id uint) *library.Responses
	FindAll() *library.Responses
	Add(data *models.Category) *library.Responses
	Update(id uint, res *models.Category) *library.Responses
	Delete(id uint) *library.Responses
}