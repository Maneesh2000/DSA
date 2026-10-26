package cache

// LFUCache stores integer key-value pairs and evicts the least frequently used
// item when full. When several entries have the same frequency, it evicts the
// least recently used among them. It is not safe for concurrent use without
// external locking. Get and Put both run in O(1) average time.
type LFUCache struct {
	capacity       int
	minFrequency   int
	entries        map[int]*lfuNode
	frequencyLists map[int]*lfuList
}

type lfuNode struct {
	key       int
	value     int
	frequency int
	prev      *lfuNode
	next      *lfuNode
}

// lfuList holds nodes with the same frequency. Nodes nearest head were used
// most recently, while the node nearest tail is the eviction candidate.
type lfuList struct {
	head *lfuNode
	tail *lfuNode
	len  int
}

// NewLFU returns an LFU cache with the supplied capacity. A non-positive
// capacity creates a cache that does not retain values.
func NewLFU(capacity int) *LFUCache {
	return &LFUCache{
		capacity:       capacity,
		entries:        make(map[int]*lfuNode),
		frequencyLists: make(map[int]*lfuList),
	}
}

// Get returns the value for key and true when it is cached. A successful read
// increments the entry's frequency and makes it most recent at that frequency.
func (c *LFUCache) Get(key int) (int, bool) {
	node, exists := c.entries[key]
	if !exists {
		return 0, false
	}

	c.incrementFrequency(node)
	return node.value, true
}

// Put inserts or updates key. Updating an existing entry counts as a use and
// increments its frequency. A new entry starts with frequency one.
func (c *LFUCache) Put(key, value int) {
	if c.capacity <= 0 {
		return
	}

	if node, exists := c.entries[key]; exists {
		node.value = value
		c.incrementFrequency(node)
		return
	}

	if len(c.entries) == c.capacity {
		leastUsed := c.frequencyLists[c.minFrequency]
		victim := leastUsed.removeLeastRecent()
		delete(c.entries, victim.key)
		if leastUsed.len == 0 {
			delete(c.frequencyLists, c.minFrequency)
		}
	}

	node := &lfuNode{key: key, value: value, frequency: 1}
	c.entries[key] = node
	c.listForFrequency(1).addMostRecent(node)
	c.minFrequency = 1
}

// Len returns the number of values currently stored in the cache.
func (c *LFUCache) Len() int {
	return len(c.entries)
}

// incrementFrequency moves node to the list for its next frequency.
func (c *LFUCache) incrementFrequency(node *lfuNode) {
	oldFrequency := node.frequency
	oldList := c.frequencyLists[oldFrequency]
	oldList.remove(node)

	if oldList.len == 0 {
		delete(c.frequencyLists, oldFrequency)
		if c.minFrequency == oldFrequency {
			// node is moving from oldFrequency to oldFrequency+1, so no lower
			// frequency can remain after its old list becomes empty.
			c.minFrequency++
		}
	}

	node.frequency++
	c.listForFrequency(node.frequency).addMostRecent(node)
}

// listForFrequency returns the list for frequency, creating it if needed.
func (c *LFUCache) listForFrequency(frequency int) *lfuList {
	list, exists := c.frequencyLists[frequency]
	if !exists {
		list = newLFUList()
		c.frequencyLists[frequency] = list
	}
	return list
}

func newLFUList() *lfuList {
	head := &lfuNode{}
	tail := &lfuNode{}
	head.next = tail
	tail.prev = head
	return &lfuList{head: head, tail: tail}
}

// addMostRecent inserts node directly after the head sentinel.
func (l *lfuList) addMostRecent(node *lfuNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
	l.len++
}

// removeLeastRecent removes and returns the node before the tail sentinel.
func (l *lfuList) removeLeastRecent() *lfuNode {
	node := l.tail.prev
	l.remove(node)
	return node
}

// remove detaches node from its frequency list.
func (l *lfuList) remove(node *lfuNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	l.len--
}
