package tmpdb

// An AllocationPool implementation. It never deallocates.
// No thread-safe.
type DummyAllocationPool struct {
	off int64
}

// AllocationPool.Allocate().
func (p *DummyAllocationPool) Allocate(length int64) (off int64) {
	off, p.off = p.off, p.off+length
	return
}

// AllocationPool.Dellocate(). It never deallocates.
func (p *DummyAllocationPool) Deallocate(off int64, length int64) {
}

// A RefCountPool implementation. It never counts.
type DummyRefCountPool struct{}

// RefCountPool.Increment(). It always returns 1.
func (p *DummyRefCountPool) Increment(int64, RefCount) RefCount {
	return 1
}
