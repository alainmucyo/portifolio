package router

import (
	"github.com/alainmucyo/my_brand/controller"
	"github.com/gorilla/mux"
)

func Register() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/query", controller.GetQuery).Methods("GET")
	r.HandleFunc("/api/query", controller.CreateQuery).Methods("POST")
	return r
}
