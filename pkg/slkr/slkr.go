package slkr

import (
	"github.com/nlopes/slack"
)

var Api *slack.Client

func Initialize(token string) {
	Api = slack.New(token)
}
