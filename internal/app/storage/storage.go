package storage

import (
	"sync"
)

// storage for hashes
type URLStorage struct {
	mu    sync.RWMutex
	idMap map[string]string
}

// Initializes a new storage
func NewURLStorage() *URLStorage {
	return &URLStorage{
		idMap: make(map[string]string),
	}
}

// Checks if the ID is in storage
func (s *URLStorage) Add(id, origURL string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.idMap[id]; !exists {
		s.idMap[id] = origURL
		return true // ID was added successfully
	}
	return false // ID already exists, not added
}

// Gets the id from the storage
func (s *URLStorage) Get(id string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	origURL, exists := s.idMap[id]
	return origURL, exists
}
