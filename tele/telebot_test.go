package tele

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	tb "github.com/tucnak/telebot"
)

func TestFiring(t *testing.T) {
	//Add key and ID to test
	t.Skip()
	var (
		apiKey       = ""
		chatID int64 = 0
	)
	bot, err := tb.NewBot(tb.Settings{
		Token:  apiKey,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	assert.NoError(t, err)
	chats := []*tb.Chat{{
		ID: chatID,
	}}
	tl := NewTele(bot, chats)
	for _, recipient := range tl.channels {
		_, err := tl.bot.Send(recipient, "lên tàu em ei")
		assert.NoError(t, err)
	}
}
