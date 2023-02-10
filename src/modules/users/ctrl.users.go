package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/interfaces"
	"github.com/AnggaArdhinata/backend-go/src/library"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type users_ctrl struct {
	service interfaces.User_ServiceIF
}

func NewCtrl(data interfaces.User_ServiceIF) *users_ctrl {
	return &users_ctrl{data}
}

func (c *users_ctrl) GetByIdCtx(w http.ResponseWriter, r *http.Request) {
	payload_id := r.Context().Value("user")
	c.service.FindById(payload_id.(uint)).Send(w)
	return
}

func (c *users_ctrl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)["id"]
	id_int, err := strconv.Atoi(vars)
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}
	id := uint(id_int)
	result := c.service.FindById(id)
	result.Send(w)
	return
}

func (c *users_ctrl) FindByUserEmail(w http.ResponseWriter, r *http.Request) {
	email_var := mux.Vars(r)["email"]
	result := c.service.FindByUserEmail(email_var)
	result.Send(w)
	return
}


func (c *users_ctrl) FindByUserName(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	result := c.service.FindByUserName(username)
	result.Send(w)
	return
}

func (c *users_ctrl) FindAll(w http.ResponseWriter, r *http.Request) {
	result := c.service.FindAll()
	result.Send(w)
	return
}

func (c *users_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var data models.User
	var decoder = schema.NewDecoder()

	file, handler, err := r.FormFile("image")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	uploadImage, publicId, err := library.UploadImage(file, handler)
	if err != nil {
		library.Response(err, 500, true)
		return
	}
	data.Image = uploadImage
	data.ImgId = publicId

	err = decoder.Decode(&data, r.Form)
	if err != nil {
		library.Response(err, 500, true)
		return
	}
	_, err = govalidator.ValidateStruct(data)
	if err != nil {
		library.Response(err, 500, true).Send(w)
		return
	}
	c.service.Add(&data).Send(w)

}

func (c *users_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_int, err := strconv.Atoi(vars["id"])
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}
	id := uint(id_int)
	var data models.User
	json.NewDecoder(r.Body).Decode(&data)
	c.service.Update(id, &data).Send(w)
}

func (c *users_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_int, err := strconv.Atoi(vars["id"])
	if err != nil {
		library.Response(err.Error(), 500, true).Send(w)
		return
	}
	id := uint(id_int)
	c.service.Delete(id).Send(w)
}
