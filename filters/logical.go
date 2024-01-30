package filters

import "github.com/haashemi/tgo"

// True does nothing and just always returns true.
func True() tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool { return true })
}

// False does nothing and just always returns false.
func False() tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool { return false })
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

// Not Behaves like the ! operator; returns the opposite of the filter result
func Not(filter tgo.Filter) tgo.Filter {
	return NewFilter(func(update *tgo.Update) bool { return !filter.Check(update) })
}
