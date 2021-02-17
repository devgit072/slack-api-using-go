package main

import (
	slack_app "github.com/devgit072/slack-api-using-go/slack-app"
	"log"
)

func main() {
	sm := slack_app.SlackMessage{
		Title:         "Simple Demonstration, how to send slack message in slack channel using Go",
		Text:          "Some Text",
		Pretext:       "Some pretext",
		MsgOptionText: "Slack has nice interface to send message",
	}

	if err := sm.PostMessageToChannel(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
