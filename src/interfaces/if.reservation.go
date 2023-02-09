package interfaces

import (
	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/library"
)

type Reservation_RepoIF interface {
	GetById(id uint) (*models.Reservation, error)
	GetAll() (*models.Reservations, error)
	Insert(data *models.Reservation) (*models.Reservation, error)
	Set(id uint, res *models.Reservation) (*models.Reservation, error)
	Remove(id uint) (*library.Messeage, error)
}

type Reservation_ServiceIF interface {
	FindByIdUser(id uint) *library.Responses
	FindAll() *library.Responses
	Add(data *models.Reservation) *library.Responses
	Update(id uint, res *models.Reservation) *library.Responses
	Delete(id uint) *library.Responses
}