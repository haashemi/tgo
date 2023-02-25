package tgo

type Party interface {
	// Party creates and returns a new Party under the current party.
	//
	// It's useful when you want to categories your handlers with the same filter.
	Party(filter Filter) Party

	OnMessage(filter Filter, handlers ...MessageHandler) // OnMessage calls the handlers when filters gets passed if not any other handler got called before this.
	BeforeMessage(handlers ...MessageHandler)            // BeforeMessage calls the handlers before calling the main handlers of the passed filter.
	AfterMessage(handlers ...MessageHandler)             // AfterMessage calls the handlers after calling the main handlers of the passed filter.

	OnEditedMessage(filter Filter, handlers ...MessageHandler) // OnEditedMessage calls the handlers when filters gets passed if not any other handler got called before this.
	BeforeEditedMessage(handlers ...MessageHandler)            // BeforeEditedMessage calls the handlers before calling the main handlers of the passed filter.
	AfterEditedMessage(handlers ...MessageHandler)             // AfterEditedMessage calls the handlers after calling the main handlers of the passed filter.

	OnChannelPost(filter Filter, handlers ...MessageHandler) // OnChannelPost calls the handlers when filters gets passed if not any other handler got called before this.
	BeforeChannelPost(handlers ...MessageHandler)            // BeforeChannelPost calls the handlers before calling the main handlers of the passed filter.
	AfterChannelPost(handlers ...MessageHandler)             // AfterChannelPost calls the handlers after calling the main handlers of the passed filter.

	OnEditedChannelPost(filter Filter, handlers ...MessageHandler) // OnEditedChannelPost calls the handlers when filters gets passed if not any other handler got called before this.
	BeforeEditedChannelPost(handlers ...MessageHandler)            // BeforeEditedChannelPost calls the handlers before calling the main handlers of the passed filter.
	AfterEditedChannelPost(handlers ...MessageHandler)             // AfterEditedChannelPost calls the handlers after calling the main handlers of the passed filter.

	OnCallbackQuery(filter Filter, handlers ...CallbackHandler) // OnCallbackQuery calls the handlers when filters gets passed if not any other handler got called before this.
	BeforeCallbackQuery(handlers ...CallbackHandler)            // BeforeCallbackQuery calls the handlers before calling the main handlers of the passed filter.
	AfterCallbackQuery(handlers ...CallbackHandler)             // AfterCallbackQuery calls the handlers after calling the main handlers of the passed filter.

	handleOnMessage(ctx MessageContext) (ok bool)             // handleOnMessage gets the update and calls the first passed-filter message handlers. returns true if finds any handler for the update.
	handleOnEditedMessage(ctx MessageContext) (done bool)     // handleOnEditedMessage gets the update and calls the first passed-filter message handlers. returns true if finds any handler for the update.
	handleOnChannelPost(ctx MessageContext) (done bool)       // handleOnChannelPost gets the update and calls the first passed-filter message handlers. returns true if finds any handler for the update.
	handleOnEditedChannelPost(ctx MessageContext) (done bool) // handleOnEditedChannelPost gets the update and calls the first passed-filter message handlers. returns true if finds any handler for the update.
	handleOnCallbackQuery(ctx CallbackContext) (ok bool)      // handleOnCallbackQuery gets the update and calls the first passed-filter callback query handlers. returns true if finds any handler for the update.
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

	onEditedMessage     []OnMessageHandler
	beforeEditedMessage []MessageHandler
	afterEditedMessage  []MessageHandler

	onChannelPost     []OnMessageHandler
	beforeChannelPost []MessageHandler
	afterChannelPost  []MessageHandler

	onEditedChannelPost     []OnMessageHandler
	beforeEditedChannelPost []MessageHandler
	afterEditedChannelPost  []MessageHandler

	onCallbackQuery     []OnCallbackHandler
	beforeCallbackQuery []CallbackHandler
	afterCallbackQuery  []CallbackHandler

	// ToDo: support updates of type:
	// 	InlineQuery
	// 	ChosenInlineResult
	// 	ShippingQuery
	// 	PreCheckoutQuery
	// 	Poll
	// 	PollAnswer
	// 	MyChatMember
	// 	ChatMember
	// 	ChatJoinRequest
}

// Party creates and returns a new Party under the current party.
//
// It's useful when you want to categories your handlers with the same filter.
func (p *party) Party(filter Filter) Party {
	newParty := &party{filter: filter}

	p.parties = append(p.parties, newParty)

	return newParty
}

// OnMessage calls the handlers when filters gets passed if not any other handler got called before this.
func (p *party) OnMessage(filter Filter, handlers ...MessageHandler) {
	p.onMessage = append(p.onMessage, OnMessageHandler{Filter: filter, Handlers: handlers})
}

// BeforeMessage calls the handlers before calling the main handlers of the passed filter.
func (p *party) BeforeMessage(handlers ...MessageHandler) {
	p.beforeMessage = append(p.beforeMessage, handlers...)
}

// AfterMessage calls the handlers after calling the main handlers of the passed filter.
func (p *party) AfterMessage(handlers ...MessageHandler) {
	p.afterMessage = append(p.afterMessage, handlers...)
}

// OnEditedMessage calls the handlers when filters gets passed if not any other handler got called before this.
func (p *party) OnEditedMessage(filter Filter, handlers ...MessageHandler) {
	p.onEditedMessage = append(p.onEditedMessage, OnMessageHandler{Filter: filter, Handlers: handlers})
}

// BeforeEditedMessage calls the handlers before calling the main handlers of the passed filter.
func (p *party) BeforeEditedMessage(handlers ...MessageHandler) {
	p.beforeEditedMessage = append(p.beforeEditedMessage, handlers...)
}

// AfterEditedMessage calls the handlers after calling the main handlers of the passed filter.
func (p *party) AfterEditedMessage(handlers ...MessageHandler) {
	p.afterEditedMessage = append(p.afterEditedMessage, handlers...)
}

// OnChannelPost calls the handlers when filters gets passed if not any other handler got called before this.
func (p *party) OnChannelPost(filter Filter, handlers ...MessageHandler) {
	p.onChannelPost = append(p.onChannelPost, OnMessageHandler{Filter: filter, Handlers: handlers})
}

// BeforeChannelPost calls the handlers before calling the main handlers of the passed filter.
func (p *party) BeforeChannelPost(handlers ...MessageHandler) {
	p.beforeChannelPost = append(p.beforeChannelPost, handlers...)
}

// AfterChannelPost calls the handlers after calling the main handlers of the passed filter.
func (p *party) AfterChannelPost(handlers ...MessageHandler) {
	p.afterChannelPost = append(p.afterChannelPost, handlers...)
}

// OnEditedChannelPost calls the handlers when filters gets passed if not any other handler got called before this.
func (p *party) OnEditedChannelPost(filter Filter, handlers ...MessageHandler) {
	p.onEditedChannelPost = append(p.onEditedChannelPost, OnMessageHandler{Filter: filter, Handlers: handlers})
}

// BeforeEditedChannelPost calls the handlers before calling the main handlers of the passed filter.
func (p *party) BeforeEditedChannelPost(handlers ...MessageHandler) {
	p.beforeEditedChannelPost = append(p.beforeEditedChannelPost, handlers...)
}

// AfterEditedChannelPost calls the handlers after calling the main handlers of the passed filter.
func (p *party) AfterEditedChannelPost(handlers ...MessageHandler) {
	p.afterEditedChannelPost = append(p.afterEditedChannelPost, handlers...)
}

// OnCallbackQuery calls the handlers when filters gets passed if not any other handler got called before this.
func (p *party) OnCallbackQuery(filter Filter, handlers ...CallbackHandler) {
	p.onCallbackQuery = append(p.onCallbackQuery, OnCallbackHandler{Filter: filter, Handlers: handlers})
}

// BeforeCallbackQuery calls the handlers before calling the main handlers of the passed filter.
func (p *party) BeforeCallbackQuery(handlers ...CallbackHandler) {
	p.beforeCallbackQuery = append(p.beforeCallbackQuery, handlers...)
}

// AfterCallbackQuery calls the handlers after calling the main handlers of the passed filter.
func (p *party) AfterCallbackQuery(handlers ...CallbackHandler) {
	p.afterCallbackQuery = append(p.afterCallbackQuery, handlers...)
}

// handleOnMessage gets the update and calls the first passed-filter message handlers.
// returns true if finds any handler for the update.
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

// handleOnEditedMessage gets the update and calls the first passed-filter message handlers.
// returns true if finds any handler for the update.
func (p *party) handleOnEditedMessage(ctx MessageContext) (done bool) {
	if p.filter != nil && !p.filter.Check(ctx.RawUpdate()) {
		return false
	}

	for _, party := range p.parties {
		if ok := party.handleOnEditedMessage(ctx); ok {
			return true
		}
	}

	for _, updateHandler := range p.onEditedMessage {
		if !updateHandler.Filter.Check(ctx.RawUpdate()) {
			continue
		}
		ctx.ResetStopped()

		for _, beforeHandler := range p.beforeEditedMessage {
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

		for _, afterHandler := range p.afterEditedMessage {
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

// handleOnChannelPost gets the update and calls the first passed-filter message handlers.
// returns true if finds any handler for the update.
func (p *party) handleOnChannelPost(ctx MessageContext) (done bool) {
	if p.filter != nil && !p.filter.Check(ctx.RawUpdate()) {
		return false
	}

	for _, party := range p.parties {
		if ok := party.handleOnChannelPost(ctx); ok {
			return true
		}
	}

	for _, updateHandler := range p.onChannelPost {
		if !updateHandler.Filter.Check(ctx.RawUpdate()) {
			continue
		}
		ctx.ResetStopped()

		for _, beforeHandler := range p.beforeChannelPost {
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

		for _, afterHandler := range p.afterChannelPost {
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

// handleOnEditedChannelPost gets the update and calls the first passed-filter message handlers.
// returns true if finds any handler for the update.
func (p *party) handleOnEditedChannelPost(ctx MessageContext) (done bool) {
	if p.filter != nil && !p.filter.Check(ctx.RawUpdate()) {
		return false
	}

	for _, party := range p.parties {
		if ok := party.handleOnEditedChannelPost(ctx); ok {
			return true
		}
	}

	for _, updateHandler := range p.onEditedChannelPost {
		if !updateHandler.Filter.Check(ctx.RawUpdate()) {
			continue
		}
		ctx.ResetStopped()

		for _, beforeHandler := range p.beforeEditedChannelPost {
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

		for _, afterHandler := range p.afterEditedChannelPost {
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

// handleOnCallbackQuery gets the update and calls the first passed-filter callback query handlers.
// returns true if finds any handler for the update.
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
