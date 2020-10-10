package router

import (
	"github.com/alainmucyo/my_brand/config"
	"github.com/alainmucyo/my_brand/controller/articles"
	"github.com/alainmucyo/my_brand/controller/queries"
	"github.com/alainmucyo/my_brand/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func Register() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(utils.MyFileSystem{http.Dir(config.STATIC_FOLDER)})))

	r.HandleFunc("/api/query", queries.GetQuery).Methods("GET")
	r.HandleFunc("/api/query", queries.CreateQuery).Methods("POST")

	r.HandleFunc("/api/article", articles.Index).Methods("GET")
	r.HandleFunc("/api/article/{id}", articles.Show).Methods("GET")
	r.HandleFunc("/api/article", articles.Store).Methods("POST")
	r.HandleFunc("/api/article/{id}", articles.Delete).Methods("DELETE")
	r.HandleFunc("/api/article/{id}", articles.Update).Methods("PUT")
	return r
}
