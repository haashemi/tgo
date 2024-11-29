package tgo

import (
	"testing"

	"github.com/haashemi/tgo/tg"
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
