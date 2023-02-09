package reservation

import (
	"github.com/AnggaArdhinata/backend-go/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRt(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/reservation").Subrouter()

	repo := NewRepo(db)
	service := NewService(repo)
	ctrl := NewCtrl(service)

	route.HandleFunc("/user", middleware.Handle(ctrl.FindByUserID, middleware.AuthMiddle("user"))).Methods("GET")
	route.HandleFunc("/", ctrl.FindAll).Methods("GET")
	route.HandleFunc("/", middleware.Handle(ctrl.Add, middleware.AuthMiddle("user"))).Methods("POST")
	route.HandleFunc("/{id}", ctrl.Update).Methods("PUT")
	route.HandleFunc("/{id}", ctrl.Delete).Methods("DELETE")
}
