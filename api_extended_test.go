package tgo

import (
	"math/rand"
	"testing"
)

var (
	_ ChatID = Username("")
	_ ChatID = ID(0)
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

var (
	parseModeSettableAnimation ParseModeSettable = &SendAnimation{}
	parseModeSettableAudio     ParseModeSettable = &SendAudio{}
	parseModeSettableDocument  ParseModeSettable = &SendDocument{}
	parseModeSettableMessage   ParseModeSettable = &SendMessage{}
	parseModeSettablePhoto     ParseModeSettable = &SendPhoto{}
	parseModeSettableVideo     ParseModeSettable = &SendVideo{}
	parseModeSettableVoice     ParseModeSettable = &SendVoice{}
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
			expected := ID(test.ID)
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
		parseMode ParseMode
	}{
		"parseModeSettable Animation None":       {T: parseModeSettableAnimation, parseMode: ParseModeNone},
		"parseModeSettable Animation Markdown":   {T: parseModeSettableAnimation, parseMode: ParseModeMarkdown},
		"parseModeSettable Animation MarkdownV2": {T: parseModeSettableAnimation, parseMode: ParseModeMarkdownV2},
		"parseModeSettable Animation HTML":       {T: parseModeSettableAnimation, parseMode: ParseModeHTML},
		"parseModeSettable Audio None":           {T: parseModeSettableAudio, parseMode: ParseModeNone},
		"parseModeSettable Audio Markdown":       {T: parseModeSettableAudio, parseMode: ParseModeMarkdown},
		"parseModeSettable Audio MarkdownV2":     {T: parseModeSettableAudio, parseMode: ParseModeMarkdownV2},
		"parseModeSettable Audio HTML":           {T: parseModeSettableAudio, parseMode: ParseModeHTML},
		"parseModeSettable Document None":        {T: parseModeSettableDocument, parseMode: ParseModeNone},
		"parseModeSettable Document Markdown":    {T: parseModeSettableDocument, parseMode: ParseModeMarkdown},
		"parseModeSettable Document MarkdownV2":  {T: parseModeSettableDocument, parseMode: ParseModeMarkdownV2},
		"parseModeSettable Document HTML":        {T: parseModeSettableDocument, parseMode: ParseModeHTML},
		"parseModeSettable Message None":         {T: parseModeSettableMessage, parseMode: ParseModeNone},
		"parseModeSettable Message Markdown":     {T: parseModeSettableMessage, parseMode: ParseModeMarkdown},
		"parseModeSettable Message MarkdownV2":   {T: parseModeSettableMessage, parseMode: ParseModeMarkdownV2},
		"parseModeSettable Message HTML":         {T: parseModeSettableMessage, parseMode: ParseModeHTML},
		"parseModeSettable Photo None":           {T: parseModeSettablePhoto, parseMode: ParseModeNone},
		"parseModeSettable Photo Markdown":       {T: parseModeSettablePhoto, parseMode: ParseModeMarkdown},
		"parseModeSettable Photo MarkdownV2":     {T: parseModeSettablePhoto, parseMode: ParseModeMarkdownV2},
		"parseModeSettable Photo HTML":           {T: parseModeSettablePhoto, parseMode: ParseModeHTML},
		"parseModeSettable Video None":           {T: parseModeSettableVideo, parseMode: ParseModeNone},
		"parseModeSettable Video Markdown":       {T: parseModeSettableVideo, parseMode: ParseModeMarkdown},
		"parseModeSettable Video MarkdownV2":     {T: parseModeSettableVideo, parseMode: ParseModeMarkdownV2},
		"parseModeSettable Video HTML":           {T: parseModeSettableVideo, parseMode: ParseModeHTML},
		"parseModeSettable Voice None":           {T: parseModeSettableVoice, parseMode: ParseModeNone},
		"parseModeSettable Voice Markdown":       {T: parseModeSettableVoice, parseMode: ParseModeMarkdown},
		"parseModeSettable Voice MarkdownV2":     {T: parseModeSettableVoice, parseMode: ParseModeMarkdownV2},
		"parseModeSettable Voice HTML":           {T: parseModeSettableVoice, parseMode: ParseModeHTML},
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
