package tmpdb

// It never deallocates. Designed for tests.
type DummyStorage struct {
	MemFile
	Offset int64
}

// Storage.Allocate().
func (p *DummyStorage) Allocate(length int64, references []Reference) (offset int64) {
	p.Lock()
	offset, p.Offset = p.Offset, p.Offset+length
	p.Unlock()
	return
}

// Storage.Count().
func (p *DummyStorage) Count(offset, length, count int64) {
}
