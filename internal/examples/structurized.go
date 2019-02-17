package examples

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/kr/pretty"
	"github.com/nlopes/slack"
)

func TestStructurized(w http.ResponseWriter, r *http.Request) {
	dialog := slack.Dialog{
		Title:       "Creating New Lendable",
		SubmitLabel: "Submit",
		TriggerID:   "faketrigger",
		CallbackID:  "whatevs",
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
					Type:  slack.InputTypeSelect,
					Name:  "groupStatus",
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
					Type:  slack.InputTypeSelect,
					Name:  "channelStatus",
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
	}
	jsonized, err := json.Marshal(dialog)
	_ = err
	fmt.Println(jsonized)

}
