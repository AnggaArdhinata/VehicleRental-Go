package interfaces

import (
	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/library"
)

type Users_RepoIF interface {
	GetById(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByName(username string) (*models.User, error)
	GetAll() (*models.Users, error)
	Insert(data *models.User) (*models.User, error)
	Set(id uint, res *models.User) (*models.User, error)
	Remove(id uint) (*library.Messeage, error)
}

type User_ServiceIF interface {
	FindById(id uint) *library.Responses
	FindByUserEmail(email string) *library.Responses
	FindByUserName(username string) *library.Responses
	FindAll() *library.Responses
	Add(data *models.User) *library.Responses
	Update(id uint, res *models.User) *library.Responses
	Delete(id uint) *library.Responses
}