package slack_app

const (
	token = "<private_token>"
)

type SlackMessage struct {
	UserEmail     string
	ChannelId     string
	Title         string
	Text          string
	Pretext       string
	MsgOptionText string
}
