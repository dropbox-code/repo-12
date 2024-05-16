package aws

import (
	"testing"

	"github.com/innogames/slack-bot/v2/bot"
	"github.com/innogames/slack-bot/v2/bot/config"
	"github.com/innogames/slack-bot/v2/bot/msg"
	"github.com/innogames/slack-bot/v2/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetCommands(t *testing.T) {
	slackClient := mocks.NewSlackClient(t)
	base := bot.BaseCommand{SlackClient: slackClient}

	t.Run("AWS is not active", func(t *testing.T) {
		cfg := config.Aws{}
		commands := GetCommands(cfg, base)
		assert.Equal(t, 0, commands.Count())
	})

	t.Run("AWS is active", func(t *testing.T) {
		cfg := config.Aws{
			Enabled: true,
			CloudFront: []config.AwsCfDistribution{
				{ID: "id", Name: "name"},
			},
		}
		commands := GetCommands(cfg, base)
		assert.Equal(t, 2, commands.Count())

		// test help
		help := commands.GetHelp()
		assert.Len(t, help, 2)

		// list the CF
		message := msg.Message{}
		message.Text = "aws cf list"
		mocks.AssertSlackBlocks(t, slackClient, message, `[{"type":"section","text":{"type":"mrkdwn","text":"\"id\": name\n"}}]`)

		actual := commands.Run(message)
		assert.True(t, actual)
	})
}
