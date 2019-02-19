package models

import (
	"fmt"
	"golend/pkg/db"
	"hash/fnv"
	"log"
	"regexp"

	_ "github.com/go-sql-driver/mysql" // must be included with gorm
	"github.com/jinzhu/gorm"
	//_ "github.com/kr/pretty"
)

type Lender struct {
	gorm.Model
	TeamId         string `gorm:"unique_index"`
	Admins         string
	AccessToken    string
	BotAccessToken string
	// Relation only, no column in DB
	Groupers       []Grouper
}

func (lendr Lender) HasAdmin(userId string) (bool, error) {
	// if Admins empty, then everyone's an admin
	var err error
	if lendr.Admins == "" {
		return true, nil
	}
	// Otherwise TRUE only if userId in Admins
	matched, err := regexp.MatchString(userId, lendr.Admins)
	if err != nil {
		log.Fatal(err.Error())
		return false, fmt.Errorf("Unexpected error")
	}
	if ! matched && err == nil{
		err = fmt.Errorf("User %s not an admin of Team %s",
			userId, lendr.TeamId)
	}
	return matched, err
}

func SearchLendersAsAdmin(teamId string, userId string) (Lender, error) {
	var lendr = Lender{}
	result := db.DB.Table("lenders").Scopes(
		ForTeam(teamId),
		WhileAdmin(userId)).
		First(&lendr)
	if result.Error != nil{
		return lendr, fmt.Errorf("no accessible team data found")
	}
	return lendr, nil
}

func (lender Lender) GetHash() string {
	lenderId := fmt.Sprint(lender.ID)
	h := fnv.New32a()
	h.Write([]byte(lenderId))
	h.Write([]byte(lender.TeamId))
	return fmt.Sprint(h.Sum32())
}


func (lender Lender) ValidateHash(hashedVal string) bool {
	if lender.GetHash() != hashedVal {
		log.Fatal("Invalid hash detected for lender " + fmt.Sprint(lender.ID))
	}
	return lender.GetHash() == hashedVal
}

func FindLenderByTeam(teamId string) (Lender, error) {
	var lendr = Lender{}
	result := db.DB.Table("lenders").
		Scopes(ForTeam(teamId)).
		First(&lendr)
	if result.Error != nil {
		return lendr, fmt.Errorf("Team %s not found in database", teamId)
	}
	return lendr, nil
}

//*******************************
//*****    QUERY HELPERS    *****
//*******************************
func ForTeam(teamId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("team_id = ?", teamId)
	}
}
// limit query results to prospective lendee's id
func WhileAdmin(userId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("admins = ? OR admins LIKE ?",
			"",
			"%"+userId+"%")
	}
}
