package manage

import (
	"fmt"
	"golend/internal/models"
	"golend/pkg/slkr"
	"log"
	"net/http"

	"github.com/nlopes/slack"
)

// ************************************
// Entry point for management commands
// ************************************
const baseActionSelectCallback string = "base_action_select"

func BaseActionSelect(requestValues slack.InteractionCallback, w http.ResponseWriter, r *http.Request) {
	var dialog slack.Dialog
	var err error
	teamId := requestValues.Team.ID
	channelId := requestValues.Channel.ID
	userId := requestValues.User.ID
	//*************************
	//Setup slack api
	//*************************
	lendr, err := models.FindLenderByTeam(teamId)
	if err != nil {
		// team must exist in database
		log.Fatal(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	slkr.Initialize(lendr.BotAccessToken)

	//*************************
	// Only allow Admins into management section
	//*************************
	adminExists, err := lendr.HasAdmin(userId)
	if  ! adminExists{
		slkr.SendError(channelId, userId, err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//*************************
	// SWITCH: Callbacks
	// selects proper slack.Dialog
	//*************************
	groupers := models.FindGroupersByLender(lendr.ID)
	switch requestValues.Actions[0].Value {
	case "create":
		var selectables []slack.DialogSelectOption
		if len(groupers) == 0 {
			selectables = []slack.DialogSelectOption {
				{
					Label: "First time? I've gotcha.",
					Value: "isGroup",
				},
			}
		} else { // len(groupers) > 0
			selectables = models.GrouperToSelectables(groupers)
		}
		dialog = LendablesCreationView(requestValues.TriggerID, lendr.GetHash(), selectables)

	case "update":
		var selectables []slack.DialogSelectOption
		if len(groupers) == 0 {
			slkr.Api.PostEphemeral(channelId, userId, ManagementSuggestCreateView())
			return
		} else {
			selectables = models.GrouperToSelectables(groupers)
			_ = selectables
		}
		// TODO - handle update requests
		w.WriteHeader(http.StatusNotImplemented)
		return
	case "delete":
		// TODO - handle delete requests
		w.WriteHeader(http.StatusNotImplemented)
		return
	case "no_action":
		// They just didn't want to do anything. Not an error
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Open selected dialog and check errors
	err = slkr.Api.OpenDialog(requestValues.TriggerID, dialog)

	if err != nil {
		fmt.Printf("%s", err)
		fmt.Println("")
	}
	return
}

func LendablesCreationView(triggerId string, state string, selectables []slack.DialogSelectOption) slack.Dialog {
	return slack.Dialog{
		Title:       "Creating New Lendable",
		SubmitLabel: "Submit",
		State: state,
		TriggerID:   triggerId,
		CallbackID:  "manage." + lendablesCreateDialogCallback,
		Elements: []slack.DialogElement {
			slack.TextInputElement {
				DialogInput: slack.DialogInput {
					Name:  "lendableName",
					Type:  slack.InputTypeText,
					Label: "Name that Lendable",
				},
				Hint: "How should people refer to your Lendable?",
			},
			slack.DialogInputSelect {
				DialogInput: slack.DialogInput {
					Type:  slack.InputTypeSelect,
					Name:  "groupStatus",
					Label: "Should this be grouped?",
				},
				OptionGroups: []slack.DialogOptionGroup {
					{
						Label: "Probably this one:",
						Options:
							[]slack.DialogSelectOption {
								{
									Label: "Make this a group",
									Value: "isGroup",
								},
							},
					},
					{
						Label: "Attach to Group:",
						Options:
							selectables,
					},
				},
			},
		},
	}
}

func ManagementSuggestCreateView () slack.MsgOption{
	return slack.MsgOptionAttachments(
		slack.Attachment{
			Text: "Weird, nothing there. Wanna make something?",
			CallbackID: "manage." + baseActionSelectCallback,
			Actions: []slack.AttachmentAction{
				{
					Name:  "create_action",
					Text:  "Add Something",
					Type:  "button",
					Value: "create",
				},
				{
					Name:  "no_action",
					Text:  "Nope.",
					Type:  "button",
					Value: "no_action",
				},
			},
		})
}
