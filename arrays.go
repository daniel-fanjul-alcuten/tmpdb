package tmpdb

import (
	"fmt"
	"runtime"
)

// An array type: continguous list of values of the same type.
type ArrayClass struct {
	Size int64
	Base Class
}

type Array struct {
	_db     *Database
	_offset int64
	class   ArrayClass
}

// "[3]bool", i.e.
func (c ArrayClass) String() string {
	return fmt.Sprintf("[%d]%s", c.Size, c.Base)
}

func (c ArrayClass) length() int64 {
	return c.Size * c.Base.length()
}

func writeArray(db *Database, class Class, values ...Object) (obj *Array, err error) {
	size, length := int64(len(values)), class.length()
	data := make([]byte, 0, size*length)
	for _, value := range values {
		if value.Class() != class {
			err = fmt.Errorf("Unexpected class %s", value.Class())
			return
		}
		var vdata []byte
		vdata, err = value.db().read(value.offset(), length)
		if err != nil {
			return
		}
		data = append(data, vdata...)
	}
	offset, err := db.write(data)
	if err != nil {
		return
	}
	obj = &Array{db, offset, ArrayClass{size, class}}
	runtime.SetFinalizer(obj, func(obj *Array) {
		obj._db.finalize(obj._offset, obj.class.length())
	})
	return
}

func writeBoolArray(db *Database, values ...bool) (obj *Array, err error) {
	class := BoolClass{}
	size, length := int64(len(values)), class.length()
	data := make([]byte, 0, size*length)
	for _, value := range values {
		if value {
			data = append(data, 1)
		} else {
			data = append(data, 0)
		}
	}
	offset, err := db.write(data)
	if err != nil {
		return
	}
	obj = &Array{db, offset, ArrayClass{size, class}}
	runtime.SetFinalizer(obj, func(obj *Array) {
		obj._db.finalize(obj._offset, obj.class.length())
	})
	return
}

func (c ArrayClass) newObject(db *Database, offset int64) Object {
	obj := &Array{db, offset, c}
	runtime.SetFinalizer(obj, func(obj *Bool) {
		obj._db.finalize(obj._offset, 1)
	})
	return obj
}

// The ArrayClass.
func (o *Array) Class() Class {
	return o.class
}

func (o *Array) db() *Database {
	return o._db
}

func (o *Array) offset() int64 {
	return o._offset
}

// One value.
func (o *Array) Value(i int64) (obj Object) {
	if i >= o.class.Size {
		panic("array index out of range")
	}
	class := o.class.Base
	length := class.length()
	off := o._offset + i*length
	obj = class.newObject(o._db, off)
	return
}

// All values.
func (o *Array) Values() (objs []Object) {
	class := o.class.Base
	length := class.length()
	off := o._offset
	for i := int64(0); i < o.class.Size; i++ {
		obj := class.newObject(o._db, off)
		objs = append(objs, obj)
		off += length
	}
	return
}
