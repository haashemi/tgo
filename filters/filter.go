// Filters are set of functions that can be used with routers to determine which
// handler should be called and which should not.
package filters // import "github.com/haashemi/tgo/filters"

import (
	"regexp"
	"strings"

	"github.com/haashemi/tgo"
)

// FilterFunc tests the update with its own filters.
type FilterFunc func(update *tgo.Update) bool

// Filter does nothing and just holds a FilterFunc.
type Filter struct{ f FilterFunc }

// Check calls the FilterFunc and returns the result.
func (f Filter) Check(update *tgo.Update) bool { return f.f(update) }

// NewFilter returns a Filter from FilterFunc f.
func NewFilter(f FilterFunc) *Filter { return &Filter{f: f} }

// Text checks the update's text is equal to the text.
//
// Currently works with Message's caption and text, CallbackQuery's data, and InlineQuery's query.
func Text(text string) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool {
		return extractUpdateText(update) == text
	})
}

// Text checks the update's text is in the texts.
//
// Currently works with Message's caption and text, CallbackQuery's data, and InlineQuery's query.
func Texts(texts ...string) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool {
		raw := extractUpdateText(update)

		for _, text := range texts {
			if raw == text {
				return true
			}
		}

		return false
	})
}

// WithPrefix tests whether the update's text begins with prefix.
//
// Currently works with Message's caption and text, CallbackQuery's data, and InlineQuery's query.
func WithPrefix(prefix string) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool {
		return strings.HasPrefix(extractUpdateText(update), prefix)
	})
}

// WithSuffix tests whether the update's text ends with suffix.
//
// Currently works with Message's caption and text, CallbackQuery's data, and InlineQuery's query.
func WithSuffix(suffix string) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool {
		return strings.HasSuffix(extractUpdateText(update), suffix)
	})
}

// Regex matches the update's text with the reg.
//
// Currently works with Message's caption and text, CallbackQuery's data, and InlineQuery's query.
func Regex(reg *regexp.Regexp) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool {
		return reg.MatchString(extractUpdateText(update))
	})
}

// Whitelist checks if the update is from the whitelisted IDs.
//
// Currently works with Message and CallbackQuery, and InlineQuery.
func Whitelist(IDs ...int64) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool {
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
	})
}
