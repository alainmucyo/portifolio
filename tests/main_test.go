package tests

import (
	"github.com/alainmucyo/my_brand/config"
	"github.com/alainmucyo/my_brand/src/database"
	"github.com/alainmucyo/my_brand/src/model"
	"github.com/alainmucyo/my_brand/src/router"
	"github.com/alainmucyo/my_brand/tests/article"
	"github.com/alainmucyo/my_brand/tests/auth"
	"github.com/alainmucyo/my_brand/tests/data"
	"github.com/alainmucyo/my_brand/tests/query"
	"github.com/joho/godotenv"
	"log"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading env variables!")
	}
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_DATABASE := os.Getenv("DB_TEST_DATABASE")
	DB_TYPE := os.Getenv("DB_TYPE")
	config.Connect(DB_PASSWORD, DB_USERNAME, DB_DATABASE, DB_TYPE)
	database.Migrate()
	emptyTables()
	ts := httptest.NewServer(router.Register())
	data.MainURL = ts.URL
	defer ts.Close()
	os.Exit(m.Run())
}

func TestBest(t *testing.T) {
	t.Run("Auth", auth.Auth)
	t.Run("Articles", article.Article)
	t.Run("Queries", query.Query)
}

func emptyTables() {
	config.Database.Where("1=1").Delete(&model.User{})
	config.Database.Where("1=1").Delete(&model.Query{})
}
