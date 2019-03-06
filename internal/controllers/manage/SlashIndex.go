package manage

import (
	"golend/internal/models"
	"golend/pkg/slkr"
	"log"
	"net/http"
	"strings"

	"github.com/nlopes/slack"
)

// ************************************
// Entry point for management commands
// lendr-manage create,update,delete,config
// ************************************
func SlashIndex(w http.ResponseWriter, r *http.Request) {
	// Parse request
	v, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	// Get Team
	// TODO - Error if team not found
	lendr, _ := models.FindLenderByTeam(v.TeamID)

	//*************************
	//Setup slack api
	//*************************
	slackBotToken := lendr.BotAccessToken
	slkr.Initialize(slackBotToken)

	var callbackID string
	switch strings.ToLower(v.Text) {
	case lendablesCreateCallback:
		callbackID = lendablesCreateCallback
	case "update":
		// TODO : Config not Implemented
		w.WriteHeader(http.StatusNotImplemented)
		log.Fatal("Config not implemented")
		return
	case "delete":
		// TODO : Config not Implemented
		w.WriteHeader(http.StatusNotImplemented)
		log.Fatal("Config not implemented")
		return
	case "config":
		// TODO : Config not Implemented
		w.WriteHeader(http.StatusNotImplemented)
		log.Fatal("Config not implemented")
		return
	default:
		slkr.Api.PostEphemeral(v.ChannelID, v.UserID, ManagementInitialOptions())
		return
	}

	result, err := slkr.Api.PostEphemeral(v.ChannelID, v.UserID, ManagementSelectRelationship(callbackID))
	_ = result
	return
}

func ManagementInitialOptions () slack.MsgOption{
	return slack.MsgOptionAttachments(
		slack.Attachment{
			Text: "What should we do? Create perhaps?",
		})
}

func ManagementSelectRelationship(callbackID string) slack.MsgOption{
	return slack.MsgOptionAttachments(
		slack.Attachment{
			Text: "What should we do? Try again like create,update,delete",
			CallbackID: "manage." + callbackID,
			Actions: []slack.AttachmentAction{
				{
					Name:  "grouper",
					Text:  "Make Group",
					Type:  "button",
					Value: "grouper",
				},
				{
					Name:  "lendable",
					Text:  "Group Under..",
					Type:  "button",
					Value: "lendable",
				},
			},
		})
}
