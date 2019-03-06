package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // must be included with gorm
	"github.com/jinzhu/gorm"
	"github.com/nlopes/slack"
	"golend/pkg/db"
)

type Grouper struct {
	gorm.Model
	Name               string
	LenderId           uint `gorm:"index"`
	Access             string
	TotalLendables     uint
	AvailableLendables uint
	// Relation only, no column in DB
	Lendables          []Lendable
}

func GrouperToSelectables(groupers []Grouper) []slack.DialogSelectOption {
	var selectables = make([]slack.DialogSelectOption, len(groupers))
	for i, grouper := range groupers {
		selectables[i] = slack.DialogSelectOption{
			Label: grouper.Name,
			Value: fmt.Sprint(grouper.ID),
		}
	}
	return selectables
}

func MakeNewGrouper (lendrId uint, name string) *Grouper {
	grouper := Grouper{
		LenderId: lendrId,
		Name: name,
		TotalLendables: 1,
		AvailableLendables: 1,
	}
	result := db.DB.Create(&grouper)
	// TODO : handle the result.Error if they insert a duplicate name
	_ = result
	return &grouper
}

func FindGroupersByLender(lenderId uint) []Grouper {
	var groupers []Grouper
	db.DB.Table("groupers").
		Select("id, name").Scopes(
		ByLender(lenderId)).
		Limit(100).
		Find(&groupers)
	return groupers
}

//*******************************
//*****    QUERY HELPERS    *****
//*******************************
// limit query results to prospective lendee's id
func ByLender(lenderId uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("lender_id = ?", lenderId)
	}
}
