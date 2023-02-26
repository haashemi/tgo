package tgo

type Handler func(ctx Context)

type Filter interface{ Check(update *Update) bool }

type Party interface {
	// Party creates and returns a new Party under the current party.
	//
	// It's useful when you want to categories your handlers with the same filter.
	Party(filter Filter, middlewares ...Handler) Party

	Use(handlers ...Handler)                   // Use calls the passed handlers as a middleware before calling the main handler
	Handle(filter Filter, handlers ...Handler) // Handle calls the handlers when filters gets passed if not any other handler got called before this.
	AfterHandler(handlers ...Handler)          // AfterHandler calls the passed handlers after calling the main handlers.

	handleUpdate(ctx Context) (ok bool) // handleOnMessage gets the update and calls the first passed-filter message handlers. returns true if finds any handler for the update.
}

type UpdateHandler struct {
	Filter   Filter
	Handlers []Handler
}

type party struct {
	filter  Filter
	parties []Party

	middlewares  []Handler
	handlers     []UpdateHandler
	afterHandler []Handler
}

// Party creates and returns a new Party under the current party.
//
// It's useful when you want to categories your handlers with the same filter.
func (p *party) Party(filter Filter, middlewares ...Handler) Party {
	newParty := &party{filter: filter, middlewares: middlewares}

	p.parties = append(p.parties, newParty)

	return newParty
}

// Use calls the passed handlers as a middleware before calling the main handler
func (p *party) Use(handlers ...Handler) {
	p.middlewares = append(p.middlewares, handlers...)
}

// Handle calls the handlers when filters gets passed if not any other handler got called before this.
func (p *party) Handle(filter Filter, handlers ...Handler) {
	p.handlers = append(p.handlers, UpdateHandler{Filter: filter, Handlers: handlers})
}

// AfterHandler calls the passed handlers after calling the main handlers.
func (p *party) AfterHandler(handlers ...Handler) {
	p.afterHandler = append(p.afterHandler, handlers...)
}

// handleOnMessage gets the update and calls the first passed-filter message handlers.
// returns true if finds any handler for the update.
func (p *party) handleUpdate(ctx Context) (done bool) {
	if p.filter != nil && !p.filter.Check(ctx.Update()) {
		return false
	}

	for _, party := range p.parties {
		if ok := party.handleUpdate(ctx); ok {
			return true
		}
	}

	for _, updateHandler := range p.handlers {
		if !updateHandler.Filter.Check(ctx.Update()) {
			continue
		}
		ctx.ResetStopped()

		for _, beforeHandler := range p.middlewares {
			if ctx.IsStopped() {
				continue
			}
			beforeHandler(ctx)
		}

		for _, handler := range updateHandler.Handlers {
			if ctx.IsStopped() {
				continue
			}
			handler(ctx)
		}

		for _, afterHandler := range p.afterHandler {
			if ctx.IsStopped() {
				continue
			}
			afterHandler(ctx)
		}

		if !ctx.IsStopped() {
			return true
		}
	}

	return false
}
