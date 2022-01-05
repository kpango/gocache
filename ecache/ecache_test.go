package benchplus

import (
	"fmt"
	"os"
	"runtime/debug"
	"sync"
	"testing"
	"time"

	"github.com/benchplus/gocache"
	"github.com/orca-zhang/ecache"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	debug.SetGCPercent(10)
}

func shutdown() {
	gocache.PrintGCPause()
}

func BenchmarkPutInt_ecache(b *testing.B) {
	cache := ecache.NewLRUCache(256, 32, 10*time.Second)
	for i := 0; i < b.N; i++ {
		cache.PutV(fmt.Sprint(i), nil, ecache.D(int64(i+1)))
	}
}

func BenchmarkGetInt_ecache(b *testing.B) {
	cache := ecache.NewLRUCache(256, 32, 10*time.Second)
	cache.PutV("0", nil, ecache.D(1))
	for i := 0; i < b.N; i++ {
		cache.GetV("0")
	}
}

func BenchmarkPut1K_ecache(b *testing.B) {
	cache := ecache.NewLRUCache(256, 32, 10*time.Second)
	for i := 0; i < b.N; i++ {
		cache.PutV(fmt.Sprint(i), nil, gocache.Data1K)
	}
}

func BenchmarkGetIntMiss_ecache(b *testing.B) {
	cache := ecache.NewLRUCache(256, 32, 10*time.Second)
	for i := 0; i < b.N; i++ {
		cache.GetV(fmt.Sprint(i))
	}
}

func BenchmarkPutTinyObject_ecache(b *testing.B) {
	cache := ecache.NewLRUCache(256, 32, 10*time.Second)
	for i := 0; i < b.N; i++ {
		cache.Put(fmt.Sprint(i), gocache.UserInfo{})
	}
}

func BenchmarkChangeOutAllInt_ecache(b *testing.B) {
	cache := ecache.NewLRUCache(256, 32, 10*time.Second)
	for i := 0; i < b.N*1024; i++ {
		cache.PutV(fmt.Sprint(i), nil, ecache.D(int64(i+1)))
	}
}

func BenchmarkHeavyReadInt_ecache(b *testing.B) {
	gocache.GCPause()

	cache := ecache.NewLRUCache(256, 32, 10*time.Second)
	for i := 0; i < 1024; i++ {
		cache.PutV(fmt.Sprint(i), nil, ecache.D(int64(i+1)))
	}
	var wg sync.WaitGroup
	for index := 0; index < 10000; index++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1024; i++ {
				cache.GetV(fmt.Sprint(i))
			}
			wg.Done()
		}()
	}
	wg.Wait()

	gocache.AddGCPause("HeavyReadInt")
}

func BenchmarkHeavyWriteInt_ecache(b *testing.B) {
	gocache.GCPause()

	cache := ecache.NewLRUCache(256, 32, 10*time.Second)
	var wg sync.WaitGroup
	for index := 0; index < 10000; index++ {
		start := index
		wg.Add(1)
		go func() {
			for i := 0; i < 8192; i++ {
				cache.PutV(fmt.Sprint(i+start), nil, ecache.D(int64(i+1)))
			}
			wg.Done()
		}()
	}
	wg.Wait()

	gocache.AddGCPause("HeavyWriteInt")
}

func BenchmarkHeavyWrite1K_ecache(b *testing.B) {
	gocache.GCPause()

	cache := ecache.NewLRUCache(256, 32, 10*time.Second)
	var wg sync.WaitGroup
	for index := 0; index < 10000; index++ {
		start := index
		wg.Add(1)
		go func() {
			for i := 0; i < 8192; i++ {
				cache.PutV(fmt.Sprint(i+start), nil, gocache.Data1K)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	gocache.AddGCPause("HeavyWrite1K")
}
