package tgo

type Party interface {
	Party(filter Filter) Party

	OnMessage(filter Filter, handlers ...MessageHandler)
	BeforeOnMessage(handlers ...MessageHandler)
	AfterOnMessage(handlers ...MessageHandler)

	OnCallbackQuery(filter Filter, handlers ...CallbackHandler)
	BeforeOnCallbackQuery(handlers ...CallbackHandler)
	AfterOnCallbackQuery(handlers ...CallbackHandler)

	handleOnMessage(ctx MessageContext, update *Update) (ok bool)
	// handleOnCallbackQuery(ctx *CallbackQuery, update *Update)
}

type Filter interface{ Check(update *Update) bool }

type MessageHandler func(ctx MessageContext)

// ToDo: this is just a placeholder for now, we will create CallbackContext soon
type CallbackHandler func(ctx *CallbackQuery)

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

	onMessage       []OnMessageHandler
	beforeOnMessage []MessageHandler
	afterOnMessage  []MessageHandler

	onCallbackQuery       []OnCallbackHandler
	beforeOnCallbackQuery []CallbackHandler
	afterOnCallbackQuery  []CallbackHandler

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

func (p *party) BeforeOnMessage(handlers ...MessageHandler) {
	p.beforeOnMessage = append(p.beforeOnMessage, handlers...)
}

func (p *party) AfterOnMessage(handlers ...MessageHandler) {
	p.afterOnMessage = append(p.afterOnMessage, handlers...)
}

func (p *party) OnCallbackQuery(filter Filter, handlers ...CallbackHandler) {
	p.onCallbackQuery = append(p.onCallbackQuery, OnCallbackHandler{Filter: filter, Handlers: handlers})
}

func (p *party) BeforeOnCallbackQuery(handlers ...CallbackHandler) {
	p.beforeOnCallbackQuery = append(p.beforeOnCallbackQuery, handlers...)
}

func (p *party) AfterOnCallbackQuery(handlers ...CallbackHandler) {
	p.afterOnCallbackQuery = append(p.afterOnCallbackQuery, handlers...)
}

func (p *party) handleOnMessage(ctx MessageContext, update *Update) (done bool) {
	if p.filter != nil && !p.filter.Check(update) {
		return false
	}

	for _, party := range p.parties {
		if ok := party.handleOnMessage(ctx, update); ok {
			return true
		}
	}

	for _, updateHandler := range p.onMessage {
		if !updateHandler.Filter.Check(ctx.RawUpdate()) {
			continue
		}
		ctx.ResetStopped()

		for _, beforeHandler := range p.beforeOnMessage {
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

		for _, afterHandler := range p.afterOnMessage {
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
