package tgo

type MessageContext = *messageContext

type messageContext struct{ *Context }

// SenderID returns the chat id of who sent the message
func (m *Message) SenderID() int64 {
	if m.From == nil {
		return 0
	}

	return m.From.Id
}

// ChatID returns the chat id of where the message is sent in
func (m *Message) ChatID() int64 {
	return m.Chat.Id
}

// ThreadID returns the chat's thread id of where the message is sent in
func (m *Message) ThreadID() int64 {
	return m.MessageThreadId
}

// MessageID returns ID of the sent message
func (m *Message) MessageID() int64 {
	return m.MessageId
}

func (ctx *messageContext) Message() *Message {
	return ctx.Contextable.(*Message)
}

// Caption returns the message's text or media caption
func (ctx *messageContext) Caption() string {
	msg := ctx.Message()

	if msg.Text != "" {
		return msg.Text
	}
	return msg.Caption
}
