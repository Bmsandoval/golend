package models

import (
	"lendr/pkg/db"
	"time"
)

type Lendable struct {
	TeamId   string `gorm:"primary_key"`
	Parent   string `gorm:"primary_key"`
	Name     string
	CurUsage string
	UpdateAt time.Time `gorm:"type:timestamp"`
}

func RemoveLendablesByLendr(lendrId string) {
	db.DB.Debug().Table("lendables").Where("lendr_id= ?", lendrId).Delete(&Lendable{})
}

func CreateLendablesTable() {
	db.DB.Debug().AutoMigrate(&Lendable{})
}

func DeleteLendablesTable() {
	db.DB.Debug().DropTableIfExists(&Lendable{})
}
