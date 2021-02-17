package slack_app

import (
	"github.com/slack-go/slack"
	"log"
)

const (
	token     = "<>"
	channelId = "<>"
)

type SlackMessage struct {
	Title       string
	Text    string
	Pretext   string
	MsgOptionText string
}
/*
You can use http://davestevens.github.io/slack-message-builder/ to preview your slack message.
 */
func (sm *SlackMessage) PostMessageToChannel() error {
	api := slack.New(token)
	attachment := slack.Attachment{
		Title: sm.Title,
		Text: sm.Text,
		Pretext: sm.Pretext,
	}

	// Though it is not required and it can add a bit of latency in this function, but just to demonstrate how to get
	// a channel details, we use this code snippet.
	channel, err := api.GetConversationInfo(channelId, true)
	if err != nil {
		log.Fatalf("Error while fetching channel details: %s", err.Error())
	}
	log.Printf("Channel details, Name: %s, ChannelId: %s", channel.Name, channel.ID)

	// Now we have all details. Lets post the message in slack channel.
	resp,timestamp, err := api.PostMessage(
		channelId,
		slack.MsgOptionText(sm.MsgOptionText, true),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
		)
	if err != nil {
		log.Fatalf("Error while sending message to channel: %s", err.Error())
	}
	log.Printf("Response: %s", resp)
	log.Printf("Message sent on timestamp: %s", timestamp)
	log.Println("Slack message posted successfully")
	return nil
}