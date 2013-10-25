package tmpdb

import (
	"io"
)

// A seekable file.
type File interface {
	io.ReaderAt
	io.WriterAt
}

// It keeps track of the used and unsed blobs.
// Implementations are not required to be thread-safe.
type AllocationPool interface {
	// Find an unused blob and mark it as used.
	Allocate(length int64) (offset int64)
	// Mark the blob as unused.
	Deallocate(offset int64, length int64)
}

// It counts references.
type RefCount int64

// It keeps track of the references to the blobs.
// Implementations are not required to be thread-safe.
type RefCountPool interface {
	// Adds or Subtracts to the counter.
	// Returns the new value of the counter.
	Increment(int64, RefCount) RefCount
}

// The type of any Object.
type Class interface {
	// The name of the class.
	String() string
	length() int64
	newObject(*Database, int64) Object
}

// Any object in the database.
type Object interface {
	// The type of the object.
	Class() Class
	db() *Database
	offset() int64
}
