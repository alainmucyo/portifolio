package database

import "github.com/alainmucyo/my_brand/config"
import "github.com/alainmucyo/my_brand/model"

func Migrate() {
	_ = config.Database.AutoMigrate(model.Query{})
}
