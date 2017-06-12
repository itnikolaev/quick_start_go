package quickstart

type Session struct {
	Settings
	ID    int64
	Email string
}

type SessionStorage interface {
	Get(key string) (*Session, error)
	Set(key string, s *Session)
}
