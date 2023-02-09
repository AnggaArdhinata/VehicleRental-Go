package history

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRt(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/history").Subrouter()

	repo := NewRepo(db)
	service := NewService(repo)
	ctrl := NewCtrl(service)

	// route.HandleFunc("/join", ctrl.JoinWithRes).Methods("GET")
	route.HandleFunc("/", ctrl.FindAll).Methods("GET")
	route.HandleFunc("/", ctrl.Add).Methods("POST")
	route.HandleFunc("/{id}", ctrl.Update).Methods("PUT")
	route.HandleFunc("/{id}", ctrl.Delete).Methods("DELETE")
}
