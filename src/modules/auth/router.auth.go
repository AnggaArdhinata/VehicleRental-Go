package auth

import (
	"github.com/AnggaArdhinata/backend-go/src/modules/users"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRt(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/auth").Subrouter()

	repo := users.NewUserRepo(db)
	service := NewService(repo)
	ctrl := NewCtrl(service)

	route.HandleFunc("/login", ctrl.SignIn).Methods("POST")
	route.HandleFunc("/register", ctrl.Register).Methods("POST")

}
