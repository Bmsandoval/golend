package slkr

import (
	"github.com/nlopes/slack"
	"hash/fnv"
)

var Api *slack.Client

func Initialize(token string) {
	Api = slack.New(token)
}

func SendError(channelId string, userId string, msg string) {
	result, _ := Api.PostEphemeral(channelId, userId, slack.MsgOptionAttachments(
		slack.Attachment{
			Title: "Whoops, something went wrong.",
			Text: msg,
		}))

	_ = result
	return

}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}