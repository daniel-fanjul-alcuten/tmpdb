package tmpdb

import (
	. "github.com/daniel-fanjul-alcuten/tmpdb"
	. "github.com/daniel-fanjul-alcuten/tmpdb/scalar"
	"testing"
)

func TestArrayClass_Bool(t *testing.T) {
	arrayClass := ArrayClass{2, BoolClass{}}
	var class Class = arrayClass
	if l := class.Length(); l != 2 {
		t.Error(l)
	}
	if r := class.References([]byte{0, 0}); r != nil {
		t.Error(r)
	}
	if s := arrayClass.String(); s != "[2]bool" {
		t.Error(s)
	}
}

func TestArrayClass_PointerBool(t *testing.T) {
	arrayClass := ArrayClass{2, PointerClass{BoolClass{}}}
	var class Class = arrayClass
	if l := class.Length(); l != 16 {
		t.Error(l)
	}
	if r := class.References([]byte{0, 0, 0, 0, 0, 0, 1, 2, 0, 0, 0, 0, 0, 0, 3, 4}); len(r) != 2 {
		t.Error(r)
	} else if r0 := r[0]; r0 != (Reference{Location{0, 8}, Location{1*256 + 2, 1}}) {
		t.Error(r0)
	} else if r1 := r[1]; r1 != (Reference{Location{8, 8}, Location{3*256 + 4, 1}}) {
		t.Error(r1)
	}
	if s := arrayClass.String(); s != "[2]*bool" {
		t.Error(s)
	}
}

func TestArrayValues(t *testing.T) {
	storage := &DummyStorage{}
	b1, err := WriteBool(storage, false)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := WriteBool(storage, true)
	if err != nil {
		t.Fatal(err)
	}
	array, err := WriteArray(storage, BoolClass{}, b1, b2)
	if err != nil {
		t.Error(err)
	}
	if s := array.Size(); s != 2 {
		t.Error(s)
	}
	if v, ok := array.Value(storage, 0).(Bool); !ok {
		t.Error(ok)
	} else if b, err := v.Value(storage); err != nil {
		t.Error(err)
	} else if b {
		t.Error(b)
	}
	if v, ok := array.Value(storage, 1).(Bool); !ok {
		t.Error(ok)
	} else if b, err := v.Value(storage); err != nil {
		t.Error(err)
	} else if !b {
		t.Error(b)
	}
	if values := array.Values(storage); len(values) != 2 {
		t.Error(len(values))
	} else {
		if v, ok := values[0].(Bool); !ok {
			t.Error(ok)
		} else if b, err := v.Value(storage); err != nil {
			t.Error(err)
		} else if b {
			t.Error(b)
		}
		if v, ok := values[1].(Bool); !ok {
			t.Error(ok)
		} else if b, err := v.Value(storage); err != nil {
			t.Error(err)
		} else if !b {
			t.Error(b)
		}
	}
}
