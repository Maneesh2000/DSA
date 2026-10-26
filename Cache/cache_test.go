package cache

import "testing"

func TestLRUEvictionAndRecency(t *testing.T) {
	cache := NewLRU(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if value, ok := cache.Get(1); !ok || value != 1 {
		t.Fatalf("Get(1) = (%d, %t), want (1, true)", value, ok)
	}

	// Key 1 was just used, so key 2 is the least recently used entry.
	cache.Put(3, 3)
	if _, ok := cache.Get(2); ok {
		t.Error("Get(2) found an evicted key")
	}
	if value, ok := cache.Get(3); !ok || value != 3 {
		t.Errorf("Get(3) = (%d, %t), want (3, true)", value, ok)
	}
}

func TestLRUUpdateAndZeroCapacity(t *testing.T) {
	cache := NewLRU(1)
	cache.Put(1, 1)
	cache.Put(1, 10)
	if value, ok := cache.Get(1); !ok || value != 10 {
		t.Errorf("updated Get(1) = (%d, %t), want (10, true)", value, ok)
	}

	noStorage := NewLRU(0)
	noStorage.Put(1, 1)
	if noStorage.Len() != 0 {
		t.Errorf("zero-capacity cache length = %d, want 0", noStorage.Len())
	}
}

func TestLFUEviction(t *testing.T) {
	cache := NewLFU(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if value, ok := cache.Get(1); !ok || value != 1 {
		t.Fatalf("Get(1) = (%d, %t), want (1, true)", value, ok)
	}

	// Key 2 has frequency one, while key 1 has frequency two.
	cache.Put(3, 3)
	if _, ok := cache.Get(2); ok {
		t.Error("Get(2) found an LFU-evicted key")
	}
	if value, ok := cache.Get(3); !ok || value != 3 {
		t.Errorf("Get(3) = (%d, %t), want (3, true)", value, ok)
	}
}

func TestLFUTieBreakAndUpdate(t *testing.T) {
	cache := NewLFU(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	// Both keys have frequency one; key 1 is least recently used at that level.
	cache.Put(3, 3)
	if _, ok := cache.Get(1); ok {
		t.Error("Get(1) found a key that should lose the LFU recency tie-break")
	}

	cache.Put(2, 20) // Updating a key changes its value and raises its frequency.
	if value, ok := cache.Get(2); !ok || value != 20 {
		t.Errorf("updated Get(2) = (%d, %t), want (20, true)", value, ok)
	}
}
