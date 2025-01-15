package main

import (
	"fmt"
	_ "fmt"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"testing/synctest"
	"time"
)

func TestGoCacheEntryExpires(t *testing.T) {
	c := cache.New(5*time.Second, 5*time.Second)
	c.Set("foo", "bar", cache.DefaultExpiration)
	v, found := c.Get("foo")
	assert.True(t, found)
	assert.Equal(t, "bar", v)
	time.Sleep(5 * time.Second)
	v, found = c.Get("foo")
	assert.False(t, found)
	assert.Nil(t, v)
}

func TestGoCacheEntryExpiresWithSynctest(t *testing.T) {
	c := cache.New(2*time.Second, 5*time.Second)
	synctest.Run(func() {
		//c := cache.New(2*time.Second, 5*time.Second)
		c.Set("foo", "bar", cache.DefaultExpiration)
		// Get an entry from the cache.
		if got, exist := c.Get("foo"); !exist && got != "bar" {
			t.Errorf("c.Get(k) = %v, want %v", got, "bar")
		}

		// Verify that we get the same entry when accessing it before the expiry.
		time.Sleep(1 * time.Second)
		if got, exist := c.Get("foo"); !exist && got != "bar" {
			t.Errorf("c.Get(k) = %v, want %v", got, 1)
		}
		// Wait for the entry to expire and verify that we now get a new one.
		time.Sleep(3 * time.Second)
		if got, exist := c.Get("foo"); exist {
			t.Errorf("c.Get(k) = %v, want %v", got, nil)
		}
	})
}

func TestSynctest(t *testing.T) {
	synctest.Run(func() {
		before := time.Now()
		fmt.Println("before", before)
		f1 := func() {
			for i := 0; i < 10e9; i++ { // time consuming, It's about 3s in my machine
			}
		}
		go f1()
		synctest.Wait()     //wait f1
		after := time.Now() // time is not affected by the running time of f1
		fmt.Println("after", after)
	})
}

func TestSyncCond(t *testing.T) {
	synctest.Run(func() {
		var mu sync.Mutex
		cond := sync.NewCond(&mu)
		go func() {
			cond.L.Lock()
			cond.Wait()
			cond.L.Unlock()
		}()
		time.Sleep(1 * time.Second)
		cond.Signal()
	})
}
