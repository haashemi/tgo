package tgo

// Sendable is an interface that represents any object that can be sent using an API client.
type Sendable interface {
	// GetChatID returns the chat ID associated with the sendable object.
	GetChatID() ChatID

	// SetChatID sets the chat ID for the sendable object.
	SetChatID(id int64)

	// Send sends the sendable object using the provided API client.
	// It returns the sent message and any error encountered.
	Send(api *API) (*Message, error)
}

// ParseModeSettable is an interface that represents any object that can have its ParseMode set
// Or in other words, messages with captions.
type ParseModeSettable interface {
	GetParseMode() ParseMode
	SetParseMode(mode ParseMode)
}

// Replyable is an interface that represents any object that can be replied to.
type Replyable interface {
	Sendable
	SetReplyToMessageId(id int64)
}

// Send sends a message with the preferred ParseMode.
func (b *Bot) Send(msg Sendable) (*Message, error) {
	if x, ok := msg.(ParseModeSettable); ok {
		if x.GetParseMode() == ParseModeNone {
			x.SetParseMode(b.DefaultParseMode)
		}
	}

	return msg.Send(b.API)
}
