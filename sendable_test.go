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
