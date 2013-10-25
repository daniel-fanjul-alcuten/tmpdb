package tmpdb

import (
	"runtime"
)

// The bool type.
type BoolClass struct{}

// A bool value.
type Bool struct {
	_db     *Database
	_offset int64
}

// "bool".
func (c BoolClass) String() string {
	return "bool"
}

func (c BoolClass) length() int64 {
	return 1
}

func writeBool(db *Database, value bool) (obj *Bool, err error) {
	data := []byte{0}
	if value {
		data[0] = 1
	}
	offset, err := db.write(data)
	if err != nil {
		return
	}
	obj = &Bool{db, offset}
	runtime.SetFinalizer(obj, func(obj *Bool) {
		obj._db.finalize(obj._offset, 1)
	})
	return
}

func (c BoolClass) newObject(db *Database, offset int64) Object {
	obj := &Bool{db, offset}
	runtime.SetFinalizer(obj, func(obj *Bool) {
		obj._db.finalize(obj._offset, 1)
	})
	return obj
}

// BoolClass{}.
func (o *Bool) Class() Class {
	return BoolClass{}
}

func (o *Bool) db() *Database {
	return o._db
}

func (o *Bool) offset() int64 {
	return o._offset
}

// The bool value.
func (o *Bool) Value() (value bool, err error) {
	data, err := o._db.read(o._offset, 1)
	if err != nil {
		return
	}
	value = data[0] > 0
	return
}
