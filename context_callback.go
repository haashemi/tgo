package tgo

// SenderID returns the chat id of who sent the query
func (q *CallbackQuery) SenderID() int64 {
	return q.From.Id
}

// ChatID returns the chat id of where the message of that callback button was in.
// it returns the sender's chat-id if the message couldn't be found.
func (q *CallbackQuery) ChatID() int64 {
	if q.Message == nil {
		return q.SenderID()
	}

	return q.Message.Chat.Id
}

// ThreadID returns the chat's thread id of where the message of that callback button was in.
// it returns zero if the message couldn't be found.
func (q *CallbackQuery) ThreadID() int64 {
	if q.Message == nil {
		return 0
	}

	return q.Message.MessageThreadId
}

// MessageID returns message ID of pressed callback button.
// it returns zero if the message couldn't be found.
func (q *CallbackQuery) MessageID() int64 {
	if q.Message == nil {
		return 0
	}

	return q.Message.MessageId
}

// Answer answers to callback queries sent from inline keyboards
func (ctx *botContext) Answer(options *AnswerCallbackQueryOptions) error {
	_, err := ctx.bot.AnswerCallbackQuery(ctx.CallbackQuery().Id, options)
	return err
}
