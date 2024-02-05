package storage

import (
	"strconv"
	"sync"
	"testing"
)

// TestNewURLStorage tests the initialization of URLStorage
func TestNewURLStorage(t *testing.T) {
	idStorage := NewURLStorage()
	if idStorage == nil {
		t.Error("NewURLStorage() should not return nil")
	}
}

func TestURLStorage_Add(t *testing.T) {
	idStorage := NewURLStorage()
	id := "UwaF9CSP"
	origURL := "https://practicum.yandex.ru/"

	added := idStorage.Add(id, origURL)
	if !added {
		t.Errorf("Add() failed to add new ID")
	}
	// Test attempting to add a duplicate ID
	added = idStorage.Add(id, origURL)
	if added {
		t.Error("Add() should not add duplicate IDs")
	}
}

func TestURLStorage_Get(t *testing.T) {
	idStorage := NewURLStorage()
	id := "UwaF9CSP"
	origURL := "https://practicum.yandex.ru/"

	added := idStorage.Add(id, origURL)
	if !added {
		t.Errorf("Add() failed to add new ID")
	}
	retrievedURL, exists := idStorage.Get(id)
	if !exists && retrievedURL != origURL {
		t.Errorf("Get() failed to retrieve URL, expected %v, got %v", origURL, retrievedURL)
	}
}

func TestURLStorage_Concurrency(t *testing.T) {
	idStorage := NewURLStorage()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			idString := strconv.Itoa(id)
			idStorage.Add(idString, "http://example.com")
		}(i)
	}

	wg.Wait()

	if len(idStorage.idMap) != 100 {
		t.Errorf("Expected 100 entries in map, got %d", len(idStorage.idMap))
	}

	for i := 0; i < 0; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			idString := strconv.Itoa(id)
			idStorage.Get(idString)
		}(i)
	}
	wg.Wait()
}
