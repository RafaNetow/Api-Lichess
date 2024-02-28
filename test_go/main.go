package main

import (
	"errors"
	"fmt"
	"sync"
)

type CacheData struct {
	data map[string]string
	mu   sync.Mutex
}

func newMainCache() *CacheData {
	return &CacheData{
		data: make(map[string]string),
	}
}

func (c *CacheData) Get(key string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.data[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return value, nil
}

func (c *CacheData) Set(key string, value string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	oldValue, _ := c.data[key]
	c.data[key] = value
	return oldValue
}

func (c *CacheData) Del(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.data[key]; !ok {
		return errors.New("key not found")
	}
	delete(c.data, key)
	return nil
}

func main() {
	cache := newMainCache()

	oldValue := cache.Set("foo", "bar")
	fmt.Println("Old value:", oldValue)
	fmt.Println("Old value:", cache)

	value, err := cache.Get("foo")
	fmt.Println("Value:", value, "Error:", err)

	_, err = cache.Get("baz")
	fmt.Println("Error:", err)

	err = cache.Del("foo")
	fmt.Println("Error:", err)

	err = cache.Del("foo")
	fmt.Println("Error:", err)
}
