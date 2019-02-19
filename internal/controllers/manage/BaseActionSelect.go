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
	adminExists, err := lendr.HasAdmin(requestValues.User.ID)
	if  ! adminExists{
		slkr.SendError(requestValues.Channel.ID, requestValues.User.ID, err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//*************************
	// SWITCH: Callbacks
	// selects proper slack.Dialog
	//*************************
	switch requestValues.Actions[0].Value {
	case "create":
		groupers := models.FindGroupersByLender(lendr.ID)
		var selectables []slack.DialogSelectOption
		if len(groupers) == 0 {
			selectables = []slack.DialogSelectOption {
				{
					Label: "First time? I've gotcha.",
					Value: "isGroup",
				},
			}
		} else { // len(groupers) > 0
			selectables = make([]slack.DialogSelectOption, len(groupers))
			for i, grouper := range groupers {
				selectables[i] = slack.DialogSelectOption{
					Label: grouper.Name,
					Value: fmt.Sprint(grouper.ID),
				}
			}
		}
		dialog = LendablesCreationDialog(requestValues.TriggerID, lendr.GetHash(), selectables)

	case "update":
		// TODO - handle update requests
		w.WriteHeader(http.StatusNotImplemented)
		return
	case "delete":
		// TODO - handle delete requests
		w.WriteHeader(http.StatusNotImplemented)
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

func LendablesCreationDialog(triggerId string, state string, selectables []slack.DialogSelectOption) slack.Dialog {
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
