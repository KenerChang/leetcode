package lrucache

import (
	"testing"
)

func TestCache(t *testing.T) {
	cache := Constructor(2 /* capacity */)

	cache.Put(1, 1)
	cache.Put(2, 2)
	result := cache.Get(1) // returns 1
	if result != 1 {
		t.Errorf("expect %d, got %d", 1, result)
	}

	cache.Put(3, 3)       // evicts key 2
	result = cache.Get(2) // returns -1 (not found)
	if result != -1 {
		t.Errorf("expect %d, got %d", -1, result)
	}

	cache.Put(4, 4)       // evicts key 1
	result = cache.Get(1) // returns -1 (not found)
	if result != -1 {
		t.Errorf("expect %d, got %d", 1, result)
	}

	result = cache.Get(3) // returns 3
	if result != 3 {
		t.Errorf("expect %d, got %d", 3, result)
	}

	result = cache.Get(4) // returns 4
	if result != 4 {
		t.Errorf("expect %d, got %d", 4, result)
	}
}

func TestCacheII(t *testing.T) {
	cache := Constructor(2 /* capacity */)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(1, 1)
	cache.Put(3, 3)
	result := cache.Get(2) // returns 1
	if result != -1 {
		t.Errorf("expect %d, got %d", -1, result)
	}

	cache.Put(4, 4)       // evicts key 2
	result = cache.Get(1) // returns -1 (not found)
	if result != -1 {
		t.Errorf("expect %d, got %d", -1, result)
	}

	result = cache.Get(3) // returns -1 (not found)
	if result != 3 {
		t.Errorf("expect %d, got %d", 3, result)
	}

	result = cache.Get(4) // returns 3
	if result != 4 {
		t.Errorf("expect %d, got %d", 4, result)
	}
}
