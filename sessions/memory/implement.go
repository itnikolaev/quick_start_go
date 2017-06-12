package memory

import (
	"sync"

	"github.com/itnikolaev/quick_start_go"
	"fmt"
)

type Storage struct {
	*sync.RWMutex
	data map[string]*quickstart.Session
}

var notFoundError error = fmt.Errorf("not found")

func NewStorage() quickstart.SessionStorage {
	return &Storage{
		&sync.RWMutex{},
		make(map[string]*quickstart.Session),
	}
}

func (s *Storage) Get(key string) (*quickstart.Session, error) {
	s.RLock()
	defer s.RUnlock()
	if sess, ok := s.data[key]; ok {
		return sess, nil
	}
	return nil, notFoundError
}

func (s *Storage) Set(key string, sess *quickstart.Session) {
	s.Lock()
	s.data[key] = sess
	s.Unlock()
}
