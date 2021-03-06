package database

import (
	"github.com/alainmucyo/my_brand/config"
	"github.com/alainmucyo/my_brand/src/model"
)

func Migrate() {
	_ = config.Database.AutoMigrate(model.Query{}, model.Article{}, model.Comment{}, model.User{})
}
