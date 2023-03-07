package tgo

type Filter interface{ Check(update *Update) bool }

type Middleware func(ctx Context) (ok bool)

type Handler func(ctx Context)

type Router interface {
	Setup(bot *Bot) error
	HandleUpdate(ctx Context) (used bool)
}

type DefaultRouter struct {
	middlewares []Middleware
	routes      []DefaultRouterRoute
}

type DefaultRouterRoute struct {
	filter      Filter
	middlewares []Middleware
	handler     Handler
}

func NewDefaultRouter() *DefaultRouter {
	return &DefaultRouter{}
}

func (r *DefaultRouter) Use(middlewares ...Middleware) {
	r.middlewares = append(r.middlewares, middlewares...)
}

func (r *DefaultRouter) Handle(filter Filter, handler Handler, middlewares ...Middleware) {
	r.routes = append(r.routes, DefaultRouterRoute{
		filter:      filter,
		middlewares: middlewares,
		handler:     handler,
	})
}

func (r *DefaultRouter) Setup(bot *Bot) error {
	// we have nothing to setup for the default router
	return nil
}

func (r *DefaultRouter) HandleUpdate(ctx Context) (used bool) {
	for _, route := range r.routes {
		if !route.filter.Check(ctx.Update) {
			continue
		}

		for _, middleware := range r.middlewares {
			if !middleware(ctx) {
				continue
			}
		}

		route.handler(ctx)

		// filters passed and we used this method, so it's used!
		return true
	}

	return false
}
