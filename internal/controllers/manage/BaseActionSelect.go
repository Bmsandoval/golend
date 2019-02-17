package manage

import (
	"fmt"
	"github.com/nlopes/slack"
	"golend/internal/models/lender"
	"golend/pkg/slkr"
	"net/http"
)


// ************************************
// Entry point for management commands
// ************************************
const baseActionSelectCallback string = "base_action_select"
func BaseActionSelect(requestValues slack.InteractionCallback, w http.ResponseWriter, r *http.Request) {
	//*************************
	//Setup slack api
	//*************************
	lendr, err := lender.GetLender(requestValues.Team.ID)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
	} else {
		w.WriteHeader(http.StatusAccepted)
	}
	slkr.Initialize(lendr.BotAccessToken)

	// Handle Management Callbacks
	if requestValues.Actions[0].Name == "create_action" &&
		requestValues.Actions[0].Value == "create" {

		err := slkr.Api.OpenDialog(requestValues.TriggerID, slack.Dialog{
			Title: "Creating New Lendable",
			SubmitLabel: "Submit",
			TriggerID: requestValues.TriggerID,
			CallbackID: "manage."+lendablesCreateDialogCallback,
			Elements: []slack.DialogElement{
				slack.TextInputElement{
					DialogInput: slack.DialogInput{
						Name:  "lendableName",
						Type:  slack.InputTypeText,
						Label: "Name that Lendable",
					},
					Hint: "How should people refer to your Lendable?",
				},
				slack.DialogInputSelect{
					DialogInput: slack.DialogInput{
						Type: slack.InputTypeSelect,
						Name: "groupStatus",
						Label: "Should this be grouped?",
					},
					Options: []slack.DialogSelectOption{
						{
							Label: "Group under something",
							Value: "isLendable",
						},
						{
							Label: "Make this a group",
							Value: "isGroup",
						},
					},
				},
				slack.DialogInputSelect{
					DialogInput: slack.DialogInput{
						Type: slack.InputTypeSelect,
						Name: "channelStatus",
						Label: "Who deserves access?",
					},
					Options: []slack.DialogSelectOption{
						{
							Label: "Everyone",
							Value: "everyone",
						},
						{
							Label: "Specific channel(s)",
							Value: "specificChannels",
						},
					},
				},
			},
		})
		if err != nil {
			fmt.Printf("%s", err)
			fmt.Println("")
		}
	}
	return
}

