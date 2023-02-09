package routers

import (
	"github.com/AnggaArdhinata/backend-go/src/databases/orm"
	"github.com/AnggaArdhinata/backend-go/src/modules/auth"
	"github.com/AnggaArdhinata/backend-go/src/modules/category"
	"github.com/AnggaArdhinata/backend-go/src/modules/history"
	"github.com/AnggaArdhinata/backend-go/src/modules/reservation"
	"github.com/AnggaArdhinata/backend-go/src/modules/users"
	"github.com/AnggaArdhinata/backend-go/src/modules/vehicles"
	"github.com/gorilla/mux"
)

func NewApp() (*mux.Router, error) {
	mainRoute := mux.NewRouter()
	
	db, err := orm.NewDb()
	if err != nil {
		return nil, err
	}
	
	users.NewRt(mainRoute, db)
	category.NewRt(mainRoute, db)
	vehicles.NewRt(mainRoute, db)
	reservation.NewRt(mainRoute, db)
	history.NewRt(mainRoute, db)
	auth.NewRt(mainRoute, db)
	
	return mainRoute, nil

}