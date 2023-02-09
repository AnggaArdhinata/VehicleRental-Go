package interfaces

import (
	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/library"
)

type Auth_serviceIF interface {
	Login(body *models.User) *library.Responses
	Register(data *models.User) *library.Responses
}