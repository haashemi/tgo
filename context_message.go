package tgo

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

// String returns the message's text or media caption
func (m *Message) String() string {
	if m.Text != "" {
		return m.Text
	}

	return m.Caption
}
