package benchplus

import (
	"fmt"
	"testing"
	"time"

	"github.com/benchplus/gocache"
	"github.com/orca-zhang/cache"
)

func BenchmarkPutInt_cacheLRU2(b *testing.B) {
	cache := cache.NewLRUCache(1024, 1024, 10*time.Second).LRU2(1024)
	for i := 0; i < b.N; i++ {
		cache.Put(fmt.Sprint(i), i)
	}
}

func BenchmarkPut1K_cacheLRU2(b *testing.B) {
	cache := cache.NewLRUCache(1024, 1024, 10*time.Second).LRU2(1024)
	for i := 0; i < b.N; i++ {
		cache.Put(fmt.Sprint(i), gocache.Data1K)
	}
}

func BenchmarkPut1M_cacheLRU2(b *testing.B) {
	cache := cache.NewLRUCache(1024, 1024, 10*time.Second).LRU2(1024)
	for i := 0; i < b.N; i++ {
		cache.Put(fmt.Sprint(i), gocache.Data1M)
	}
}

func BenchmarkPutTinyObject_cacheLRU2(b *testing.B) {
	cache := cache.NewLRUCache(1024, 1024, 10*time.Second).LRU2(1024)
	for i := 0; i < b.N; i++ {
		cache.Put(fmt.Sprint(i), gocache.UserInfo{})
	}
}

func BenchmarkPutChangeOutAll_cacheLRU2(b *testing.B) {
	cache := cache.NewLRUCache(1024, 1024, 10*time.Second).LRU2(1024)
	for i := 0; i < b.N*1024; i++ {
		cache.Put(fmt.Sprint(i), i)
	}
}
