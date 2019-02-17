package manage

import (
	"golend/internal/models/lender"
	"golend/pkg/slkr"
	"log"
	"net/http"

	"github.com/nlopes/slack"
)

// ************************************
// Entry point for management commands
// ************************************
func Index(w http.ResponseWriter, r *http.Request) {
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	//if !s.ValidateToken(os.Getenv("SLACK_SIGNING_SECRET")) {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}

	lendr, _ := lender.GetLender(s.TeamID)

	//*************************
	//Setup slack api
	//*************************
	slackBotToken := lendr.BotAccessToken
	slkr.Initialize(slackBotToken)
	result, err := slkr.Api.PostEphemeral(s.ChannelID, s.UserID, slack.MsgOptionAttachments(
		slack.Attachment{
			CallbackID: "manage." + baseActionSelectCallback,
			Actions: []slack.AttachmentAction{
				{
					Name:  "create_action",
					Text:  "Add Something",
					Type:  "button",
					Value: "create",
				},
				{
					Name:  "remove_action",
					Text:  "Remove Something",
					Type:  "button",
					Value: "remove",
				},
			},
		}))

	_ = result
	return

	//t, err := template.New("test").Parse(view.ManageBaseForm)
	//if err != nil {
	//	log.Print(err)
	//	return
	//}
	//
	//type Inventory struct {
	//	Material string
	//	Count    uint
	//}
	//sweaters := Inventory{"wool", 17}
	//
	//w.Header().Add("Content-Type", "application/json")
	//// TODO : make next line work
	//// t.Delims("<<", ">>")
	//err = t.Execute(w, sweaters) //, config)
	//if err != nil {
	//	log.Print("execute: ", err)
	//	return
	//}

}
