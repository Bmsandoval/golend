package models

import (
	_ "github.com/go-sql-driver/mysql" // must be included with gorm
	"github.com/jinzhu/gorm"
	"golend/pkg/db"
)

type Lendable struct {
	gorm.Model
	GrouperId uint `gorm:"index"`
	Name      string
	HeldBy    string
}

func MakeNewLendable( grouperId uint, name string) *gorm.DB {
	lendable := Lendable {
		Name: name,
		GrouperId: grouperId,
	}
	// TODO : handle the error if they insert a duplicate name
	result := db.DB.Create(&lendable)
	return result
}
