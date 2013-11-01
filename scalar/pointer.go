package tmpdb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/daniel-fanjul-alcuten/tmpdb"
)

const OffsetLength = 64 / 8

// A pointer type. The value is the offset of the referenced Object.
type PointerClass struct {
	Base Class
}

// 8.
func (c PointerClass) Length() int64 {
	return OffsetLength
}

// Class.References()
func (c PointerClass) References(data []byte) (refs []Reference) {
	var offset int64
	reader := bytes.NewReader(data)
	if err := binary.Read(reader, binary.BigEndian, &offset); err != nil {
		panic(err)
	}
	refs = []Reference{Reference{Location{0, OffsetLength},
		Location{offset, c.Base.Length()}}}
	return
}

// A Pointer.
func (c PointerClass) NewObject(blob *Blob) Object {
	return Pointer{blob, blob.Class.(PointerClass)}
}

// "*bool", i.e.
func (c PointerClass) String() string {
	return fmt.Sprintf("*%s", c.Base)
}

// It invokes WriteBlob().
func WritePointer(storage Storage, object Object) (pointer Pointer, err error) {
	oblob := object.Blob()
	var buffer bytes.Buffer
	if err = binary.Write(&buffer, binary.BigEndian, oblob.Offset); err != nil {
		return
	}
	class := PointerClass{oblob.Class}
	blob, _, err := WriteBlob(storage, buffer.Bytes(), class)
	pointer = Pointer{blob, class}
	return
}

// A Pointer, the offset of the referenced Object.
type Pointer struct {
	blob  *Blob
	class PointerClass
}

// Object.Blob()
func (p Pointer) Blob() *Blob {
	return p.blob
}

// The offset of the referenced Object.
func (p Pointer) Offset(storage Storage) (offset int64, err error) {
	reader := p.blob.NewReader(storage)
	err = binary.Read(reader, binary.BigEndian, &offset)
	return
}

// The referenced Object.
func (p Pointer) Value(storage Storage) (object Object, err error) {
	offset, err := p.Offset(storage)
	if err != nil {
		return
	}
	base := p.class.Base
	blob := NewBlob(storage, offset, base)
	object = base.NewObject(blob)
	return
}
