package tmpdb

import (
	"io"
)

// A seekable file. An *os.File, for example. Implementations must be
// thread-safe.
type File interface {
	io.ReaderAt
	io.WriterAt
}

// The location of a blob.
type Location struct {
	Offset, Length int64
}

// The source blob holds a reference to the target blob. The target blob must
// not be freed until the source blob has been freed.
type Reference struct {
	Source, Target Location
}

// It counts the number of references to the blobs. Implementations must be
// thread-safe.
type Storage interface {

	// The reads should be able to run in parallel.
	io.ReaderAt

	// The writes should be able to run in parallel.
	io.WriterAt

	// It finds an unused blob of the given length and marks it with the initial
	// reference count of 1. The offsets of the source blobs are relative to the
	// offset of the whole allocation. The reference count of the target blobs
	// are incremented by 1.
	Allocate(length int64, references []Reference) (offset int64)

	// It updates the counter of the references to the blobs. When a source blob
	// of any reference of the previous allocations becomes unreferenced, the
	// target blob decrements its counter. When a counter reaches zero, the blob
	// is marked as unused. This method is typically invoked by runtime
	// finalizers so it must meet their requirements.
	Count(offset, length, count int64)
}
