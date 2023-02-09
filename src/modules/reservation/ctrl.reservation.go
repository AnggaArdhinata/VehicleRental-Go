package reservation

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/interfaces"
	"github.com/AnggaArdhinata/backend-go/src/library"
	"github.com/gorilla/mux"
)

type reservation_ctrl struct {
	service interfaces.Reservation_ServiceIF
}

func NewCtrl(data interfaces.Reservation_ServiceIF) *reservation_ctrl {
	return &reservation_ctrl{data}
}

func (c *reservation_ctrl) FindByUserID(w http.ResponseWriter, r *http.Request) {
	payload_id := r.Context().Value("user")
	c.service.FindByIdUser(payload_id.(uint)).Send(w)
	return
}

func (c *reservation_ctrl) FindAll(w http.ResponseWriter, r *http.Request) {
	result := c.service.FindAll()
	result.Send(w)
	return

}

func (c *reservation_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var data models.Reservation
	userId := r.Context().Value("user")
	data.User_Id = userId.(uint)
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}

	c.service.Add(&data).Send(w)
}

func (c *reservation_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_int, err := strconv.Atoi(vars["id"])
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}
	id := uint(id_int)
	var data models.Reservation
	json.NewDecoder(r.Body).Decode(&data)
	c.service.Update(id, &data).Send(w)
}

func (c *reservation_ctrl) Delete(w  http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_int, err := strconv.Atoi(vars["id"])
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}
	id := uint(id_int)
	c.service.Delete(id).Send(w)
}