package tgo

import "sync"

type Session struct {
	mut  sync.RWMutex
	data map[string]any
}

// GetSession returns the stored session.
// if creates a new session if session id didn't exists.
func (bot *Bot) GetSession(sessionID int64) *Session {
	bot.sessionsMut.Lock()
	defer bot.sessionsMut.Unlock()

	if s, ok := bot.sessions[sessionID]; ok {
		return s
	}

	s := &Session{data: make(map[string]any)}
	bot.sessions[sessionID] = s
	return s
}

func (s *Session) Set(key string, value any) {
	s.mut.Lock()
	s.data[key] = value
	s.mut.Unlock()
}

func (s *Session) Get(key string) (data any, ok bool) {
	s.mut.RLock()
	defer s.mut.RUnlock()

	data, ok = s.data[key]
	return
}

func (s *Session) Del(key string) {
	s.mut.Lock()
	delete(s.data, key)
	s.mut.Unlock()
}
