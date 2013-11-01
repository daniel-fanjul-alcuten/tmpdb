package tmpdb

import (
	"bytes"
	"fmt"
	. "github.com/daniel-fanjul-alcuten/tmpdb"
)

// An array type: contiguous list of values of the same Class.
type ArrayClass struct {
	Size int64
	Base Class
}

// c.Size * c.Base.Length()
func (c ArrayClass) Length() int64 {
	return c.Size * c.Base.Length()
}

// Class.References()
func (c ArrayClass) References(data []byte) (refs []Reference) {
	offset, length := int64(0), c.Base.Length()
	for i := int64(0); i < c.Size; i++ {
		for _, ref := range c.Base.References(data[offset : offset+length]) {
			ref.Source.Offset += offset
			refs = append(refs, ref)
		}
		offset += length
	}
	return
}

// Class.NewObject()
func (c ArrayClass) NewObject(blob *Blob) Object {
	return Array{blob, blob.Class.(ArrayClass)}
}

// "[3]bool", i.e.
func (c ArrayClass) String() string {
	return fmt.Sprintf("[%d]%s", c.Size, c.Base)
}

// It invokes WriteBlob(). All Objects must belong to the given Class.
func WriteArray(storage Storage, base Class, objects ...Object) (array Array, err error) {
	size := int64(len(objects))
	buffer := bytes.NewBuffer(make([]byte, 0, size*base.Length()))
	for _, object := range objects {
		blob := object.Blob()
		if blob.Class != base {
			err = fmt.Errorf("Unexpected class %s", blob.Class)
			return
		}
		if _, err = buffer.ReadFrom(blob.NewReader(storage)); err != nil {
			return
		}
	}
	class := ArrayClass{size, base}
	blob, _, err := WriteBlob(storage, buffer.Bytes(), class)
	array = Array{blob, class}
	return
}

// An array, a contiguous list of values of the same Class.
type Array struct {
	blob  *Blob
	class ArrayClass
}

// Object.Blob()
func (a Array) Blob() *Blob {
	return a.blob
}

// The number of objects.
func (a Array) Size() (size int64) {
	size = a.class.Size
	return
}

// One value.
func (a Array) Value(storage Storage, i int64) (object Object) {
	if i >= a.class.Size {
		panic("array index out of range")
	}
	base := a.class.Base
	offset := a.blob.Offset + i*base.Length()
	blob := NewBlob(storage, offset, base)
	object = base.NewObject(blob)
	return
}

// All values.
func (a Array) Values(storage Storage) (objects []Object) {
	base := a.class.Base
	offset, length := a.blob.Offset, base.Length()
	for i := int64(0); i < a.class.Size; i++ {
		object := base.NewObject(NewBlob(storage, offset, base))
		objects = append(objects, object)
		offset += length
	}
	return
}
