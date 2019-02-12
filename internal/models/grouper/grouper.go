package grouper

import (
	"github.com/jinzhu/gorm"
	"golend/internal/models/lendable"
)

type Grouper struct {
	gorm.Model
	Name string
	LenderId  uint `gorm:"index"`
	Access     string
	Lendables []lendable.Lendable
}
