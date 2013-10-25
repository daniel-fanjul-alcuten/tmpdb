package tmpdb

import (
	"io"
)

// No thread-safe.
type Database struct {
	file   File
	allocs AllocationPool
	counts RefCountPool
}

// New instance.
func NewDatabase(file File) *Database {
	return &Database{file, &DummyAllocationPool{}, &DummyRefCountPool{}}
}

// Creates a new Bool Object.
func (db *Database) NewBool(value bool) (obj *Bool, err error) {
	return writeBool(db, value)
}

// Creates a new Array Object.
func (db *Database) NewArray(class Class, values ...Object) (obj *Array, err error) {
	return writeArray(db, class, values...)
}

// Creates a new BoolArray Object.
func (db *Database) NewBoolArray(values ...bool) (obj *Array, err error) {
	return writeBoolArray(db, values...)
}

func (db *Database) read(offset, length int64) (data []byte, err error) {
	data = make([]byte, length)
	n, err := db.file.ReadAt(data, offset)
	if n == len(data) && err == io.EOF {
		err = nil
	}
	return
}

func (db *Database) write(data []byte) (offset int64, err error) {
	length := int64(len(data))
	offset = db.allocs.Allocate(length)
	if _, err = db.file.WriteAt(data, offset); err != nil {
		db.allocs.Deallocate(offset, length)
		return
	}
	db.counts.Increment(offset, 1)
	return
}

func (db *Database) finalize(offset, length int64) {
	if db.counts.Increment(offset, -1) == 0 {
		db.allocs.Deallocate(offset, length)
		// TODO recursive deallocation
	}
}
