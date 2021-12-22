package benchplus

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/allegro/bigcache"
	"github.com/benchplus/gocache"
	"github.com/golang/protobuf/proto"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
}

func shutdown() {
}

func BenchmarkPutInt_bigcache(b *testing.B) {
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:             1024,
		LifeWindow:         10 * time.Second,
		MaxEntriesInWindow: 1024 * 10,
		MaxEntrySize:       1024,
		Verbose:            false,
	})
	for i := 0; i < b.N; i++ {
		v := fmt.Sprint(i)
		cache.Set(v, []byte(v))
	}
}

func BenchmarkGetInt_bigcache(b *testing.B) {
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:             1024,
		LifeWindow:         10 * time.Second,
		MaxEntriesInWindow: 1024 * 10,
		MaxEntrySize:       1024,
		Verbose:            false,
	})
	cache.Set("0", []byte("0"))
	for i := 0; i < b.N; i++ {
		cache.Get("0")
	}
}

func BenchmarkPut1K_bigcache(b *testing.B) {
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:             1024,
		LifeWindow:         10 * time.Second,
		MaxEntriesInWindow: 1024 * 10,
		MaxEntrySize:       1024,
		Verbose:            false,
	})
	for i := 0; i < b.N; i++ {
		cache.Set(fmt.Sprint(i), gocache.Data1K)
	}
}

func BenchmarkPut1M_bigcache(b *testing.B) {
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:             1024,
		LifeWindow:         10 * time.Second,
		MaxEntriesInWindow: 1024 * 10,
		MaxEntrySize:       1024,
		Verbose:            false,
	})
	for i := 0; i < b.N; i++ {
		cache.Set(fmt.Sprint(i), gocache.Data1M)
	}
}

func BenchmarkPutTinyObject_bigcache(b *testing.B) {
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:             1024,
		LifeWindow:         10 * time.Second,
		MaxEntriesInWindow: 1024 * 10,
		MaxEntrySize:       1024,
		Verbose:            false,
	})
	for i := 0; i < b.N; i++ {
		data, _ := proto.Marshal(&gocache.UserInfo{})
		cache.Set(fmt.Sprint(i), data)
	}
}

func BenchmarkChangeOutAllInt_bigcache(b *testing.B) {
	cache, _ := bigcache.NewBigCache(bigcache.Config{
		Shards:             1024,
		LifeWindow:         10 * time.Second,
		MaxEntriesInWindow: 1024 * 10,
		MaxEntrySize:       1024,
		Verbose:            false,
	})
	for i := 0; i < b.N*1024; i++ {
		v := fmt.Sprint(i)
		cache.Set(v, []byte(v))
	}
}
