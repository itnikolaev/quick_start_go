package memory

import (
	"sync"

	"github.com/itnikolaev/quick_start_go"
)

type Storage struct {
	mux  sync.RWMutex
	data map[string]*quickstart.Session
}

func NewStorage() quickstart.SessionStorage {
	return &Storage{
		data: make(map[string]*quickstart.Session),
	}
}

func (s *Storage) Get(key string) *quickstart.Session {
	s.mux.RLock()
	defer s.mux.RUnlock()
	if sess, ok := s.data[key]; ok {
		return sess
	}
	return nil
}

func (s *Storage) Set(key string, sess *quickstart.Session) {
	s.mux.Lock()
	s.data[key] = sess
	s.mux.Unlock()
}
