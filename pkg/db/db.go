package db

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Open(username string, password string, hostname string, database string) error {
	var err error
	DB, err = gorm.Open("mysql", username+":"+password+"@tcp("+hostname+")/"+database)
	return err
}

func Close() error {
	return DB.Close()
}
