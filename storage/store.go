package storage

import "sync"

type Store struct {
	urls map[string]string
	mu   sync.RWMutex
}

func NewStore() *Store {
	return &Store{urls: map[string]string{}}
}

func (s *Store) Save(shortUrl string, longUrl string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[shortUrl] = longUrl
}

func (s *Store) Get(shortUrl string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, exists := s.urls[shortUrl]

	return val, exists
}
