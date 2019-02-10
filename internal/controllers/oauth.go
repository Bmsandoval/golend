package controllers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"lendr/internal/models"
	"log"
	"net/http"
	"os"
)

type slackBasicStruct struct {
	AccessToken string `json:"access_token"`
	TeamId      string `json:"team_id"`
	Bot         slackBotStruct
}
type slackBotStruct struct {
	BotAccessToken string `json:"bot_access_token"`
}

func CollectAuthToken(_ http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	err := r.FormValue("error")
	if err == "access_denied" {
		models.DeleteProspectiveLendr(state)
		return
	}
	tempCode := r.FormValue("code")
	lendrExists := models.LendrExists(state)
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
			var t slackBasicStruct
			err := decoder.Decode(&t)
			if err != nil {
				log.Fatal(err)
			} else {
				models.ConvertProspectiveLendr(state, t.TeamId, t.AccessToken, t.Bot.BotAccessToken)
			}
		}
	}
}

func InitializeTeamSetupProcedure(w http.ResponseWriter, r *http.Request) {
	slackClientID := os.Getenv("SLACK_CLIENT_ID")
	slackState := models.MakeProspectiveLendr(SecureRandomAlphaString)
	url := "https://slack.com/oauth/authorize?" +
		"client_id=" + slackClientID + "&" +
		"scope=chat:write:bot bot commands" + "&" +
		"state=" + slackState
	http.Redirect(w, r, url, 302)
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 52 possibilities
	letterIdxBits = 6                                                      // 6 bits to represent 64 possibilities / indexes
	letterIdxMask = 1<<letterIdxBits - 1                                   // All 1-bits, as many as letterIdxBits
)

func SecureRandomAlphaString(length int) string {

	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = SecureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & letterIdxMask); idx < len(letterBytes) {
			result[i] = letterBytes[idx]
			i++
		}
	}

	return string(result)
}

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func SecureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}
	return randomBytes
}
