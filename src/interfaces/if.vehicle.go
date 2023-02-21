package interfaces

import (
	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/library"
)

type Vehicles_RepoIF interface {
	GetById(id uint) (*models.Vehicle, error)
	GetAll() (*models.Vehicles, error)
	GetByRating() (*models.Vehicles, error)
	Insert(data *models.Vehicle) (*models.Vehicle, error)
	Set(id uint, res *models.Vehicle) (*models.Vehicle, error)
	Remove(id uint) (*library.Messeage, error)
}

type Vehicles_ServiceIF interface {
	FindById(id uint) *library.Responses
	FindAll() *library.Responses
	FindByRating() *library.Responses
	Add(data *models.Vehicle) *library.Responses
	Update(id uint, res *models.Vehicle) *library.Responses
	Delete(id uint) *library.Responses
}
