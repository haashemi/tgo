package tgo

type MessageContext = *messageContext

type messageContext struct {
	Context

	Message *Message
}

// SenderID returns the chat id of who sent the message
func (m *Message) SenderID() int64 {
	if m.From == nil {
		return 0
	}

	return m.From.Id
}

// SenderID returns the chat id of where the message is sent in
func (m *Message) ChatID() int64 {
	if m.Chat == nil {
		return 0
	}

	return m.Chat.Id
}

// SenderID returns the chat's thread id of where the message is sent in
func (m *Message) ThreadID() int64 {
	return m.MessageThreadId
}

// MessageID returns ID of the sent message
func (m *Message) MessageID() int64 {
	return m.MessageId
}

// Caption returns the message's text or media caption
func (ctx *messageContext) Caption() string {
	if ctx.Message.Text != "" {
		return ctx.Message.Text
	}

	return ctx.Message.Caption
}
