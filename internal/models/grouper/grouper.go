package grouper

import (
	"github.com/jinzhu/gorm"
<<<<<<< HEAD
	"golend/internal/models/lendable"
=======
	"lendr/internal/models/lendable"
>>>>>>> 055efae1be042cd5ac2cc6e396bc45da4afe753e
)

type Grouper struct {
	gorm.Model
	Name string
	LenderId  uint `gorm:"index"`
	Access     string
	Lendables []lendable.Lendable
}
