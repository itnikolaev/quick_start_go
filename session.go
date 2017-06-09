package quickstart

type Session struct {
	Settings
	ID    int64
	Email string
}

type SessionStorage interface {
	Get(key string) *Session
	Set(key string, s *Session)
}
