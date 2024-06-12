package tgo

import (
	"regexp"
	"time"
)

var RateLimitDescriptionRegex = regexp.MustCompile(`^Too Many Requests: retry after [0-9]+$`)

type Error struct {
	ErrorCode   int                 `json:"error_code,omitempty"`
	Description string              `json:"description,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

func (e Error) Error() string { return e.Description }

// A list of most of the common/possible API errors.
//
// source: https://github.com/TelegramBotAPI/errors/tree/ab18477b56d5894e61a015c0775cad20721faa8d
var (
	// Bot token is incorrect.
	ErrUnauthorized = Error{ErrorCode: 401, Description: "Unauthorized"}

	// The chat is unknown to the bot.
	ErrChatNotFound = Error{ErrorCode: 400, Description: "Bad Request: chat not found"}

	// UserId is incorrect.
	ErrUserNotFound = Error{ErrorCode: 400, Description: "[Error]: Bad Request: user not found"}

	// You're trying to perform an action on a user account that has been deactivated or deleted.
	ErrUserIsDeactivated = Error{ErrorCode: 403, Description: "Forbidden: user is deactivated"}

	// Bot was kicked.
	ErrBotWasKicked = Error{ErrorCode: 403, Description: "Forbidden: bot was kicked from the group chat"}

	// The user have blocked the bot.
	ErrBotBlockedByUser = Error{ErrorCode: 403, Description: "Forbidden: bot was blocked by the user"}

	// You tried to send a message to another bot. This is not possible.
	ErrBotCantSendMessageToBots = Error{ErrorCode: 403, Description: "Forbidden: bot can't send messages to bots"}

	// Occurs when a group chat has been converted/migrated to a supergroup.
	//
	// NOTE: DO NOT use this error as it doesn't contain parameters. use IsGroupMigratedToSupergroupErr instead.
	ErrGroupMigratedToSupergroup = Error{ErrorCode: 400, Description: "Bad Request: group chat was migrated to a supergroup chat"}

	// The file id you are trying to retrieve doesn't exist.
	ErrInvalidFileID = Error{ErrorCode: 400, Description: "Bad Request: invalid file id"}

	// The current and new message text and reply markups are the same.
	ErrMessageNotModified = Error{ErrorCode: 400, Description: "Bad Request: message is not modified"}

	// You have already set up a webhook and are trying to get the updates via getUpdates.
	ErrTerminatedByOtherLongPoll = Error{ErrorCode: 409, Description: "Conflict: terminated by other long poll or webhook"}

	// Occurs when the ``action'' property value is invalid.
	ErrWrongParameterActionInRequest = Error{ErrorCode: 400, Description: "Bad Request: wrong parameter action in request"}

	// The message text is empty or not provided.
	ErrMessageTextIsEmpty = Error{ErrorCode: 400, Description: "Bad Request: message text is empty"}

	// The message text cannot be edited.
	ErrMessageCantBeEdited = Error{ErrorCode: 400, Description: "Bad Request: message can't be edited"}

	// You are trying to use getUpdates while a webhook is active.
	ErrCantUseGetUpdatesWithWebhook = Error{ErrorCode: 409, Description: "Conflict: can't use getUpdates method while webhook is active; use deleteWebhook to delete the webhook first"}
)

// IsRateLimitErr returns rate-limited duration and yes as true if the error is
// about when you are hitting the API limit.
func IsRateLimitErr(err error) (retryAfter time.Duration, yes bool) {
	tgErr, isTgError := err.(Error)
	if !isTgError {
		return 0, false
	}

	codeMatches := tgErr.ErrorCode == 429
	descriptionMatches := RateLimitDescriptionRegex.Match([]byte(tgErr.Description))

	if codeMatches && descriptionMatches && tgErr.Parameters != nil {
		return time.Duration(tgErr.Parameters.RetryAfter) * time.Second, true
	}

	return 0, false
}

// IsGroupMigratedToSupergroupErr returns new ChatID and yes as true if the error
// is about when a group chat has been converted/migrated to a supergroup.
func IsGroupMigratedToSupergroupErr(err error) (newChatID int64, yes bool) {
	tgErr, isTgError := err.(Error)
	if !isTgError {
		return 0, false
	}

	if tgErr.ErrorCode == ErrGroupMigratedToSupergroup.ErrorCode &&
		tgErr.Description == ErrGroupMigratedToSupergroup.Description &&
		tgErr.Parameters != nil {
		return tgErr.Parameters.MigrateToChatId, true
	}

	return 0, false
}
