package filters

import "github.com/haashemi/tgo"

type FilterFunc func(update *tgo.Update) bool

type Filter struct{ f FilterFunc }

func (f Filter) Check(update *tgo.Update) bool { return f.f(update) }

func NewFilter(f FilterFunc) *Filter { return &Filter{f: f} }

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

func Text(text string) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool {
		if msg := update.Message; msg != nil {
			return msg.Text == text || msg.Caption == text
		}

		return false
	})
}

func CallbackQuery(query string) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool {
		if q := update.CallbackQuery; q != nil {
			return q.Data == query
		}

		return false
	})
}

func InlineQuery(query string) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool {
		if q := update.InlineQuery; q != nil {
			return q.Query == query
		}

		return false
	})
}
