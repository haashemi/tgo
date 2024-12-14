package tgo

import (
	"math/rand"
	"testing"

	"github.com/haashemi/tgo/tg"
)

var (
	sendableAnimation Sendable = &SendAnimation{}
	sendableAudio     Sendable = &SendAudio{}
	sendableContact   Sendable = &SendContact{}
	sendableDice      Sendable = &SendDice{}
	sendableDocument  Sendable = &SendDocument{}
	sendableGame      Sendable = &SendGame{}
	sendableInvoice   Sendable = &SendInvoice{}
	sendableLocation  Sendable = &SendLocation{}
	sendableMessage   Sendable = &SendMessage{}
	sendablePhoto     Sendable = &SendPhoto{}
	sendablePoll      Sendable = &SendPoll{}
	sendableSticker   Sendable = &SendSticker{}
	sendableVenue     Sendable = &SendVenue{}
	sendableVideo     Sendable = &SendVideo{}
	sendableVideoNote Sendable = &SendVideoNote{}
	sendableVoice     Sendable = &SendVoice{}
)

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

var (
	parseModeSettableAnimation ParseModeSettable = &SendAnimation{}
	parseModeSettableAudio     ParseModeSettable = &SendAudio{}
	parseModeSettableDocument  ParseModeSettable = &SendDocument{}
	parseModeSettableMessage   ParseModeSettable = &SendMessage{}
	parseModeSettablePhoto     ParseModeSettable = &SendPhoto{}
	parseModeSettableVideo     ParseModeSettable = &SendVideo{}
	parseModeSettableVoice     ParseModeSettable = &SendVoice{}
)

// TestSendables tests Sendables' GetChatID and SetChatID methods and ensures
// that they sets and returns the right value.
//
// TODO: Test if they call the right API method on Send method.
func TestSendables(t *testing.T) {
	var tests = map[string]struct {
		T  Sendable
		ID int64
	}{
		"Sendable Animation": {T: sendableAnimation, ID: rand.Int63()},
		"Sendable Audio":     {T: sendableAudio, ID: rand.Int63()},
		"Sendable Contact":   {T: sendableContact, ID: rand.Int63()},
		"Sendable Dice":      {T: sendableDice, ID: rand.Int63()},
		"Sendable Document":  {T: sendableDocument, ID: rand.Int63()},
		"Sendable Game":      {T: sendableGame, ID: rand.Int63()},
		"Sendable Invoice":   {T: sendableInvoice, ID: rand.Int63()},
		"Sendable Location":  {T: sendableLocation, ID: rand.Int63()},
		"Sendable Message":   {T: sendableMessage, ID: rand.Int63()},
		"Sendable Photo":     {T: sendablePhoto, ID: rand.Int63()},
		"Sendable Poll":      {T: sendablePoll, ID: rand.Int63()},
		"Sendable Sticker":   {T: sendableSticker, ID: rand.Int63()},
		"Sendable Venue":     {T: sendableVenue, ID: rand.Int63()},
		"Sendable Video":     {T: sendableVideo, ID: rand.Int63()},
		"Sendable VideoNote": {T: sendableVideoNote, ID: rand.Int63()},
		"Sendable Voice":     {T: sendableVoice, ID: rand.Int63()},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.T.SetChatID(test.ID)

			got := test.T.GetChatID()
			expected := tg.ID(test.ID)
			if got != expected {
				t.Errorf("got: %v, expected: %v", got, expected)
			}
		})
	}
}

// TestParseModeSettables tests ParseModeSettables' GetParseMode and SetParseMode
// methods with all of the available parse modes and ensures that they set and
// return the right parse mode.
func TestParseModeSettables(t *testing.T) {
	var tests = map[string]struct {
		T         ParseModeSettable
		parseMode tg.ParseMode
	}{
		"parseModeSettable Animation None":       {T: parseModeSettableAnimation, parseMode: tg.ParseModeNone},
		"parseModeSettable Animation Markdown":   {T: parseModeSettableAnimation, parseMode: tg.ParseModeMarkdown},
		"parseModeSettable Animation MarkdownV2": {T: parseModeSettableAnimation, parseMode: tg.ParseModeMarkdownV2},
		"parseModeSettable Animation HTML":       {T: parseModeSettableAnimation, parseMode: tg.ParseModeHTML},
		"parseModeSettable Audio None":           {T: parseModeSettableAudio, parseMode: tg.ParseModeNone},
		"parseModeSettable Audio Markdown":       {T: parseModeSettableAudio, parseMode: tg.ParseModeMarkdown},
		"parseModeSettable Audio MarkdownV2":     {T: parseModeSettableAudio, parseMode: tg.ParseModeMarkdownV2},
		"parseModeSettable Audio HTML":           {T: parseModeSettableAudio, parseMode: tg.ParseModeHTML},
		"parseModeSettable Document None":        {T: parseModeSettableDocument, parseMode: tg.ParseModeNone},
		"parseModeSettable Document Markdown":    {T: parseModeSettableDocument, parseMode: tg.ParseModeMarkdown},
		"parseModeSettable Document MarkdownV2":  {T: parseModeSettableDocument, parseMode: tg.ParseModeMarkdownV2},
		"parseModeSettable Document HTML":        {T: parseModeSettableDocument, parseMode: tg.ParseModeHTML},
		"parseModeSettable Message None":         {T: parseModeSettableMessage, parseMode: tg.ParseModeNone},
		"parseModeSettable Message Markdown":     {T: parseModeSettableMessage, parseMode: tg.ParseModeMarkdown},
		"parseModeSettable Message MarkdownV2":   {T: parseModeSettableMessage, parseMode: tg.ParseModeMarkdownV2},
		"parseModeSettable Message HTML":         {T: parseModeSettableMessage, parseMode: tg.ParseModeHTML},
		"parseModeSettable Photo None":           {T: parseModeSettablePhoto, parseMode: tg.ParseModeNone},
		"parseModeSettable Photo Markdown":       {T: parseModeSettablePhoto, parseMode: tg.ParseModeMarkdown},
		"parseModeSettable Photo MarkdownV2":     {T: parseModeSettablePhoto, parseMode: tg.ParseModeMarkdownV2},
		"parseModeSettable Photo HTML":           {T: parseModeSettablePhoto, parseMode: tg.ParseModeHTML},
		"parseModeSettable Video None":           {T: parseModeSettableVideo, parseMode: tg.ParseModeNone},
		"parseModeSettable Video Markdown":       {T: parseModeSettableVideo, parseMode: tg.ParseModeMarkdown},
		"parseModeSettable Video MarkdownV2":     {T: parseModeSettableVideo, parseMode: tg.ParseModeMarkdownV2},
		"parseModeSettable Video HTML":           {T: parseModeSettableVideo, parseMode: tg.ParseModeHTML},
		"parseModeSettable Voice None":           {T: parseModeSettableVoice, parseMode: tg.ParseModeNone},
		"parseModeSettable Voice Markdown":       {T: parseModeSettableVoice, parseMode: tg.ParseModeMarkdown},
		"parseModeSettable Voice MarkdownV2":     {T: parseModeSettableVoice, parseMode: tg.ParseModeMarkdownV2},
		"parseModeSettable Voice HTML":           {T: parseModeSettableVoice, parseMode: tg.ParseModeHTML},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.T.SetParseMode(test.parseMode)

			got := test.T.GetParseMode()
			expected := test.parseMode
			if got != expected {
				t.Errorf("got: %v, expected: %v", got, expected)
			}
		})
	}
}
