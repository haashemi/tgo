package filters

import (
	"strings"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/tg"
)

// IsPrivate checks if the message (and only message) is inside the private chat.
func IsPrivate() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool {
		if msg := update.Message; msg != nil {
			// if message sender's id is equal to the chat-id,
			// then it's a private message.
			return msg.From.Id == msg.Chat.Id
		}

		return false
	})
}

// Commands tests if the message's (and only message) text or caption
// matches the cmd.
func Command(cmd, botUsername string) tgo.Filter {
	return Commands(botUsername, cmd)
}

// Commands tests if the message's (and only message) text or caption
// matches any of the cmds.
func Commands(botUsername string, cmds ...string) tgo.Filter {
	// make sure they are all lower-cased
	for index, command := range cmds {
		cmds[index] = strings.ToLower("/" + command)
	}

	// add a '@' prefix if not set already
	if !strings.HasPrefix(botUsername, "@") {
		botUsername = "@" + botUsername
	}

	return NewFilter(func(update *tg.Update) bool {
		if msg := update.Message; msg != nil {
			text := msg.Text
			if text == "" {
				text = msg.Caption
			}

			// As the commands are already lowercased,
			// the text itself should get lowercased too.
			text = strings.ToLower(text)

			for _, cmd := range cmds {
				// valid cases are:
				// /command
				// /command@username
				// /command args...
				// /command@username args...
				if text == cmd || text == cmd+botUsername || strings.HasPrefix(text, cmd+" ") || strings.HasPrefix(text, cmd+botUsername+" ") {
					return true
				}
			}
		}

		return false
	})
}
