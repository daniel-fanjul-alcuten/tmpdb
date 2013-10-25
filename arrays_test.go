package tmpdb

import (
	"bytes"
	"runtime"
	"testing"
)

func TestArrayClass(t *testing.T) {
	obj := &Array{class: ArrayClass{2, BoolClass{}}}
	class := obj.Class()
	if _, ok := class.(ArrayClass); !ok {
		t.Error(ok)
	}
	if s := class.String(); s != "[2]bool" {
		t.Error(s)
	}
	if l := class.length(); l != 2 {
		t.Error(l)
	}
}

func TestArrayValues(t *testing.T) {
	file := &MemFile{}
	db := NewDatabase(file)
	b1, err := db.NewBool(false)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := db.NewBool(true)
	if err != nil {
		t.Fatal(err)
	}
	obj, err := db.NewArray(b1.Class(), b1, b2)
	if err != nil {
		t.Error(err)
	}
	if c := obj.Class(); c != (ArrayClass{2, BoolClass{}}) {
		t.Error(c)
	}
	if v, ok := obj.Value(0).(*Bool); !ok {
		t.Error(err)
	} else if b, err := v.Value(); err != nil {
		t.Error(err)
	} else if b {
		t.Error(b)
	}
	if v, ok := obj.Value(1).(*Bool); !ok {
		t.Error(err)
	} else if b, err := v.Value(); err != nil {
		t.Error(err)
	} else if !b {
		t.Error(b)
	}
	if vv := obj.Values(); len(vv) != 2 {
		t.Error(len(vv))
	} else {
		if v, ok := vv[0].(*Bool); !ok {
			t.Error(err)
		} else if b, err := v.Value(); err != nil {
			t.Error(err)
		} else if b {
			t.Error(b)
		}
		if v, ok := vv[1].(*Bool); !ok {
			t.Error(err)
		} else if b, err := v.Value(); err != nil {
			t.Error(err)
		} else if !b {
			t.Error(b)
		}
	}
}

func TestArrayDealloc(t *testing.T) {
	file := &MemFile{}
	db := NewDatabase(file)
	var b1, b2 *Bool
	{
		var err error
		if b1, err = db.NewBool(false); err != nil {
			t.Fatal(err)
		}
		if b2, err = db.NewBool(true); err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(file.data, []byte{0, 1}) {
			t.Error(file.data)
		}
	}
	{
		if _, err := db.NewArray(b1.Class(), b1); err != nil {
			t.Error(err)
		}
		if !bytes.Equal(file.data, []byte{0, 1, 0}) {
			t.Error(file.data)
		}
		runtime.GC()
	}
	{
		if _, err := db.NewArray(b2.Class(), b2); err != nil {
			t.Error(err)
		}
		if !bytes.Equal(file.data, []byte{0, 1, 0, 1}) {
			t.Error(file.data)
		}
		runtime.GC()
	}
	if !bytes.Equal(file.data, []byte{0, 1, 0, 1}) {
		t.Error(file.data)
	}
}
