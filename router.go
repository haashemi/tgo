package tgo

import "github.com/haashemi/tgo/tg"

type Filter func(update *tg.Update) (ok bool)

type Handler func(ctx Context)

type Middleware func(next Handler) Handler

type Route struct {
	filter  Filter
	handler Handler
}

type Router struct {
	routes []Route
}

// Handle adds a new route to the Router instance.
func (r *Router) Handle(filter Filter, handler Handler, middlewares ...Middleware) {
	route := Route{filter: filter, handler: handler}

	for i := len(middlewares) - 1; i >= 0; i-- {
		route.handler = middlewares[i](handler)
	}

	r.routes = append(r.routes, route)
}

func (r *Router) HandleUpdate(bot *Bot, upd *tg.Update) {
	for _, route := range r.routes {
		if !route.filter(upd) {
			continue
		}

		ctx := NewContext(bot, upd)
		route.handler(ctx)

		// Filters passed and the handler got called, so we can stop matching.
		break
	}
}
