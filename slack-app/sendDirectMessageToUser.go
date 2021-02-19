package slack_app

import (
	"github.com/slack-go/slack"
	"log"
)

func (sm *SlackMessage) SendDirectMessage() error {
	api := slack.New(token)
	attachment := slack.Attachment{
		Title:   sm.Title,
		Text:    sm.Text,
		Pretext: sm.Pretext,
	}
	user, err := api.GetUserByEmail(sm.UserEmail)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
		return err
	}
	log.Printf("User details: %s", user.Name)
	// now send message
	openConvParam := slack.OpenConversationParameters{
		ReturnIM: true,
		Users:    []string{user.ID},
	}
	ch, _, _, err := api.OpenConversation(&openConvParam)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
		return err
	}
	// Now we have all details. Lets post the message in slack channel.
	resp, timestamp, err := api.PostMessage(
		ch.ID,
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
