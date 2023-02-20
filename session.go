package tgo

import "sync"

type botSession struct {
	mut      sync.Mutex
	sessions map[int64]*Session
}

func newBotSession() *botSession {
	return &botSession{sessions: make(map[int64]*Session)}
}

// GetSession returns the stored session.
// if creates a new session if session id didn't exists.
func (bs *botSession) GetSession(sessionID int64) *Session {
	bs.mut.Lock()
	defer bs.mut.Unlock()

	if s, ok := bs.sessions[sessionID]; ok {
		return s
	}

	s := &Session{data: make(map[string]any)}
	bs.sessions[sessionID] = s
	return s
}

type Session struct {
	mut  sync.RWMutex
	data map[string]any
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
