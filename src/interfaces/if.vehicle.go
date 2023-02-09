package interfaces

import (
	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/library"
)

type Vehicles_RepoIF interface {
	GetAll() (*models.Vehicles, error)
	Insert(data *models.Vehicle) (*models.Vehicle, error)
	Set(id uint, res *models.Vehicle) (*models.Vehicle, error)
	Remove(id uint) (*library.Messeage, error)
}

type Vehicles_ServiceIF interface {
	FindAll() *library.Responses
	Add(data *models.Vehicle) *library.Responses
	Update(id uint, res *models.Vehicle) *library.Responses
	Delete(id uint) *library.Responses
}
