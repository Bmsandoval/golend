package lendable

import (
	"github.com/jinzhu/gorm"
)

type Lendable struct {
	gorm.Model
	GrouperId     uint `gorm:"index"`
	Name   string
	HeldBy         string
}
