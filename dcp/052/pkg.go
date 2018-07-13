/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Implement an LRU (Least Recently Used) cache. It should be able to be
initialized with a cache size n, and contain the following methods:

set(key, value): sets key to value. If there are already n items in the cache
and we are adding a new item, then it should also remove the least recently used
item.

get(key): gets the value at key. If no such key exists, return null.

Each operation should run in O(1) time.

*/
package dcp052

// Cache with LRU replacement policy.
type lruCache struct {
	entries  map[string]*node
	head     *node
	capacity int
}

// Node in a doubly linked-list of cache elements.
type node struct {
	key, value string
	prev, next *node
}

// Inserts n immediately after hd.
func (hd *node) insertAfter(n *node) {
	n.prev = hd
	n.next = hd.next
	hd.next.prev = n
	hd.next = n
}

// Removes n from the list.
func (n *node) unlink() {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func newLRUCache(cap int) *lruCache {
	if cap < 1 {
		panic("capacity must be >= 1")
	}
	c := &lruCache{
		entries:  make(map[string]*node, cap),
		head:     &node{}, // sentinel node
		capacity: cap,
	}
	c.head.next, c.head.prev = c.head, c.head
	return c
}

func (c *lruCache) set(key, value string) {
	if len(c.entries) == c.capacity {
		dead := c.head.prev
		dead.unlink()
		delete(c.entries, dead.key)
	}
	n := &node{key: key, value: value}
	c.head.insertAfter(n)
	c.entries[key] = n
}

func (c *lruCache) get(key string) (string, bool) {
	if n := c.entries[key]; n != nil {
		n.unlink()
		c.head.insertAfter(n)
		return n.value, true
	}
	return "", false
}
