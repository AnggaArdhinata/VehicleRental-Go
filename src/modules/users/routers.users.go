package users

import (
	// "github.com/AnggaArdhinata/backend-go/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRt(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/user").Subrouter()

	repo := NewUserRepo(db)
	service := NewService(repo)
	ctrl := NewCtrl(service)

	//enable if you want running app with authorization
	// route.HandleFunc("/profile", middleware.Handle(ctrl.GetByIdCtx, middleware.AuthMiddle("user"))).Methods("GET")
	// route.HandleFunc("/{id}", middleware.Handle(ctrl.FindById, middleware.AuthMiddle("admin"))).Methods("GET")

	route.HandleFunc("/{id}", ctrl.FindById).Methods("GET")

	route.HandleFunc("", ctrl.FindAll).Methods("GET")
	route.HandleFunc("/user/{username}", ctrl.FindByUserName).Methods("GET")
	route.HandleFunc("/email/{email}", ctrl.FindByUserEmail).Methods("GET")
	route.HandleFunc("/", ctrl.Add).Methods("POST")
	route.HandleFunc("/{id}", ctrl.Update).Methods("PUT")
	route.HandleFunc("/{id}", ctrl.Delete).Methods("DELETE")
}