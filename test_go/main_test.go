package main

import (
	"testing"
)

func TestCache_Get(t *testing.T) {
	cache := newMainCache()
	cache.Set("key1", "value1")

	value, err := cache.Get("key1")
	if value != "value1" || err != nil {
		t.Errorf("Expected value 'value1', got '%s' with error: %v", value, err)
	}

	_, err = cache.Get("key2")
	if err == nil || err.Error() != "key not found" {
		t.Errorf("Expected error 'key not found', got: %v", err)
	}
}

func TestCache_Set(t *testing.T) {
	cache := newMainCache()

	oldValue := cache.Set("key1", "value1")
	if oldValue != "" {
		t.Errorf("Expected no old value, got: %s", oldValue)
	}

	oldValue = cache.Set("key1", "value2")
	if oldValue != "value1" {
		t.Errorf("Expected old value 'value1', got: %s", oldValue)
	}
}

func TestCache_Del(t *testing.T) {
	cache := newMainCache()
	cache.Set("key1", "value1")

	err := cache.Del("key1")
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	err = cache.Del("key2")
	if err == nil || err.Error() != "key not found" {
		t.Errorf("Expected error 'key not found', got: %v", err)
	}
}
