package filters

import (
	"regexp"

	"github.com/haashemi/tgo"
)

type FilterFunc func(update *tgo.Update) bool

type Filter struct{ f FilterFunc }

func (f Filter) Check(update *tgo.Update) bool { return f.f(update) }

func NewFilter(f FilterFunc) *Filter { return &Filter{f: f} }

// True does nothing and just always returns true.
func True() tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool { return true })
}

// False does nothing and just always returns false.
func False() tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool { return false })
}

// Or behaves like the || operator; returns true if at least one of the passed filters passes.
// returns false if none of them passes.
func Or(filters ...tgo.Filter) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool {
		for _, filter := range filters {
			if filter.Check(update) {
				return true
			}
		}

		return false
	})
}

// And Behaves like the && operator; returns true if all of the passes filters passes, otherwise returns false.
func And(filters ...tgo.Filter) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool {
		for _, filter := range filters {
			if !filter.Check(update) {
				return false
			}
		}

		return true
	})
}

// Not Behaves like the ! operator; returns the opposite of the filter result
func Not(filter tgo.Filter) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool { return !filter.Check(update) })
}

// Text compares the update (message's text or caption, callback query, inline query) with the passed text.
func Text(text string) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool {
		return extractUpdateText(update) == text
	})
}

// Texts compares the update (message's text or caption, callback query, inline query) with the passed texts.
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

// Regex matches the update (message's text or caption, callback query, inline query) with the passed regexp.
func Regex(reg *regexp.Regexp) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool {
		return reg.MatchString(extractUpdateText(update))
	})
}

// Whitelist compares IDs with the sender-id of the message or callback query. returns true if sender-id is in the blacklist.
func Whitelist(IDs ...int64) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool {
		var senderID int64

		switch data := extractUpdate(update).(type) {
		case *tgo.Message:
			senderID = data.SenderID()
		case *tgo.CallbackQuery:
			senderID = data.SenderID()
		default:
			// avoid unnecessary id comparisons.
			return false
		}

		for _, id := range IDs {
			if id == senderID {
				return true
			}
		}

		return false
	})
}

// Blacklist compares IDs with the sender-id of the message or callback query. returns false if sender-id is in the blacklist.
func Blacklist(IDs ...int64) tgo.Filter {
	// Blacklist works the same as Whitelist, So, why not reducing duplicate code!
	return Not(Whitelist(IDs...))
}
