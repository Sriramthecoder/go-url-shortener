package store

import "sync"

var (
	data = make(map[string]string)
	mu   sync.RWMutex
)

func Save(code, url string) {
	mu.Lock()
	defer mu.Unlock()
	data[code] = url
}

func Get(code string) string {
	mu.RLock()
	defer mu.RUnlock()
	return data[code]
}
