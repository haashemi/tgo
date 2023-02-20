package filter

import "github.com/haashemi/tgo"

func Or(filters ...tgo.Filter) tgo.Filter {
	return func(update *tgo.Update) bool {
		for _, filter := range filters {
			if filter(update) {
				return true
			}
		}

		return false
	}
}

func And(filters ...tgo.Filter) tgo.Filter {
	return func(update *tgo.Update) bool {
		for _, filter := range filters {
			if !filter(update) {
				return false
			}
		}

		return true
	}
}

func Text(text string) tgo.Filter {
	return func(update *tgo.Update) bool {
		if msg := update.Message; msg != nil {
			return msg.Text == text || msg.Caption == text
		}

		return false
	}
}

func CallbackQuery(query string) tgo.Filter {
	return func(update *tgo.Update) bool {
		if q := update.CallbackQuery; q != nil {
			return q.Data == query
		}

		return false
	}
}

func InlineQuery(query string) tgo.Filter {
	return func(update *tgo.Update) bool {
		if q := update.InlineQuery; q != nil {
			return q.Query == query
		}

		return false
	}
}
