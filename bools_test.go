package tmpdb

import (
	"bytes"
	"runtime"
	"testing"
)

func TestBoolClass(t *testing.T) {
	obj := &Bool{}
	class := obj.Class()
	if _, ok := class.(BoolClass); !ok {
		t.Error(ok)
	}
	if s := class.String(); s != "bool" {
		t.Error(s)
	}
	if l := class.length(); l != 1 {
		t.Error(l)
	}
}

func TestBoolValues(t *testing.T) {
	file := &MemFile{}
	db := NewDatabase(file)
	if obj1, err := db.NewBool(false); err != nil {
		t.Error(err)
	} else if b, err := obj1.Value(); err != nil {
		t.Error(err)
	} else if b {
		t.Error(b)
	}
	if obj2, err := db.NewBool(true); err != nil {
		t.Error(err)
	} else if b, err := obj2.Value(); err != nil {
		t.Error(err)
	} else if !b {
		t.Error(b)
	}
}

func TestBoolDealloc(t *testing.T) {
	file := &MemFile{}
	db := NewDatabase(file)
	{
		if _, err := db.NewBool(false); err != nil {
			t.Error(err)
		}
		if !bytes.Equal(file.data, []byte{0}) {
			t.Error(file.data)
		}
		runtime.GC()
	}
	{
		if _, err := db.NewBool(true); err != nil {
			t.Error(err)
		}
		if !bytes.Equal(file.data, []byte{0, 1}) {
			t.Error(file.data)
		}
		runtime.GC()
	}
	if !bytes.Equal(file.data, []byte{0, 1}) {
		t.Error(file.data)
	}
}
