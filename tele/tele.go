package tele

import (
	"fmt"

	tb "github.com/tucnak/telebot"

	"github.com/cyberliem/volume-surge-alarm/common"
)

//Tele abstract a tele bot
type Tele struct {
	bot      *tb.Bot
	channels []*tb.Chat
}

//Fire send the alarm message
func (tl *Tele) Fire(cp common.ChangeCriteria) error {
	if len(cp.Scores) == 0 {
		return nil
	}
	var msg string
	msg += fmt.Sprintf("In period of %s, these pair's change has crossed the threholds: \n", cp.Duration.String())
	for _, data := range cp.Scores {
		msg += fmt.Sprintf("%s : %.2f%% \n", data.Pair, data.Percent)
	}
	for _, recipient := range tl.channels {
		_, err := tl.bot.Send(recipient, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

//NewTele return a tele instance
func NewTele(bot *tb.Bot, channels []*tb.Chat) *Tele {
	return &Tele{
		bot:      bot,
		channels: channels,
	}
}
