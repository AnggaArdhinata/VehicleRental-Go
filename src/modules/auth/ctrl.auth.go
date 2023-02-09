package auth

import (
	"encoding/json"
	"net/http"

	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/interfaces"
	"github.com/AnggaArdhinata/backend-go/src/library"
	"github.com/asaskevich/govalidator"
)

type auth_ctrl struct {
	service interfaces.Auth_serviceIF
}

func NewCtrl(service interfaces.Auth_serviceIF) *auth_ctrl {
	return &auth_ctrl{service}
}

func (c *auth_ctrl) SignIn(w http.ResponseWriter, r *http.Request) {
	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		library.Response(err.Error(), 500, true)
		return
	}
	c.service.Login(&data).Send(w)
}

func (c *auth_ctrl) Register(w http.ResponseWriter, r *http.Request) {
	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}
	_, err = govalidator.ValidateStruct(data)
	if err != nil {
		library.Response(err, 500, true).Send(w)
		return
	}
	c.service.Register(&data).Send(w)
}
