package models

import (
	"lendr/pkg/db"
	"log"

	_ "github.com/kr/pretty"
)

type Lendr struct {
	LendrId        string `gorm:"primary_key"`
	Admins         string
	AccessToken    string
	BotAccessToken string
	CreatedAt      []uint8 `gorm:"type:timestamp"`
}

func MakeProspectiveLendr(rando func(int) string) string {
	slackState := rando(20)
	result := db.DB.Debug().Create(&Lendr{LendrId: slackState})
	for {
		if result.Error == nil {
			break
		}
		slackState := rando(20)
		result = db.DB.Debug().Create(&Lendr{LendrId: slackState})
	}
	return slackState
}

func DeleteProspectiveLendr(state string) {
	db.DB.Debug().Table("lendrs").Where("lendr_id= ?", state).Delete(&Lendr{})
}

func ConvertProspectiveLendr(lendrId string, teamId string, accessToken string, botAccessToken string) {
	result := db.DB.Debug().
		Table("lendrs").
		Where("lendr_id= ?", lendrId).
		Updates(map[string]interface{}{
			"lendr_id":         teamId,
			"access_token":     accessToken,
			"bot_access_token": botAccessToken})

	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

func GetLendr(lendrId string) (Lendr, error) {
	var lendr = Lendr{}
	result := db.DB.Debug().Table("lendrs").Where("lendr_id= ?", lendrId).First(&lendr)
	return lendr, result.Error
}

func LendrExists(lendrId string) bool {
	_, err := GetLendr(lendrId)
	return err == nil
}

func AddNewLendr(lendrId string) {
	db.DB.Debug().Create(&Lendr{LendrId: lendrId})
}

func RemoveExistingLendr(lendrId string) {
	db.DB.Debug().Table("lendrs").Where("lendr_id= ?", lendrId).Delete(&Lendr{})
}

func CreateLendrsTable() {
	db.DB.Debug().AutoMigrate(&Lendr{})
}

func DeleteLendrsTable() {
	db.DB.Debug().DropTableIfExists(&Lendr{})
}
