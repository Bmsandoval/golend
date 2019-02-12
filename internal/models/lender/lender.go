package lender

import (
	"github.com/jinzhu/gorm"
	"golend/internal/models/grouper"
	"golend/pkg/db"
	"log"

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

func CreateProspect(rando func(int) string) string {
	slackState := rando(20)
	result := db.DB.Debug().Create(&Lender{TeamId: slackState})
	for {
		if result.Error == nil {
			break
		}
		slackState := rando(20)
		result = db.DB.Debug().Create(&Lender{TeamId: slackState})
	}
	return slackState
}

func RemoveProspect(prospectId string) {
	db.DB.Debug().
		Table("lenders").
		Where("team_id= ?", prospectId).
		Where("access_token = ?", "").
		Where("bot_access_token = ?", "").
		Delete(&Lender{})
}

func ConvertProspect(prospectId string, teamId string, accessToken string, botAccessToken string) {
	result := db.DB.Debug().
		Table("lenders").
		Where("team_id= ?", prospectId).
		Where("access_token = ?", "").
		Where("bot_access_token = ?", "").
		Updates(map[string]interface{}{
			"team_id":         teamId,
			"access_token":     accessToken,
			"bot_access_token": botAccessToken})

	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

func GetLender(teamId string) (Lender, error) {
	var lendr = Lender{}
	result := db.DB.Debug().
		Table("lenders").
		Where("team_id= ?", teamId).
		First(&lendr)
	return lendr, result.Error
}

func ProspectExists(prospectId string) bool {
	var lendr = Lender{}
	result := db.DB.Debug().
		Table("lenders").
		Where("team_id= ?", prospectId).
		Where("access_token = ?", "").
		Where("bot_access_token = ?", "").
		First(&lendr)
	return result.Error == nil
}

