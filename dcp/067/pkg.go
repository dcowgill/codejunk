/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Implement an LFU (Least Frequently Used) cache. It should be able to be
initialized with a cache size n, and contain the following methods:

set(key, value): sets key to value. If there are already n items in the cache
and we are adding a new item, then it should also remove the least frequently
used item. If there is a tie, then the least recently used key should be
removed.

get(key): gets the value at key. If no such key exists, return null.

Each operation should run in O(1) time.

*/
package dcp067

import "fmt"

// An O(1) algorithm for implementing the LFU cache eviction scheme
// Prof. Ketan Shah Anirban Mitra Dhruv Matani
// August 16, 2010
// http://dhruvbird.com/lfu.pdf

type (
	Key   string
	Value string

	FreqNode struct {
		next   *FreqNode  // next value (greater frequency)
		prev   *FreqNode  // previous value (lesser frequency)
		count  int        // number of accesses of values under this node
		values *ValueNode // the actual values
	}

	ValueNode struct {
		freqNode *FreqNode  // points to head of frequency list
		next     *ValueNode // next value w/ same frequency
		prev     *ValueNode // previous value
		key      Key
		value    Value
	}

	Cache struct {
		values   map[Key]*ValueNode // lookup table for keys
		freqs    *FreqNode          // list of frequencies, low to high
		capacity int                // limit on entries in Cache.values
	}
)

func newFreqNode(count int) *FreqNode {
	n := &FreqNode{count: count, values: newValueNode()}
	n.prev = n
	n.next = n
	return n
}

func (fnode *FreqNode) insertAfter(n *FreqNode) {
	n.prev = fnode
	n.next = fnode.next
	fnode.next.prev = n
	fnode.next = n
}

func (fnode *FreqNode) unlink() {
	fnode.prev.next = fnode.next
	fnode.next.prev = fnode.prev
	fnode.prev = fnode
	fnode.next = fnode
}

func (fnode *FreqNode) isEmpty() bool {
	return fnode.values.next == fnode.values
}

func (fnode *FreqNode) removeOne() (*FreqNode, *ValueNode) {
	if fnode.isEmpty() {
		return fnode.next.removeOne()
	}
	dead := fnode.values.next
	dead.unlink()
	return fnode, dead
}

func newValueNode() *ValueNode {
	n := &ValueNode{}
	n.prev = n
	n.next = n
	return n
}

func (vnode *ValueNode) insertAfter(n *ValueNode) {
	n.prev = vnode
	n.next = vnode.next
	vnode.next.prev = n
	vnode.next = n
}

func (vnode *ValueNode) unlink() {
	vnode.prev.next = vnode.next
	vnode.next.prev = vnode.prev
	vnode.prev = vnode
	vnode.next = vnode
}

// Creates a new, empty cache.
func newCache(capacity int) *Cache {
	return &Cache{
		values:   make(map[Key]*ValueNode, capacity),
		freqs:    newFreqNode(0),
		capacity: capacity,
	}
}

// Adds the association (key, value) to the cache with an access count of zero.
// If key already exists, updates its value without adding to its access count.
func (c *Cache) Set(key Key, value Value) {
	if vnode := c.values[key]; vnode != nil {
		vnode.value = value
		return
	}
	if len(c.values) >= c.capacity {
		fnode, vnode := c.freqs.removeOne()
		delete(c.values, vnode.key)
		if fnode != c.freqs {
			fnode.unlink()
		}
	}
	n := &ValueNode{freqNode: c.freqs, key: key, value: value}
	c.values[key] = n
	c.freqs.values.insertAfter(n)
}

// Retrieves the value associated with the key and increments its access count.
func (c *Cache) Get(key Key) (Value, bool) {
	vnode := c.values[key]
	if vnode == nil {
		return "", false
	}
	vnode.unlink()
	fnode := vnode.freqNode
	if fnode.next == c.freqs || fnode.next.count != fnode.count+1 {
		new := newFreqNode(fnode.count + 1)
		new.values.insertAfter(vnode)
		fnode.insertAfter(new)
		fnode = new
	} else {
		fnode = fnode.next
		fnode.values.insertAfter(vnode)
	}
	if vnode.freqNode.isEmpty() && vnode.freqNode != c.freqs {
		vnode.freqNode.unlink()
	}
	vnode.freqNode = fnode
	return vnode.value, true
}

// For debugging.
func (c *Cache) dump() {
	n := c.freqs
	for {
		fmt.Printf("freq %d:\n", n.count)
		for v := n.values.next; v != n.values; v = v.next {
			fmt.Printf("\t(%q, %q)\n", v.key, v.value)
		}
		n = n.next
		if n == c.freqs {
			break
		}
	}
}
