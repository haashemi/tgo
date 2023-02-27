package tgo

import "sync"

// Session is a simple key-value storage to store some user's data
type Session struct {
	mut  sync.RWMutex
	data map[string]any
}

// GetSession returns the stored session.
// if creates a new session if session id didn't exists.
func (bot *Bot) GetSession(sessionID int64) *Session {
	bot.sessionsMut.Lock()
	defer bot.sessionsMut.Unlock()

	if bot.sessions == nil {
		bot.sessions = map[int64]*Session{}
	}

	if s, ok := bot.sessions[sessionID]; ok {
		return s
	}

	s := &Session{data: make(map[string]any)}
	bot.sessions[sessionID] = s
	return s
}

// Set stores the value in the session
func (s *Session) Set(key string, value any) {
	s.mut.Lock()
	if s.data == nil {
		s.data = map[string]any{}
	}
	s.data[key] = value
	s.mut.Unlock()
}

// Get returns the stored value from the session
func (s *Session) Get(key string) (data any, ok bool) {
	s.mut.RLock()
	defer s.mut.RUnlock()

	data, ok = s.data[key]
	return
}

// Del deletes the value from the session
func (s *Session) Del(key string) {
	s.mut.Lock()
	delete(s.data, key)
	s.mut.Unlock()
}
