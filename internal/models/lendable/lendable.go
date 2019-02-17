package lendable

import (
	_ "github.com/go-sql-driver/mysql" // must be included with gorm
	"github.com/jinzhu/gorm"
)

type Lendable struct {
	gorm.Model
	GrouperId uint `gorm:"index"`
	Name      string
	HeldBy    string
}
