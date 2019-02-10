package controllers

import (
	"github.com/nlopes/slack"
	"lendr/internal/models"
	"lendr/pkg/slkr"
	"log"
	"net/http"
)

// ************************************
// Entry point for management commands
// ************************************
func ManagementIndex(w http.ResponseWriter, r *http.Request) {
	//// Save a copy of this request for debugging.
	//requestDump, err := httputil.DumpRequest(r, true)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(requestDump))
	//teamId := r.FormValue("team_id")
	//fmt.Println(teamId)



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

	lendr, _ := models.GetLendr(s.TeamID)

	//*************************
	//Setup slack api
	//*************************
	slackBotToken := lendr.BotAccessToken
	slkr.Initialize(slackBotToken)

	//w.WriteHeader(http.StatusOK)
	//res, err := w.Write([]byte("Alright.. alright. Just give me a sec"))
	//_, _ = res, err

	result, err := slkr.Api.PostEphemeral(s.ChannelID, s.UserID, slack.MsgOptionAttachments(
		slack.Attachment{
			CallbackID: "manage.main_selection",
			Actions: []slack.AttachmentAction{
				{
					Name: "create_action",
					Text: "Add Something",
					Type: "button",
					Value: "create",
				},
				{
					Name: "remove_action",
					Text: "Remove Something",
					Type: "button",
					Value: "remove",
				},
			},
		}))

	_ = result






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
