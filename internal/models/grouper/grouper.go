package grouper

import (
	//_ "github.com/lib/pq" // TODO : consider postgres for "DISTINCT ON"
	"golend/internal/models/lendable"

	_ "github.com/go-sql-driver/mysql" // must be included with gorm
	"github.com/jinzhu/gorm"
)

type Grouper struct {
	gorm.Model
	Name               string
	LenderId           uint `gorm:"index"`
	Access             string
	Lendables          []lendable.Lendable
	TotalLendables     uint
	AvailableLendables uint
}
