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

// Most well-known telegram bot-api errors.
//
// Thanks to github.com/TelegramBotAPI/errors
var (
	ErrUnamortized                   = &Error{ErrorCode: 401, Description: "Unauthorized"}                                                                                                 // The bot token is incorrect.
	ErrChatNotFound                  = &Error{ErrorCode: 400, Description: "Bad Request: chat not found"}                                                                                  // Chat is unknown to the bot.
	ErrUserNotFound                  = &Error{ErrorCode: 400, Description: "[Error]: Bad Request: user not found"}                                                                         // user_id is incorrect.
	ErrUserIsDeactivated             = &Error{ErrorCode: 403, Description: "Forbidden: user is deactivated"}                                                                               // You're trying to perform an action on a user account that has been deactivated or deleted.
	ErrBotWasKicked                  = &Error{ErrorCode: 403, Description: "Forbidden: bot was kicked from the group chat"}                                                                // Bot was kicked.
	ErrBotBlockedByUser              = &Error{ErrorCode: 403, Description: "Forbidden: bot was blocked by the user"}                                                                       // The user have blocked the bot.
	ErrBotCantSendMessageToBots      = &Error{ErrorCode: 403, Description: "Forbidden: bot can't send messages to bots"}                                                                   // You tried to send a message to another bot. This is not possible.
	ErrInvalidFileID                 = &Error{ErrorCode: 400, Description: "Bad Request: invalid file id"}                                                                                 // The file id you are trying to retrieve doesn't exist
	ErrMessageNotModified            = &Error{ErrorCode: 400, Description: "Bad Request: message is not modified"}                                                                         // The current and new message text and reply markups are the same
	ErrTerminatedByOtherLongPoll     = &Error{ErrorCode: 409, Description: "Conflict: terminated by other long poll or webhook"}                                                           // You have already set up a webhook and are trying to get the updates via getUpdates
	ErrWrongParameterActionInRequest = &Error{ErrorCode: 400, Description: "Bad Request: wrong parameter action in request"}                                                               // Occurs when the action property value is invalid
	ErrMessageTextIsEmpty            = &Error{ErrorCode: 400, Description: "Bad Request: message text is empty"}                                                                           // The message text is empty or not provided
	ErrCantUseGetUpdatesWithWebhook  = &Error{ErrorCode: 409, Description: "Conflict: can't use getUpdates method while webhook is active; use deleteWebhook to delete the webhook first"} // You are trying to use getUpdates while a webhook is active
)

// IsRateLimitErr returns rate-limited duration and yes as true if the error is about when you are hitting the API limit.
func IsRateLimitErr(err error) (retryAfter time.Duration, yes bool) {
	tgErr, isTgError := err.(*Error)
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

// IsGroupMigratedToSupergroupErr returns new ChatID and yes as true if the error is about when a group chat has been converted/migrated to a supergroup.
func IsGroupMigratedToSupergroupErr(err error) (newChatID int64, yes bool) {
	tgErr, isTgError := err.(*Error)
	if !isTgError {
		return 0, false
	}

	codeMatches := tgErr.ErrorCode == 400
	descriptionMatches := tgErr.Description == "Bad Request: group chat was migrated to a supergroup chat"

	if codeMatches && descriptionMatches && tgErr.Parameters != nil {
		return tgErr.Parameters.MigrateToChatId, true
	}

	return 0, false
}
