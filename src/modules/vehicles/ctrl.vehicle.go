package vehicles

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/interfaces"
	"github.com/AnggaArdhinata/backend-go/src/library"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type vehicles_ctrl struct {
	service interfaces.Vehicles_ServiceIF
}

func NewCtrl(data interfaces.Vehicles_ServiceIF) *vehicles_ctrl {
	return &vehicles_ctrl{data}
}

func (c *vehicles_ctrl) FindAll(w http.ResponseWriter, r *http.Request) {
	result := c.service.FindAll()
	result.Send(w)
	return

}

func (c *vehicles_ctrl) FindByRating(w http.ResponseWriter, r *http.Request) {
	result := c.service.FindByRating()
	result.Send(w)
	return

}

func (c *vehicles_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var data models.Vehicle
	var decoder = schema.NewDecoder()

	file, handler, err := r.FormFile("image")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	uploadImage, publicId, err := library.UploadImage("vehicle", file, handler)
	if err != nil {
		library.Response(err, 500, true)
		return
	}
	data.Image = uploadImage
	// fmt.Println(publicId)
	data.ImageUuid = publicId

	err = decoder.Decode(&data, r.Form)
	if err != nil {
		library.Response(err, 500, true)
		return
	}
	
	c.service.Add(&data).Send(w)
}

func (c *vehicles_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_int, err := strconv.Atoi(vars["id"])
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}
	id := uint(id_int)
	var data models.Vehicle
	json.NewDecoder(r.Body).Decode(&data)
	c.service.Update(id, &data).Send(w)
}

func (c *vehicles_ctrl) Delete(w  http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_int, err := strconv.Atoi(vars["id"])
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}
	id := uint(id_int)
	c.service.Delete(id).Send(w)
}