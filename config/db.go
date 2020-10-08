package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Database *gorm.DB

func Connect(DB_PASSWORD string, DB_USERNAME string, DB_DATABASE string, DB_TYPE string) {
	var err error
	dvn := DB_TYPE + "://" + DB_USERNAME + ":" + DB_PASSWORD + "@localhost/" + DB_DATABASE
	Database, err = gorm.Open(postgres.Open(dvn), &gorm.Config{})
	if err != nil {
		os.Exit(2)
	}
}
