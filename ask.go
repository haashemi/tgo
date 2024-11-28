package tgo

import (
	"context"
	"fmt"
	"time"

	"github.com/haashemi/tgo/tg"
)

// GetChatAndSenderID extracts and returns the chat ID and sender ID from a given message.
func GetChatAndSenderID(msg *tg.Message) (chatID, senderID int64) {
	chatID = msg.Chat.Id

	if msg.From != nil {
		senderID = msg.From.Id
	} else if msg.SenderChat != nil {
		senderID = msg.SenderChat.Id
	} else {
		senderID = msg.Chat.Id
	}

	return chatID, senderID
}

// getAskUID returns a unique identifier for the ask operation based on the chat ID and sender ID.
func getAskUID(chatID, senderID int64) string {
	return fmt.Sprintf("ask:%d:%d", chatID, senderID)
}

// waitForAnswer waits for an answer from the given UID within the specified timeout duration.
// It returns the received answer message or an error if the timeout is exceeded.
func (bot *Bot) waitForAnswer(uid string, timeout time.Duration) (*tg.Message, error) {
	waiter := make(chan *tg.Message, 1)

	bot.askMut.Lock()
	bot.asks[uid] = waiter
	bot.askMut.Unlock()

	defer func() {
		bot.askMut.Lock()
		delete(bot.asks, uid)
		bot.askMut.Unlock()

		close(waiter)
	}()

	aCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case answer := <-waiter:
		return answer, nil

	case <-aCtx.Done():
		return nil, aCtx.Err()
	}
}

// sendAnswerIfAsked sends the message into the asks channel if it was a response to an ask.
// It returns true if the message was the response to an ask or false otherwise.
func (bot *Bot) sendAnswerIfAsked(msg *tg.Message) (sent bool) {
	bot.askMut.RLock()
	receiver, ok := bot.asks[getAskUID(GetChatAndSenderID(msg))]
	bot.askMut.RUnlock()

	if ok {
		receiver <- msg
		return true
	}

	return false
}

// Ask sends a question message to the specified chat and waits for an answer within the given timeout duration.
// It returns the question message, the received answer message, and any error that occurred.
func (bot *Bot) Ask(chatId, userId int64, msg Sendable, timeout time.Duration) (question, answer *tg.Message, err error) {
	if msg.GetChatID() == nil {
		msg.SetChatID(chatId)
	}
	question, err = bot.Send(msg)
	if err != nil {
		return nil, nil, err
	}

	answer, err = bot.waitForAnswer(getAskUID(chatId, userId), timeout)
	return question, answer, err
}
