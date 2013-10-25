package tmpdb

import (
	"io"
	"sync"
)

// In memory implementation of File.
// Suitable for tests. Thread-safe.
type MemFile struct {
	m    sync.RWMutex
	data []byte
}

// io.ReadAt
func (f *MemFile) ReadAt(p []byte, off int64) (n int, err error) {
	f.m.RLock()
	data := f.data
	f.m.RUnlock()
	n = copy(p, data[off:])
	if n < len(p) {
		err = io.ErrUnexpectedEOF
	}
	return
}

// io.WriteAt
func (f *MemFile) WriteAt(p []byte, off int64) (n int, err error) {
	f.m.Lock()
	diff := int64(len(f.data) - len(p))
	for diff < off {
		f.data, diff = append(f.data, 0), diff+1
	}
	data := f.data
	f.m.Unlock()
	n = copy(data[off:], p)
	return
}
