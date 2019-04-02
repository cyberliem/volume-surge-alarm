package tele

import (
	"errors"
	"time"

	tb "github.com/tucnak/telebot"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

const (
	teleBotAPIFlag = "telebot-api-key"
	chatIDFlag     = "telebot-chat-ids"
)

//NewCliFlags return cli flags to configure cex-trade client
func NewCliFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   teleBotAPIFlag,
			Usage:  "API key for telebot",
			EnvVar: "TELEBOT_API_KEY",
		},
		cli.Int64SliceFlag{
			Name:   teleBotAPIFlag,
			Usage:  `chat IDs for the telebot, can be mutiple: example: --telebot-chat-ids={0, 1}`,
			EnvVar: "TELEBOT_CHAT_IDS",
		},
	}
}

//NewTeleFromContext return binance client
func NewTeleFromContext(c *cli.Context, sugar *zap.SugaredLogger) (*Tele, error) {
	var (
		apiKey string
		chats  []*tb.Chat
	)
	apiKey = c.String(teleBotAPIFlag)
	if len(apiKey) == 0 {
		return nil, errors.New("no api key is provided for tele bot")
	}

	bot, err := tb.NewBot(tb.Settings{
		Token:  apiKey,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}
	chatIDs := c.Int64Slice(chatIDFlag)
	if len(chatIDs) == 0 {
		return nil, errors.New("no chat id is provided for tele bot")
	}
	for _, id := range chatIDs {
		chat := tb.Chat{
			ID: id,
		}
		chats = append(chats, &chat)
	}
	return NewTele(bot, chats), nil
}
