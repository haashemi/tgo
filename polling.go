package tgo

import (
	"errors"
	"syscall"
	"time"
)

// StartPolling does an infinite GetUpdates with the timeout of the passed timeoutSeconds.
// allowedUpdates by default passes nothing and uses the telegram's default.
//
// see tgo.GetUpdate for more detailed information.
func (bot *Bot) StartPolling(timeoutSeconds int64, allowedUpdates ...string) error {
	var offset int64

	for {
		data, err := bot.GetUpdates(&GetUpdates{
			Offset:         offset, // Is there any better way to do this? open an issue/pull-request if you know. thx.
			Timeout:        timeoutSeconds,
			AllowedUpdates: allowedUpdates,
		})
		if err != nil {
			if errors.Is(err, syscall.ECONNRESET) {
				time.Sleep(time.Second / 2)
				continue
			}
			return err
		}

		for _, update := range data {
			offset = update.UpdateId + 1

			go func(update *Update) {
				if update.Message != nil && bot.sendAnswerIfAsked(update.Message) {
					return
				}

				for _, router := range bot.routers {
					if used := router.HandleUpdate(bot, update); used {
						return
					}
				}
			}(update)
		}
	}
}
