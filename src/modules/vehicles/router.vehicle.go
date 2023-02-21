package vehicles

import (
	// "github.com/AnggaArdhinata/backend-go/src/middleware" 
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRt(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicle").Subrouter()

	repo := NewRepo(db)
	service := NewService(repo)
	ctrl := NewCtrl(service)

	//* Enable if u want to enable bearer token authorization
	// route.HandleFunc("/", middleware.Handle(ctrl.Add, middleware.AuthMiddle("admin"))).Methods("POST")
	// route.HandleFunc("/{id}", middleware.Handle(ctrl.Update, middleware.AuthMiddle("admin"))).Methods("PUT")
	// route.HandleFunc("/{id}", middleware.Handle(ctrl.Delete, middleware.AuthMiddle("admin"))).Methods("DELETE")

	route.HandleFunc("/", ctrl.FindAll).Methods("GET")
	route.HandleFunc("/popular", ctrl.FindByRating).Methods("GET")

	route.HandleFunc("/", ctrl.Add).Methods("POST")
	route.HandleFunc("/{id}", ctrl.Update).Methods("PUT")
	route.HandleFunc("/{id}", ctrl.Delete).Methods("DELETE")

}
