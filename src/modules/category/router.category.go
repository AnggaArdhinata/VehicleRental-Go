package category

import (
	"github.com/AnggaArdhinata/backend-go/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRt(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicle/category").Subrouter()

	repo := NewRepo(db)
	service := NewService(repo)
	ctrl := NewCtrl(service)

	route.HandleFunc("/", middleware.Handle(ctrl.FindAll, middleware.AuthMiddle("admin"))).Methods("GET")
	route.HandleFunc("/{id}", middleware.Handle(ctrl.FindById, middleware.AuthMiddle("admin"))).Methods("GET")
	route.HandleFunc("/", middleware.Handle(ctrl.Add, middleware.AuthMiddle("admin"))).Methods("POST")
	route.HandleFunc("/{id}", middleware.Handle(ctrl.Update, middleware.AuthMiddle("admin"))).Methods("PUT")
	route.HandleFunc("/{id}", middleware.Handle(ctrl.Delete, middleware.AuthMiddle("admin"))).Methods("DELETE")
}
