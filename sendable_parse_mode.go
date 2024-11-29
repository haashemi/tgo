package tgo

import "github.com/haashemi/tgo/tg"

// ParseModeSettable is an interface that represents any object that can have its ParseMode set
// Or in other words, messages with captions.
type ParseModeSettable interface {
	Sendable

	// GetParseMode returns the current set parse mode.
	GetParseMode() tg.ParseMode

	// SetParseMode updates the parse mode.
	SetParseMode(mode tg.ParseMode)
}

func (x *SendAnimation) GetParseMode() tg.ParseMode { return x.ParseMode }
func (x *SendAudio) GetParseMode() tg.ParseMode     { return x.ParseMode }
func (x *SendDocument) GetParseMode() tg.ParseMode  { return x.ParseMode }
func (x *SendMessage) GetParseMode() tg.ParseMode   { return x.ParseMode }
func (x *SendPhoto) GetParseMode() tg.ParseMode     { return x.ParseMode }
func (x *SendVideo) GetParseMode() tg.ParseMode     { return x.ParseMode }
func (x *SendVoice) GetParseMode() tg.ParseMode     { return x.ParseMode }

func (x *SendAnimation) SetParseMode(mode tg.ParseMode) { x.ParseMode = mode }
func (x *SendAudio) SetParseMode(mode tg.ParseMode)     { x.ParseMode = mode }
func (x *SendDocument) SetParseMode(mode tg.ParseMode)  { x.ParseMode = mode }
func (x *SendMessage) SetParseMode(mode tg.ParseMode)   { x.ParseMode = mode }
func (x *SendPhoto) SetParseMode(mode tg.ParseMode)     { x.ParseMode = mode }
func (x *SendVideo) SetParseMode(mode tg.ParseMode)     { x.ParseMode = mode }
func (x *SendVoice) SetParseMode(mode tg.ParseMode)     { x.ParseMode = mode }
