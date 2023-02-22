package tgo

import "sync"

type Filter interface{ Check(update *Update) bool }

type MessageHandler func(ctx MessageContext)

type MessageParty struct {
	filter      Filter
	middlewares []MessageHandler
	// parties     []*MessageParty // ToDo

	onMessageMut sync.RWMutex
	onMessage    map[Filter][]MessageHandler
	// onEditedMessage      map[Filter][]MessageHandler // TODO++
	// onChannelPost        map[Filter][]MessageHandler // TODO++
	// onEditedChannelPost  map[Filter][]MessageHandler // TODO++
	// onInlineQuery        map[Filter][]MessageHandler // TODO+
	// onChosenInlineResult map[Filter][]MessageHandler // TODO+
	// onCallbackQuery      map[Filter][]MessageHandler // TODO+++
	// onShippingQuery      map[Filter][]MessageHandler // TODO
	// onPreCheckoutQuery   map[Filter][]MessageHandler // TODO
	// onPoll               map[Filter][]MessageHandler // TODO
	// onPollAnswer         map[Filter][]MessageHandler // TODO
	// onMyChatMember       map[Filter][]MessageHandler // TODO
	// onChatMember         map[Filter][]MessageHandler // TODO
	// onChatJoinRequest    map[Filter][]MessageHandler // TODO
}

func (p *MessageParty) OnMessage(filter Filter, handlers ...MessageHandler) {
	p.onMessageMut.Lock()
	p.onMessage[filter] = handlers
	p.onMessageMut.Unlock()
}

func (p *MessageParty) checkUpdate(ctx MessageContext, update *Update) bool {
	// 1- Filter validation
	if p.filter != nil && !p.filter.Check(update) {
		return false
	}

	// 2- Calling middlewares
	for _, middleware := range p.middlewares {
		middleware(ctx)
		if ctx.IsStopped() {
			return false
		}
	}

	return true
}

func (p *MessageParty) handleOnMessage(ctx MessageContext, update *Update) {
	if !p.checkUpdate(ctx, update) {
		return
	}

	for filter, handlers := range p.onMessage {

		if filter.Check(ctx.RawUpdate()) {

			go func(handlers []MessageHandler, ctx MessageContext) {

				for _, handler := range handlers {
					handler(ctx)

					if ctx.IsStopped() {
						return
					}
				}

			}(handlers, ctx)

			return
		}
	}

	// ToDo: work on nested parties
	// for _, party := range p.parties {
	// 	party.handleOnMessage(ctx)
	// }
}
