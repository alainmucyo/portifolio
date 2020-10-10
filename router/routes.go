package router

import (
	"github.com/alainmucyo/my_brand/config"
	"github.com/alainmucyo/my_brand/controller/articles"
	"github.com/alainmucyo/my_brand/controller/auth"
	"github.com/alainmucyo/my_brand/controller/queries"
	"github.com/alainmucyo/my_brand/middlewares"
	"github.com/alainmucyo/my_brand/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func Register() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(utils.MyFileSystem{http.Dir(config.STATIC_FOLDER)})))

	r.HandleFunc("/api/query", middlewares.IsAuth(queries.GetQuery)).Methods("GET")
	r.HandleFunc("/api/query", queries.CreateQuery).Methods("POST")

	r.HandleFunc("/api/article", articles.Index).Methods("GET")
	r.HandleFunc("/api/article/{id}", articles.Show).Methods("GET")
	r.HandleFunc("/api/article", middlewares.IsAuth(articles.Store)).Methods("POST")
	r.HandleFunc("/api/article/{id}", middlewares.IsAuth(articles.Delete)).Methods("DELETE")
	r.HandleFunc("/api/article/like/{article}", articles.Like).Methods("PUT")
	r.HandleFunc("/api/article/{id}", middlewares.IsAuth(articles.Update)).Methods("PUT")

	r.HandleFunc("/api/comment/{article}", articles.Comment).Methods("POST")

	r.HandleFunc("/api/auth/mock", middlewares.IsAuth(auth.MockUser))
	r.HandleFunc("/api/auth/login", auth.Login).Methods("POST")
	r.HandleFunc("/api/auth/details", middlewares.IsAuth(auth.UserDetails)).Methods("GET")
	return r
}
