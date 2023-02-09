package history

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/interfaces"
	"github.com/AnggaArdhinata/backend-go/src/library"
	"github.com/gorilla/mux"
)

type history_ctrl struct {
	service interfaces.History_ServiceIF
}

func NewCtrl(data interfaces.History_ServiceIF) *history_ctrl {
	return &history_ctrl{data}
}

// func (c *history_ctrl) JoinWithRes(w http.ResponseWriter, r *http.Request) {
// 	var id models.Reservation
// 	idres := id.Id
// 	result := c.service.JoinReserv(idres)
// 	result.Send(w)
// 	return

// }

func (c *history_ctrl) FindAll(w http.ResponseWriter, r *http.Request) {
	result := c.service.FindAll()
	result.Send(w)
	return

}

// func (c history_ctrl) AddFromReserv(w http.ResponseWriter, r *http.Request){
// 	var data models.History
// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		library.Response(err.Error(), 500, true).Send(w)
// 		return
// 	}
// 	c.service.AddFromReserv(&data).Send(w)
// }

func (c *history_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var data models.History
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}
	c.service.Add(&data).Send(w)
}

func (c *history_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_int, err := strconv.Atoi(vars["id"])
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}
	id := uint(id_int)
	var data models.History
	json.NewDecoder(r.Body).Decode(&data)
	c.service.Update(id, &data).Send(w)
}

func (c *history_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_int, err := strconv.Atoi(vars["id"])
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}
	id := uint(id_int)
	c.service.Delete(id).Send(w)
}
