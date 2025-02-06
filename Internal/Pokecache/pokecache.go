package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache  map[string]cacheEntry
	mutex_ *sync.Mutex
}

// You'll probably want to expose a NewCache() function that creates a new cache
//with a configurable interval (time.Duration).

func NewCache(duration time.Duration) Cache {
	cache := Cache{
		make(map[string]cacheEntry),
		&sync.Mutex{},
	}

	go cache.reaploop(duration)

	return cache
}

//	    	Create a cache.Add() method that adds a new entry to the cache.
//			It should take a key (a string) and a val (a []byte).
func (c *Cache) Add(key string, val []byte) {
	c.mutex_.Lock()
	defer c.mutex_.Unlock()

	c.cache[key] = cacheEntry{time.Now().UTC(), val}
}

//	    	Create a cache.Get() method that gets an entry from the cache.
//			It should take a key (a string) and return a []byte and a bool.
//			The bool should be true if the entry was found and false if it wasn't.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex_.Lock()
	defer c.mutex_.Unlock()

	result, ok := c.cache[key]

	return result.val, ok
}

//     	Create a cache.reapLoop() method that is called when the cache is created (by the NewCache function).
//		Each time an interval (the time.Duration passed to NewCache) passes it should remove any
//		entries that are older than the interval.
//		This makes sure that the cache doesn't grow too large over time.
//		For example, if the interval is 5 seconds, and an entry was added 7 seconds ago, that entry should be removed.

func (c *Cache) reaploop(duration time.Duration) {
	ticker := time.NewTicker(duration)

	for range ticker.C {
		func() {
			c.mutex_.Lock()
			defer c.mutex_.Unlock()

			for index, result := range c.cache {
				if result.createdAt.Before(time.Now().UTC().Add(-duration)) {
					delete(c.cache, index)
				}
			}
		}()
	}
}
