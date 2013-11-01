package tmpdb

import (
	"io"
	"runtime"
)

// The type of the Blobs. Implementations must be thread-safe, immutable when
// possible.
type Class interface {
	// The length of the values of a Class is constant.
	Length() int64
	// The offsets of the source blobs are relative to the offset of the Blob.
	References(data []byte) []Reference
	// The proper Object implementation for the given Blob. The Class of the Blob
	// must be the same as the receiver.
	NewObject(blob *Blob) Object
}

// A reference to a value that has been stored in a Storage. On creation and
// finalization, an instance of Storage must be notified in order to update the
// reference counts. It must be immutable.
type Blob struct {
	Offset int64
	Class  Class
}

// An Object is the high level abstraction of a Blob, it allows to expand the
// method set.
type Object interface {
	Blob() *Blob
}

// It returns a SectionReader for the whole value.
func (b *Blob) NewReader(reader io.ReaderAt) *io.SectionReader {
	return io.NewSectionReader(reader, b.Offset, b.Class.Length())
}

// An instance of Blob is created, the reference count is incremented and the
// runtime finalizer set to decrement it.
func NewBlob(storage Storage, offset int64, class Class) (blob *Blob) {
	blob = &Blob{offset, class}
	storage.Count(offset, class.Length(), 1)
	runtime.SetFinalizer(blob, func(blob *Blob) {
		storage.Count(blob.Offset, blob.Class.Length(), -1)
	})
	return
}

// The space for the value is allocated in the Storage, an instance of Blob
// is created, the runtime finalizer set to decrement the reference count, and
// the data is written.
func WriteBlob(storage Storage, data []byte, class Class) (blob *Blob, n int, err error) {
	length, references := int64(len(data)), class.References(data)
	offset := storage.Allocate(length, references)
	blob = &Blob{offset, class}
	runtime.SetFinalizer(blob, func(blob *Blob) {
		storage.Count(blob.Offset, blob.Class.Length(), -1)
	})
	n, err = storage.WriteAt(data, offset)
	return
}
