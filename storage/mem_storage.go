package tmpdb

import (
	"container/list"
	. "github.com/daniel-fanjul-alcuten/tmpdb"
	"sort"
	"sync"
)

type refcount struct {
	offset, length int64
	count          int64
}

type group struct {
	length   int64
	elements []*list.Element
}

// Work in progress, NOT FINISHED.
type MemStorage struct {
	file File
	m    sync.Mutex
	// last offset
	eof int64
	// instances of refcount, sorted by offset.
	refcounts list.List
	// from offset to elements in refcounts.
	offsets map[int64]*list.Element
	// groups whose elements have count == 0, sorted by length
	unused []group
}

func NewMemStorage(file File) *MemStorage {
	return &MemStorage{file, sync.Mutex{},
		0, list.List{}, map[int64]*list.Element{}, []group{}}
}

// io.ReadAt().
func (s *MemStorage) ReadAt(p []byte, off int64) (n int, err error) {
	return s.file.ReadAt(p, off)
}

// io.WriteAt().
func (s *MemStorage) WriteAt(p []byte, off int64) (n int, err error) {
	return s.file.WriteAt(p, off)
}

// Storage.Allocate().
func (s *MemStorage) Allocate(length int64, references []Reference) (offset int64) {
	s.m.Lock()
	defer s.m.Unlock()

	// binary search of unused blob
	i := sort.Search(len(s.unused), func(i int) bool {
		return s.unused[i].length >= length
	})
	if i < len(s.unused) {
		// found, take last element of the group
		elements := s.unused[i].elements
		index := len(elements) - 1
		e := elements[index]
		if index == 0 {
			// remove group because it is empty
			panic("TODO remove unused group")
		} else {
			// remove last element of the group
			elements[index] = nil
			s.unused[i].elements = elements[:index]
		}
		rc := e.Value.(*refcount)
		if rc.length == length {
			// the length is exact
			rc.count++
			offset = rc.offset
			return
		}
		// must split the refcount
		panic("TODO found, must split")
		return
	}

	// not found, allocate in the eof
	offset, s.eof = s.eof, s.eof+length
	value := &refcount{offset, length, 1}
	element := s.refcounts.PushBack(value)
	s.offsets[offset] = element
	return
}

// Storage.Count().
func (s *MemStorage) Count(offset, length, count int64) {
	s.m.Lock()
	defer s.m.Unlock()

	// TODO
	return
}
