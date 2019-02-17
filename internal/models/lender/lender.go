package lender

import (
	_ "github.com/go-sql-driver/mysql" // must be included with gorm
	"github.com/jinzhu/gorm"
	"golend/internal/models/grouper"
	"golend/pkg/db"
	//_ "github.com/kr/pretty"
)

type Lender struct {
	gorm.Model
	TeamId         string `gorm:"unique_index"`
	Groupers  []grouper.Grouper
	Admins         string
	AccessToken    string
	BotAccessToken string
}

func GetLender(teamId string) (Lender, error) {
	var lendr = Lender{}
	result := db.DB.Table("lenders").
		Where("team_id= ?", teamId).
		First(&lendr)
	return lendr, result.Error
}

