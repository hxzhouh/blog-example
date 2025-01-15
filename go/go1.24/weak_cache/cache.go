package main

import (
	"container/list"
	"fmt"
	"runtime"
	"sync"
	"time"
	"weak"
)

type CacheItem struct {
	key   string
	value any
}
type WeakCache struct {
	cache   map[string]weak.Pointer[list.Element] // Use weak references to store values
	mu      sync.RWMutex
	storage Storage
}

// Storage is a fixed-length cache based on doubly linked tables and weak
type Storage struct {
	capacity int // Maximum size of the cache
	list     *list.List
}

// NewWeakCache creates a fixed-length weak reference cache.
func NewWeakCache(capacity int) *WeakCache {
	return &WeakCache{
		cache:   make(map[string]weak.Pointer[list.Element]),
		storage: Storage{capacity: capacity, list: list.New()},
	}
}

// Set adds or updates cache entries
func (c *WeakCache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// If the element already exists, update the value and move it to the head of the chain table
	if elem, exists := c.cache[key]; exists {
		if elemValue := elem.Value(); elemValue != nil {
			elemValue.Value = &CacheItem{key: key, value: value}
			c.storage.list.MoveToFront(elemValue)
			elemWeak := c.newElemWeak(elemValue)
			c.cache[key] = elemWeak
			return
		} else {
			c.removeElement(key)
		}
	}
	// remove the oldest unused element if capacity is full
	if c.storage.list.Len() >= c.storage.capacity {
		c.evict()
	}

	// Add new element
	elem := c.storage.list.PushFront(&CacheItem{key: key, value: value})
	elemWeak := c.newElemWeak(elem)
	c.cache[key] = elemWeak
}

func (c *WeakCache) newElemWeak(elem *list.Element) weak.Pointer[list.Element] {
	elemWeak := weak.Make(elem)
	runtime.AddCleanup(elem, func(name string) {
		delete(c.cache, name)
	}, elem.Value.(*CacheItem).key)
	return elemWeak
}

// Get gets the value of the cached item
func (c *WeakCache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if elem, exists := c.cache[key]; exists {
		// Check if the weak reference is still valid
		if elemValue := elem.Value(); elemValue != nil {
			// Moving to the head of the chain indicates the most recent visit
			c.storage.list.MoveToFront(elemValue)
			return elemValue.Value.(*CacheItem).value, true
		} else {
			c.removeElement(key)
		}
	}

	return nil, false
}

// evict removes the cache item that has not been used for the longest time
func (c *WeakCache) evict() {
	if elem := c.storage.list.Back(); elem != nil {
		item := elem.Value.(*CacheItem)
		c.removeElement(item.key)
	}
}

// removeElement removes elements from chains and dictionaries.
func (c *WeakCache) removeElement(key string) {
	if elem, exists := c.cache[key]; exists {
		// Check if the weak reference is still valid
		if elemValue := elem.Value(); elemValue != nil {
			c.storage.list.Remove(elemValue)
		}
	}
}

// Debug prints the contents of the cache
func (c *WeakCache) Debug() {
	fmt.Println("Cache content:Size: ", len(c.cache))
	for k, v := range c.cache {
		if v.Value() != nil {
			fmt.Printf("Key: %s, Value: %v\n", k, v.Value().Value.(*CacheItem).value)
		}
	}
}

func (c *WeakCache) CleanCache() {
	c.storage.list.Init()
}

func main() {
	cache := NewWeakCache(3)

	cache.Set("a", "value1")
	cache.Set("b", "value2")
	cache.Set("c", "value3")
	runtime.GC()
	time.Sleep(1 * time.Millisecond)
	cache.Debug()

	// Access "a" to make it the most recently used
	_, _ = cache.Get("a")
	runtime.GC()
	time.Sleep(1 * time.Millisecond)
	cache.Debug()

	// Add new element "d", triggering the elimination of the oldest unused "b".
	cache.Set("d", "value4")
	runtime.GC()
	time.Sleep(1 * time.Millisecond)
	cache.Debug()

	cache.CleanCache()
	runtime.GC()
	time.Sleep(1 * time.Millisecond)
	cache.Debug()
}
