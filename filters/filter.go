// Package filters are set of functions that can be used with routers to determine which
// handler should be called and which should not.
package filters // import "github.com/haashemi/tgo/filters"

import (
	"regexp"
	"strings"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/tg"
)

// Text checks the update's text is equal to the text.
//
// Currently works with Message's caption and text, CallbackQuery's data, and InlineQuery's query.
func Text(text string) tgo.Filter {
	return func(update *tg.Update) bool {
		return extractUpdateText(update) == text
	}
}

// Texts checks the update's text is in the texts.
//
// Currently works with Message's caption and text, CallbackQuery's data, and InlineQuery's query.
func Texts(texts ...string) tgo.Filter {
	return func(update *tg.Update) bool {
		raw := extractUpdateText(update)

		for _, text := range texts {
			if raw == text {
				return true
			}
		}

		return false
	}
}

// WithPrefix tests whether the update's text begins with prefix.
//
// Currently works with Message's caption and text, CallbackQuery's data, and InlineQuery's query.
func WithPrefix(prefix string) tgo.Filter {
	return func(update *tg.Update) bool {
		return strings.HasPrefix(extractUpdateText(update), prefix)
	}
}

// WithSuffix tests whether the update's text ends with suffix.
//
// Currently works with Message's caption and text, CallbackQuery's data, and InlineQuery's query.
func WithSuffix(suffix string) tgo.Filter {
	return func(update *tg.Update) bool {
		return strings.HasSuffix(extractUpdateText(update), suffix)
	}
}

// Regex matches the update's text with the reg.
//
// Currently works with Message's caption and text, CallbackQuery's data, and InlineQuery's query.
func Regex(reg *regexp.Regexp) tgo.Filter {
	return func(update *tg.Update) bool {
		return reg.MatchString(extractUpdateText(update))
	}
}

// Whitelist checks if the update is from the whitelisted IDs.
//
// Currently works with Message and CallbackQuery, and InlineQuery.
func Whitelist(IDs ...int64) tgo.Filter {
	return func(update *tg.Update) bool {
		var senderID int64

		if update.Message != nil {
			senderID = update.Message.From.Id
		} else if update.CallbackQuery != nil {
			senderID = update.CallbackQuery.From.Id
		} else if update.InlineQuery != nil {
			senderID = update.InlineQuery.From.Id
		}

		for _, id := range IDs {
			if id == senderID {
				return true
			}
		}

		return false
	}
}
