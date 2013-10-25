package tmpdb

import (
	"testing"
)

func TestDummyAllocationPool(t *testing.T) {
	dummyPool := &DummyAllocationPool{0}
	var pool AllocationPool = dummyPool
	if off := pool.Allocate(1); off != 0 {
		t.Error(off)
	}
	if off := pool.Allocate(1); off != 1 {
		t.Error(off)
	}
	if off := pool.Allocate(2); off != 2 {
		t.Error(off)
	}
	if off := pool.Allocate(3); off != 4 {
		t.Error(off)
	}
	if off := pool.Allocate(5); off != 7 {
		t.Error(off)
	}
	if off := pool.Allocate(8); off != 12 {
		t.Error(off)
	}
	if off := pool.Allocate(13); off != 20 {
		t.Error(off)
	}
	if off := pool.Allocate(21); off != 33 {
		t.Error(off)
	}
}

func TestDummyRefCountPool(t *testing.T) {
	dummyPool := &DummyRefCountPool{}
	var pool RefCountPool = dummyPool
	if i := pool.Increment(1, 2); i != 1 {
		t.Error(i)
	}
}
