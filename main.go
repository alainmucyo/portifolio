package main

import (
	"github.com/alainmucyo/my_brand/config"
	"github.com/alainmucyo/my_brand/database"
	"github.com/alainmucyo/my_brand/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variables!")
	}
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_DATABASE := os.Getenv("DB_DATABASE")
	DB_TYPE := os.Getenv("DB_TYPE")
	config.Connect(DB_PASSWORD, DB_USERNAME, DB_DATABASE, DB_TYPE)
	PORT := os.Getenv("PORT")
	URL := os.Getenv("URL")
	database.Migrate()
	println("Server started at " + URL + ":" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router.Register()))
}
