package main

import (
	"sync"
	"testing"
)

func BenchmarkSafeCache(b *testing.B) {
	cache := &SafeCache[int]{data: make(map[string]int)}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.Set("key", 42)
			cache.Get("key")
			cache.Delete("key")
		}
	})
}

func BenchmarkSyncMap(b *testing.B) {
	var sm sync.Map

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sm.Store("key", 42)
			sm.Load("key")
			sm.Delete("key")
		}
	})
}
