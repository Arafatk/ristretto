package ristretto

import (
	"bytes"
	"testing"
)

func TestCache(t *testing.T) {
	cacheSize := 100 * 1024 * 1024
	cache := New(cacheSize)
	key := []byte("abc")
	val := []byte("def")
	cache.Set(key, val)
	got, err := cache.Get(key)
	if err != nil || !bytes.Equal(got, val) {
		t.Error("value not equal")
	}
}
