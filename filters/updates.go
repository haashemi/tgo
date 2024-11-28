package filters

import (
	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/tg"
)

func HasMessage() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool {
		return update.Message != nil ||
			update.EditedMessage != nil ||
			update.ChannelPost != nil ||
			update.EditedChannelPost != nil
	})
}

func IsMessage() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool { return update.Message != nil })
}

func IsEditedMessage() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool { return update.EditedMessage != nil })
}

func IsChannelPost() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool { return update.ChannelPost != nil })
}

func IsEditedChannelPost() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool { return update.EditedChannelPost != nil })
}

func IsInlineQuery() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool { return update.InlineQuery != nil })
}

func IsChosenInlineResult() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool { return update.ChosenInlineResult != nil })
}

func IsCallbackQuery() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool { return update.CallbackQuery != nil })
}

func IsShippingQuery() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool { return update.ShippingQuery != nil })
}

func IsPreCheckoutQuery() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool { return update.PreCheckoutQuery != nil })
}

func IsPoll() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool { return update.Poll != nil })
}

func IsPollAnswer() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool { return update.PollAnswer != nil })
}

func IsMyChatMember() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool { return update.MyChatMember != nil })
}

func IsChatMember() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool { return update.ChatMember != nil })
}

func IsChatJoinRequest() tgo.Filter {
	return NewFilter(func(update *tg.Update) bool { return update.ChatJoinRequest != nil })
}
