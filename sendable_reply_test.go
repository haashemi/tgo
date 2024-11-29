package tgo

// TODO: Write tests for these.
var (
	_ Replyable = &SendAnimation{}
	_ Replyable = &SendAudio{}
	_ Replyable = &SendContact{}
	_ Replyable = &SendDice{}
	_ Replyable = &SendDocument{}
	_ Replyable = &SendGame{}
	_ Replyable = &SendInvoice{}
	_ Replyable = &SendLocation{}
	_ Replyable = &SendMessage{}
	_ Replyable = &SendPhoto{}
	_ Replyable = &SendPoll{}
	_ Replyable = &SendSticker{}
	_ Replyable = &SendVenue{}
	_ Replyable = &SendVideo{}
	_ Replyable = &SendVideoNote{}
	_ Replyable = &SendVoice{}
)
