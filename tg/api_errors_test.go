package tg

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestIsRateLimitErr(t *testing.T) {
	retryAfterTime := rand.Int63()
	rateLimitError := Error{
		ErrorCode:   429,
		Description: fmt.Sprintf("Too Many Requests: retry after %d", retryAfterTime),
		Parameters:  &ResponseParameters{RetryAfter: retryAfterTime},
	}

	tests := []struct {
		Name           string
		Error          error
		WantRetryAfter time.Duration
		WantYes        bool
	}{
		{Name: "Not tgo error", Error: errors.New("a random error"), WantRetryAfter: time.Duration(0), WantYes: false},
		{Name: "Not RateLimit error", Error: ErrBotWasKicked, WantRetryAfter: time.Duration(0), WantYes: false},
		{Name: "RateLimit error", Error: rateLimitError, WantRetryAfter: time.Duration(retryAfterTime) * time.Second, WantYes: true},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			haveRetryAfter, haveYes := IsRateLimitErr(test.Error)

			if haveRetryAfter != test.WantRetryAfter {
				t.Errorf("unmatched retry after. got: %v, expected: %v", haveRetryAfter, test.WantRetryAfter)
			} else if haveYes != test.WantYes {
				t.Errorf("unmatched yes boolean. got: %v, expected: %v", haveYes, test.WantYes)
			}
		})
	}
}

func TestIsGroupMigratedToSupergroupErr(t *testing.T) {
	newChatID := rand.Int63()
	groupMigratedToSupergroupError := Error{
		ErrorCode:   ErrGroupMigratedToSupergroup.ErrorCode,
		Description: ErrGroupMigratedToSupergroup.Description,
		Parameters:  &ResponseParameters{MigrateToChatId: newChatID},
	}

	tests := []struct {
		Name          string
		Error         error
		WantNewChatID int64
		WantYes       bool
	}{
		{Name: "Not tgo error", Error: errors.New("a random error"), WantNewChatID: 0, WantYes: false},
		{Name: "Not GroupMigratedToSupergroup error", Error: ErrInvalidFileID, WantNewChatID: 0, WantYes: false},
		{Name: "GroupMigratedToSupergroup error", Error: groupMigratedToSupergroupError, WantNewChatID: newChatID, WantYes: true},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			haveNewChatID, haveYes := IsGroupMigratedToSupergroupErr(test.Error)

			if haveNewChatID != test.WantNewChatID {
				t.Errorf("unmatched new chat id. got: %v, expected: %v", haveNewChatID, test.WantNewChatID)
			} else if haveYes != test.WantYes {
				t.Errorf("unmatched yes boolean. got: %v, expected: %v", haveYes, test.WantYes)
			}
		})
	}

}
