package tmpdb

import (
	"bytes"
	. "github.com/daniel-fanjul-alcuten/tmpdb"
	"io"
)

// The bool type.
type BoolClass struct{}

// 1.
func (c BoolClass) Length() int64 {
	return 1
}

// nil.
func (c BoolClass) References(data []byte) []Reference {
	return nil
}

// A Bool.
func (c BoolClass) NewObject(blob *Blob) Object {
	return Bool{blob}
}

// "bool".
func (c BoolClass) String() string {
	return "bool"
}

// It invokes WriteBlob().
func WriteBool(storage Storage, value bool) (object Bool, err error) {
	data := []byte{0}
	if value {
		data[0] = 1
	}
	blob, _, err := WriteBlob(storage, data, BoolClass{})
	object = Bool{blob}
	return
}

// A bool.
type Bool struct {
	blob *Blob
}

// Object.Blob()
func (b Bool) Blob() *Blob {
	return b.blob
}

// The bool value.
func (b Bool) Value(reader io.ReaderAt) (value bool, err error) {
	var buffer bytes.Buffer
	if _, err = buffer.ReadFrom(b.blob.NewReader(reader)); err != nil {
		return
	}
	var c byte
	if c, err = buffer.ReadByte(); err != nil {
		return
	}
	value = c > 0
	return
}
