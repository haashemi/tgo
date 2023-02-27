package tgo

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TestKeyValues = map[string]any{
	"int":      69,
	"int2":     85,
	"float":    69.85,
	"float2":   85.69,
	"string":   "tgo is gonna be awesome...",
	"string2":  "tgo is actually awesome!",
	"complex":  69 + 85i,
	"complex2": 85 + 69i,
}

func TestGetSession(t *testing.T) {
	var sessionIds []int64

	for i := 0; i < 50; i++ {
		sessionIds = append(sessionIds, rand.Int63())
	}

	var bot = &Bot{}

	var wg sync.WaitGroup

	wg.Add(len(sessionIds) * 2)

	for _, sid := range sessionIds {
		go func(sid int64) {
			session := bot.GetSession(sid)

			assert.NotNil(t, session)

			wg.Done()
		}(sid)
	}

	// Make sure they are still exists and didn't gone for some reason
	for _, sid := range sessionIds {
		go func(sid int64) {
			session := bot.GetSession(sid)

			assert.NotNil(t, session)

			wg.Done()
		}(sid)
	}

	wg.Wait()
}

func TestSession(t *testing.T) {
	session := &Session{
		data: map[string]any{},
	}

	var wg sync.WaitGroup

	for key, value := range TestKeyValues {
		wg.Add(1)
		go func(key string, value any) {
			defer wg.Done()

			session.Set(key, value)
			gotValue, ok := session.Get(key)

			assert.True(t, ok)
			assert.Equal(t, value, gotValue)

			session.Del(key)
			gotValue, ok = session.Get(key)

			assert.False(t, ok)
			assert.Empty(t, gotValue)
		}(key, value)
	}

	wg.Wait()
}

func TestNilDataSession(t *testing.T) {
	assert.NotPanics(t, func() {
		nilDataSession := &Session{}
		nilDataSession.Set("key", "value")
	})

	assert.NotPanics(t, func() {
		nilDataSession := &Session{}
		nilDataSession.Get("key")
	})

	assert.NotPanics(t, func() {
		nilDataSession := &Session{}
		nilDataSession.Del("key")
	})
}
