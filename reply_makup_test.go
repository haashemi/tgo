package tgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Interface check
var (
	TestInlineKeyboardMarkup ReplyMarkup = &InlineKeyboardMarkup{}
	TestReplyKeyboardMarkup  ReplyMarkup = &ReplyKeyboardMarkup{}
	TestReplyKeyboardRemove  ReplyMarkup = &ReplyKeyboardRemove{}
	TestForceReply           ReplyMarkup = &ForceReply{}
)

func TestNewInlineKeyboardMarkup(t *testing.T) {
	emptyKeyboard := NewInlineKeyboardMarkup()
	assert.Empty(t, emptyKeyboard)

	row1, row2 := NewInlineKeyboardRow(), NewInlineKeyboardRow()

	expected := &InlineKeyboardMarkup{InlineKeyboard: [][]*InlineKeyboardButton{row1, row2}}
	got := NewInlineKeyboardMarkup(row1, row2)

	assert.Equal(t, expected, got)
}

func TestNewInlineKeyboardRow(t *testing.T) {
	emptyRow := NewInlineKeyboardRow()
	assert.Empty(t, emptyRow)

	btn1 := CallbackDataInlineKeyboardButton("text1", "callback1")
	btn2 := CallbackDataInlineKeyboardButton("text2", "callback2")

	expected := []*InlineKeyboardButton{btn1, btn2}
	got := NewInlineKeyboardRow(btn1, btn2)

	assert.Equal(t, expected, got)
}

func TestUrlInlineKeyboardButton(t *testing.T) {
	expected := &InlineKeyboardButton{Text: "some_text", Url: "https://example.com"}
	btn := UrlInlineKeyboardButton("some_text", "https://example.com")

	assert.Equal(t, expected, btn)
}

func TestCallbackDataInlineKeyboardButton(t *testing.T) {
	expected := &InlineKeyboardButton{Text: "some_text", CallbackData: "someCallbackData"}
	btn := CallbackDataInlineKeyboardButton("some_text", "someCallbackData")

	assert.Equal(t, expected, btn)
}

func TestWebAppInlineKeyboardButton(t *testing.T) {
	expected := &InlineKeyboardButton{Text: "some_text", WebApp: &WebAppInfo{Url: "https://example.com/webAppURL"}}
	btn := WebAppInlineKeyboardButton("some_text", "https://example.com/webAppURL")

	assert.Equal(t, expected, btn)
}

func TestLoginUrlInlineKeyboardButton(t *testing.T) {
	expected := &InlineKeyboardButton{Text: "some_text", LoginUrl: &LoginUrl{Url: "https://example.com/login"}}
	btn := LoginUrlInlineKeyboardButton("some_text", &LoginUrl{Url: "https://example.com/login"})

	assert.Equal(t, expected, btn)
}

func TestSwitchInlineQueryInlineKeyboardButton(t *testing.T) {
	expected := &InlineKeyboardButton{Text: "some_text", SwitchInlineQuery: "someQuery"}
	btn := SwitchInlineQueryInlineKeyboardButton("some_text", "someQuery")

	assert.Equal(t, expected, btn)
}

func TestSwitchInlineQueryCurrentChatInlineKeyboardButton(t *testing.T) {
	expected := &InlineKeyboardButton{Text: "some_text", SwitchInlineQueryCurrentChat: "someQuery"}
	btn := SwitchInlineQueryCurrentChatInlineKeyboardButton("some_text", "someQuery")

	assert.Equal(t, expected, btn)
}

func TestCallbackGameInlineKeyboardButton(t *testing.T) {
	expected := &InlineKeyboardButton{Text: "some_text", CallbackGame: &CallbackGame{}}
	btn := CallbackGameInlineKeyboardButton("some_text")

	assert.Equal(t, expected, btn)
}

func TestPayInlineKeyboardButton(t *testing.T) {
	expected := &InlineKeyboardButton{Text: "some_text", Pay: true}
	btn := PayInlineKeyboardButton("some_text")

	assert.Equal(t, expected, btn)
}
