package filters

import "github.com/haashemi/tgo"

func extractUpdate(update *tgo.Update) any {
	switch {
	case update.Message != nil:
		return update.Message
	case update.EditedMessage != nil:
		return update.EditedMessage
	case update.ChannelPost != nil:
		return update.ChannelPost
	case update.EditedChannelPost != nil:
		return update.EditedChannelPost
	case update.InlineQuery != nil:
		return update.InlineQuery
	case update.ChosenInlineResult != nil:
		return update.ChosenInlineResult
	case update.CallbackQuery != nil:
		return update.CallbackQuery
	case update.ShippingQuery != nil:
		return update.ShippingQuery
	case update.PreCheckoutQuery != nil:
		return update.PreCheckoutQuery
	case update.Poll != nil:
		return update.Poll
	case update.PollAnswer != nil:
		return update.PollAnswer
	case update.MyChatMember != nil:
		return update.MyChatMember
	case update.ChatMember != nil:
		return update.ChatMember
	case update.ChatJoinRequest != nil:
		return update.ChatJoinRequest
	}

	return nil
}

func extractUpdateText(update *tgo.Update) string {
	switch data := extractUpdate(update).(type) {
	case *tgo.Message:
		if data.Caption != "" {
			return data.Caption
		}
		return data.Text
	case *tgo.CallbackQuery:
		return data.Data
	case *tgo.InlineQuery:
		return data.Query
	}

	return ""
}
