package graphdb

import (
	"math/rand"
	"sync"
)

// syncMap wraps built-in map by using RWMutex for concurrent safe.
type safeMap struct {
	items map[uint]bool
	sync.RWMutex
}

// SyncMap keeps a slice of *syncMap with length of `shardCount`.
// Using a slice of syncMap instead of a large one is to avoid lock bottlenecks.
type SafeMap struct {
	shardCount uint8
	shards     []*safeMap
}

// Create a new SyncMap with default shard count.
func SafemapNew() *SafeMap {
	return SafemapNewWithShard(defaultShardCount)
}

// Create a new SyncMap with given shard count.
// NOTE: shard count must be power of 2, default shard count will be used otherwise.
func SafemapNewWithShard(shardCount uint8) *SafeMap {
	if !isPowerOfTwo(shardCount) {
		shardCount = defaultShardCount
	}
	m := new(SafeMap)
	m.shardCount = shardCount
	m.shards = make([]*safeMap, m.shardCount)
	for i, _ := range m.shards {
		m.shards[i] = &safeMap{items: make(map[uint]bool)}
	}
	return m
}

// Find the specific shard with the given key
func (m *SafeMap) locate(key uint) *safeMap {
	return m.shards[key&uint((m.shardCount-1))]
}

// Retrieves a value
func (m *SafeMap) Get(key uint) (value bool, ok bool) {
	shard := m.locate(key)
	shard.RLock()
	value, ok = shard.items[key]
	shard.RUnlock()
	return
}

// Sets value with the given key
func (m *SafeMap) Set(key uint, value bool) {
	shard := m.locate(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

// Removes an item
func (m *SafeMap) Delete(key uint) {
	shard := m.locate(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

// Pop delete and return a random item in the cache
func (m *SafeMap) Pop() (uint, bool) {
	if m.Size() == 0 {
		panic("safemap: map is empty")
	}

	var (
		key   uint
		value bool
		found = false
		n     = int(m.shardCount)
	)

	for !found {
		idx := rand.Intn(n)
		shard := m.shards[idx]
		shard.Lock()
		if len(shard.items) > 0 {
			found = true
			for key, value = range shard.items {
				break
			}
			delete(shard.items, key)
		}
		shard.Unlock()
	}

	return key, value
}

// Whether SyncMap has the given key
func (m *SafeMap) Has(key uint) bool {
	_, ok := m.Get(key)
	return ok
}

// Returns the number of items
func (m *SafeMap) Size() int {
	size := 0
	for _, shard := range m.shards {
		shard.RLock()
		size += len(shard.items)
		shard.RUnlock()
	}
	return size
}

// Wipes all items from the map
func (m *SafeMap) Flush() int {
	size := 0
	for _, shard := range m.shards {
		shard.Lock()
		size += len(shard.items)
		shard.items = make(map[uint]bool)
		shard.Unlock()
	}
	return size
}

// Returns a channel from which each key in the map can be read
func (m *SafeMap) IterKeys() <-chan uint {
	ch := make(chan uint)
	go func() {
		for _, shard := range m.shards {
			shard.RLock()
			for key, _ := range shard.items {
				ch <- key
			}
			shard.RUnlock()
		}
		close(ch)
	}()
	return ch
}

// Item is a pair of key and value
type SafeItem struct {
	Key   uint
	Value bool
}

// Return a channel from which each item (key:value pair) in the map can be read
func (m *SafeMap) IterItems() <-chan SafeItem {
	ch := make(chan SafeItem)
	go func() {
		for _, shard := range m.shards {
			shard.RLock()
			for key, value := range shard.items {
				ch <- SafeItem{key, value}
			}
			shard.RUnlock()
		}
		close(ch)
	}()
	return ch
}
