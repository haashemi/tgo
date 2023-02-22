package tgo

type Party interface {
	Party(filter Filter) Party

	OnMessage(filter Filter, handlers ...MessageHandler)
	BeforeMessage(handlers ...MessageHandler)
	AfterMessage(handlers ...MessageHandler)

	OnCallbackQuery(filter Filter, handlers ...CallbackHandler)
	BeforeCallbackQuery(handlers ...CallbackHandler)
	AfterCallbackQuery(handlers ...CallbackHandler)

	handleOnMessage(ctx MessageContext) (ok bool)
	handleOnCallbackQuery(ctx CallbackContext) (ok bool)
}

type Filter interface{ Check(update *Update) bool }

type MessageHandler func(ctx MessageContext)

type CallbackHandler func(ctx CallbackContext)

type OnMessageHandler struct {
	Filter   Filter
	Handlers []MessageHandler
}

type OnCallbackHandler struct {
	Filter   Filter
	Handlers []CallbackHandler
}

type party struct {
	filter  Filter
	parties []Party

	onMessage     []OnMessageHandler
	beforeMessage []MessageHandler
	afterMessage  []MessageHandler

	onCallbackQuery     []OnCallbackHandler
	beforeCallbackQuery []CallbackHandler
	afterCallbackQuery  []CallbackHandler

	// onEditedMessage      map[Filter][]MessageHandler // TODO++
	// onChannelPost        map[Filter][]MessageHandler // TODO++
	// onEditedChannelPost  map[Filter][]MessageHandler // TODO++
	// onInlineQuery        map[Filter][]MessageHandler // TODO+
	// onChosenInlineResult map[Filter][]MessageHandler // TODO+
	// onShippingQuery      // TODO
	// onPreCheckoutQuery   // TODO
	// onPoll               // TODO
	// onPollAnswer         // TODO
	// onMyChatMember       // TODO
	// onChatMember         // TODO
	// onChatJoinRequest    // TODO
}

func (p *party) Party(filter Filter) Party {
	newParty := &party{filter: filter}

	p.parties = append(p.parties, newParty)

	return newParty
}

func (p *party) OnMessage(filter Filter, handlers ...MessageHandler) {
	p.onMessage = append(p.onMessage, OnMessageHandler{Filter: filter, Handlers: handlers})
}

func (p *party) BeforeMessage(handlers ...MessageHandler) {
	p.beforeMessage = append(p.beforeMessage, handlers...)
}

func (p *party) AfterMessage(handlers ...MessageHandler) {
	p.afterMessage = append(p.afterMessage, handlers...)
}

func (p *party) OnCallbackQuery(filter Filter, handlers ...CallbackHandler) {
	p.onCallbackQuery = append(p.onCallbackQuery, OnCallbackHandler{Filter: filter, Handlers: handlers})
}

func (p *party) BeforeCallbackQuery(handlers ...CallbackHandler) {
	p.beforeCallbackQuery = append(p.beforeCallbackQuery, handlers...)
}

func (p *party) AfterCallbackQuery(handlers ...CallbackHandler) {
	p.afterCallbackQuery = append(p.afterCallbackQuery, handlers...)
}

func (p *party) handleOnMessage(ctx MessageContext) (done bool) {
	if p.filter != nil && !p.filter.Check(ctx.RawUpdate()) {
		return false
	}

	for _, party := range p.parties {
		if ok := party.handleOnMessage(ctx); ok {
			return true
		}
	}

	for _, updateHandler := range p.onMessage {
		if !updateHandler.Filter.Check(ctx.RawUpdate()) {
			continue
		}
		ctx.ResetStopped()

		for _, beforeHandler := range p.beforeMessage {
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

		for _, afterHandler := range p.afterMessage {
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

func (p *party) handleOnCallbackQuery(ctx CallbackContext) (done bool) {
	if p.filter != nil && !p.filter.Check(ctx.RawUpdate()) {
		return false
	}

	for _, party := range p.parties {
		if ok := party.handleOnCallbackQuery(ctx); ok {
			return true
		}
	}

	for _, updateHandler := range p.onCallbackQuery {
		if !updateHandler.Filter.Check(ctx.RawUpdate()) {
			continue
		}
		ctx.ResetStopped()

		for _, beforeHandler := range p.beforeCallbackQuery {
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

		for _, afterHandler := range p.afterCallbackQuery {
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
