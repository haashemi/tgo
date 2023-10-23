package message

import "github.com/haashemi/tgo"

type Handler func(ctx *Context)

type Middleware func(ctx *Context) (ok bool)

type Route struct {
	filter      tgo.Filter
	middlewares []Middleware
	handler     Handler
}

type Router struct {
	middlewares []Middleware
	routes      []Route
}

// NewRouter returns a new message router
func NewRouter(middlewares ...Middleware) *Router {
	return &Router{
		middlewares: middlewares,
	}
}

// Handle adds a new route to the Router
func (r *Router) Handle(filter tgo.Filter, handler Handler, middlewares ...Middleware) {
	r.routes = append(r.routes, Route{filter: filter, middlewares: middlewares, handler: handler})
}

// Setup implements tgo.Router interface
func (r *Router) Setup(bot *tgo.Bot) error { return nil }

// HandleUpdate implements tgo.Router interface
func (r *Router) HandleUpdate(bot *tgo.Bot, upd *tgo.Update) (used bool) {
	if upd.Message == nil {
		return false
	}

	for _, route := range r.routes {
		if !route.filter.Check(upd) {
			continue
		}

		ctx := &Context{Message: upd.Message, Bot: bot}

		allMiddlewares := append(r.middlewares, route.middlewares...)
		for _, middleware := range allMiddlewares {
			if !middleware(ctx) {
				// we used the update, but as the middleware is failed
				// we'll stop the execution and return true as "update is used"
				return true
			}
		}

		route.handler(ctx)

		// filters passed and we used this method, so it's used!
		return true
	}

	return false
}
