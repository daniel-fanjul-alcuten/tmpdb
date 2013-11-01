package tmpdb

import (
	"io"
	"sync"
)

// In-memory file. Designed for tests.
type MemFile struct {
	sync.RWMutex
	Data []byte
}

// io.ReadAt
func (f *MemFile) ReadAt(p []byte, off int64) (n int, err error) {
	f.RLock()
	data := f.Data
	f.RUnlock()
	if n = copy(p, data[off:]); n < len(p) {
		err = io.ErrUnexpectedEOF
	}
	return
}

// io.WriteAt
func (f *MemFile) WriteAt(p []byte, off int64) (n int, err error) {
	f.Lock()
	for diff := int64(len(f.Data) - len(p)); diff < off; {
		f.Data, diff = append(f.Data, 0), diff+1
	}
	data := f.Data
	f.Unlock()
	n = copy(data[off:], p)
	return
}
