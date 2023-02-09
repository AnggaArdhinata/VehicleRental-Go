package category

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/interfaces"
	"github.com/AnggaArdhinata/backend-go/src/library"
	"github.com/gorilla/mux"
)

type category_ctrl struct {
	service interfaces.Category_ServiceIF
}

func NewCtrl(data interfaces.Category_ServiceIF) *category_ctrl {
	return &category_ctrl{data}
}

func (c *category_ctrl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_int, err := strconv.Atoi(vars["id"])
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}
	id := uint(id_int)
	c.service.FindById(id).Send(w)
}

func (c *category_ctrl) FindAll(w http.ResponseWriter, r *http.Request) {
	result := c.service.FindAll()
	result.Send(w)
	return

}

func (c *category_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var data models.Category
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}
	c.service.Add(&data).Send(w)
}

func (c *category_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_int, err := strconv.Atoi(vars["id"])
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}
	id := uint(id_int)
	var data models.Category
	json.NewDecoder(r.Body).Decode(&data)
	c.service.Update(id, &data).Send(w)
}

func (c *category_ctrl) Delete(w  http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_int, err := strconv.Atoi(vars["id"])
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}
	id := uint(id_int)
	c.service.Delete(id).Send(w)
}