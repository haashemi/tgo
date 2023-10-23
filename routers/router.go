package routers

import "github.com/haashemi/tgo"

// Handler is a basic handler used in Route
type Handler func(bot *tgo.Bot, upd *tgo.Update)

// Middleware is a basic middleware used in Router and Route
type Middleware func(bot *tgo.Bot, upd *tgo.Update) (ok bool)

// Route is a basic route for the Router
type Route struct {
	filter      tgo.Filter
	middlewares []Middleware
	handler     Handler
}

// Router is a basic router for all type of updates.
type Router struct {
	middlewares []Middleware
	routes      []Route
}

// NewRouter returns a new Router; nothing special
func NewRouter() *Router {
	return &Router{}
}

// Use, adds the passed middlewares to the Router
func (r *Router) Use(middlewares ...Middleware) {
	r.middlewares = append(r.middlewares, middlewares...)
}

// Handle, adds a new route to the Router
func (r *Router) Handle(filter tgo.Filter, handler Handler, middlewares ...Middleware) {
	r.routes = append(r.routes, Route{filter: filter, middlewares: middlewares, handler: handler})
}

// Setup implements tgo.Router interface
func (r *Router) Setup(bot *tgo.Bot) error { return nil }

// HandleUpdate implements tgo.Router interface
func (r *Router) HandleUpdate(bot *tgo.Bot, upd *tgo.Update) (used bool) {
	for _, route := range r.routes {
		if !route.filter.Check(upd) {
			continue
		}

		allMiddlewares := append(r.middlewares, route.middlewares...)
		for _, middleware := range allMiddlewares {
			if !middleware(bot, upd) {
				// we used the update, but as the middleware is failed
				// we'll stop the execution and return true as "update is used"
				return true
			}
		}

		route.handler(bot, upd)

		// filters passed and we used this method, so it's used!
		return true
	}

	return false
}
