package lender

import (
	"golend/pkg/db"
	"log"

	"github.com/jinzhu/gorm"
)

/**
Create a new prospect with Id generated as defined by input function
*/
func CreateProspect(prospectIdGenerator func(int) string) string {
	slackState := ""
	for {
		// generate id
		slackState = prospectIdGenerator(20)
		// try to create prospect
		result := db.DB.Create(&Lender{TeamId: slackState})
		// prospect creation succeeds if id not in table
		if result.Error == nil {
			break
		}
		// otherwise just keep trying
		continue
	}
	return slackState
}

/**
Remove a prospect, perhaps they cancelled or ignored invite
*/
func RemoveProspect(prospectId string) {
	db.DB.Table("lenders").
		Scopes(ProspectById(prospectId)).
		Delete(&Lender{})
}

/**
After 0Auth flow complete, assign prospective lendee their own lender
*/
func ConvertProspect(prospectId string, teamId string, accessToken string, botAccessToken string) {
	result := db.DB.Table("lenders").
		Scopes(ProspectById(prospectId)).
		Updates(map[string]interface{}{
			"team_id":          teamId,
			"access_token":     accessToken,
			"bot_access_token": botAccessToken})

	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

/**
Check database for prospective lendee by id
*/
func ProspectExists(prospectId string) bool {
	var lendr = Lender{}
	result := db.DB.Table("lenders").
		Scopes(ProspectById(prospectId)).
		First(&lendr)
	return result.Error == nil
}

//*******************************
//*****    QUERY HELPERS    *****
//*******************************
// limit query results to prospective lendee's id
func ProspectById(prospectId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Where("team_id= ?", prospectId).
			Where("access_token = ?", "").
			Where("bot_access_token = ?", "")
	}
}
