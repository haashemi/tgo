package tgo

import (
	"context"
	"fmt"
	"time"
)

func GetChatAndSenderID(msg *Message) (chatID, senderID int64) {
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

func GetAskUID(chatID, senderID int64) string {
	return fmt.Sprintf("ask:%d:%d", chatID, senderID)
}

func (bot *Bot) waitForAnswer(uid string, timeout time.Duration) (*Message, error) {
	waiter := make(chan *Message, 1)

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

func (bot *Bot) sendAnswerIfAsked(msg *Message) (sent bool) {
	bot.askMut.RLock()
	receiver, ok := bot.asks[GetAskUID(GetChatAndSenderID(msg))]
	bot.askMut.RUnlock()

	if ok {
		receiver <- msg
		return true
	}

	return false
}

func (bot *Bot) Ask(chatId, userId int64, msg Sendable, timeout time.Duration) (question, answer *Message, err error) {
	question, err = bot.Send(msg)
	if err != nil {
		return nil, nil, err
	}

	answer, err = bot.waitForAnswer(GetAskUID(chatId, userId), timeout)
	return question, answer, err
}
