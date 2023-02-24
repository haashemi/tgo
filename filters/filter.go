package filters

import (
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
