package tgo

import "github.com/haashemi/tgo/botapi"

type Party struct {
	filter      Filter
	middlewares []Handler
	handlers    []handlerData
	parties     []*Party
}

type Filter func(update *botapi.Update) bool

type Handler func(ctx Context)

type handlerData struct {
	filter   Filter
	handlers []Handler
}

func (p *Party) handleUpdate(ctx Context) {
	// 1- Filter validation
	if p.filter != nil && !p.filter(ctx.RawUpdate()) {
		return
	}

	// 2- Calling middlewares
	for _, middleware := range p.middlewares {
		middleware(ctx)
		if ctx.isStopped() {
			return
		}
	}

	// 3- Calling handlers with a clone of ctx (as other handlers may stop their context)
	for _, handler := range p.handlers {
		if !handler.filter(ctx.RawUpdate()) {
			continue
		}
		newCtx := ctx.clone()

		go func(hs []Handler, ctx Context) {
			for _, h := range hs {
				h(newCtx)
				if newCtx.isStopped() {
					return
				}
			}
		}(handler.handlers, newCtx)
	}

	// 4- After calling the in-party handlers, we will pass the context to the nested parties
	for _, party := range p.parties {
		party.handleUpdate(ctx)
	}
}

func (p *Party) Handle(filter Filter, handlers ...Handler) {
	p.handlers = append(p.handlers, handlerData{
		filter:   filter,
		handlers: handlers,
	})
}
