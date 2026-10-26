// Package cache contains common cache eviction policies implemented in Go.
package cache

// LRUCache stores integer key-value pairs and evicts the least recently used
// item when full. It is not safe for concurrent use without external locking.
// Get and Put both run in O(1) average time.
type LRUCache struct {
	capacity int
	entries  map[int]*lruNode
	head     *lruNode // Sentinel before the most recently used node.
	tail     *lruNode // Sentinel after the least recently used node.
}

type lruNode struct {
	key   int
	value int
	prev  *lruNode
	next  *lruNode
}

// NewLRU returns an LRU cache with the supplied capacity. A non-positive
// capacity creates a cache that does not retain values.
func NewLRU(capacity int) *LRUCache {
	head := &lruNode{}
	tail := &lruNode{}
	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		entries:  make(map[int]*lruNode),
		head:     head,
		tail:     tail,
	}
}

// Get returns the value for key and true when it is cached. A successful read
// makes the entry the most recently used item.
func (c *LRUCache) Get(key int) (int, bool) {
	node, exists := c.entries[key]
	if !exists {
		return 0, false
	}

	c.moveToFront(node)
	return node.value, true
}

// Put inserts or updates key. Updating an existing key also makes it most
// recently used. When the cache is full, the least recently used key is evicted.
func (c *LRUCache) Put(key, value int) {
	if c.capacity <= 0 {
		return
	}

	if node, exists := c.entries[key]; exists {
		node.value = value
		c.moveToFront(node)
		return
	}

	if len(c.entries) == c.capacity {
		leastRecent := c.tail.prev
		c.remove(leastRecent)
		delete(c.entries, leastRecent.key)
	}

	node := &lruNode{key: key, value: value}
	c.entries[key] = node
	c.addAfterHead(node)
}

// Len returns the number of values currently stored in the cache.
func (c *LRUCache) Len() int {
	return len(c.entries)
}

// moveToFront records a node as the most recently used entry.
func (c *LRUCache) moveToFront(node *lruNode) {
	c.remove(node)
	c.addAfterHead(node)
}

// addAfterHead inserts node directly after the head sentinel.
func (c *LRUCache) addAfterHead(node *lruNode) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

// remove detaches node from the doubly linked list.
func (c *LRUCache) remove(node *lruNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
