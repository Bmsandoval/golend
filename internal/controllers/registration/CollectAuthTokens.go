package registration

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"lendr/internal/models/lender"
)

func CollectAuthTokens(_ http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	err := r.FormValue("error")
	if err == "access_denied" {
		lender.RemoveProspect(state)
		return
	}
	tempCode := r.FormValue("code")
	lendrExists := lender.ProspectExists(state)
	if lendrExists {
		slackClientID := os.Getenv("SLACK_CLIENT_ID")
		slackClientSecret := os.Getenv("SLACK_CLIENT_SECRET")
		url := "https://slack.com/api/oauth.access?" +
			"client_id=" + slackClientID + "&" +
			"client_secret=" + slackClientSecret + "&" +
			"code=" + tempCode
		response, err := http.Post(url, "application/json", nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			decoder := json.NewDecoder(response.Body)
			var t struct {
				AccessToken string `json:"access_token"`
				TeamId      string `json:"team_id"`
				Bot         struct {
					BotAccessToken string `json:"bot_access_token"`
				}
			}
			err := decoder.Decode(&t)
			if err != nil {
				log.Fatal(err)
			} else {
				lender.ConvertProspect(state, t.TeamId, t.AccessToken, t.Bot.BotAccessToken)
			}
		}
	}
}
