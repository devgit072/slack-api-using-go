package main

import (
	slack_app "github.com/devgit072/slack-api-using-go/slack-app"
	"log"
)

func main() {
	sendMessageToChannel()
	sendMessageToUser()
}
func sendMessageToChannel() {
	sm := slack_app.SlackMessage{
		Title:         "Simple Demonstration, how to send slack message in slack channel using Go",
		Text:          "Some Text",
		Pretext:       "Some pretext",
		MsgOptionText: "Slack has nice interface to send message",
		ChannelId:     "C01MZ6570MD",
	}

	if err := sm.PostMessageToChannel(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

func sendMessageToUser() {
	sm := slack_app.SlackMessage{
		Title:         "Simple Demonstration, how to send slack message in slack channel using Go",
		Text:          "Some Text",
		Pretext:       "Some pretext",
		MsgOptionText: "Slack has nice interface to send message",
		UserEmail:     "devraj@cohesity.com",
	}

	if err := sm.SendDirectMessage(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
