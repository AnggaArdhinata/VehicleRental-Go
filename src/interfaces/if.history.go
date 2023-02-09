package interfaces

import (
	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/library"
)

type History_RepoIF interface {
	// JoinData(idres uint) (*models.History, error)
	// AddBooking(data *models.History) (*models.History, error)
	GetAll() (*models.Historis, error)
	Insert(data *models.History) (*models.History, error)
	Set(id uint, res *models.History) (*models.History, error)
	Remove(id uint) (*library.Messeage, error)
}

type History_ServiceIF interface {
	// JoinReserv(idres uint) *library.Responses
	// AddFromReserv(data *models.History) *library.Responses
	FindAll() *library.Responses
	Add(data *models.History) *library.Responses
	Update(id uint, res *models.History) *library.Responses
	Delete(id uint) *library.Responses
}