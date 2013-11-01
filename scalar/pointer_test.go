package tmpdb

import (
	. "github.com/daniel-fanjul-alcuten/tmpdb"
	"testing"
)

func TestPointerClass(t *testing.T) {
	pointerClass := PointerClass{BoolClass{}}
	var class Class = pointerClass
	if l := class.Length(); l != 8 {
		t.Error(l)
	}
	if r := class.References([]byte{0, 0, 0, 0, 0, 0, 1, 2}); len(r) != 1 {
		t.Error(r)
	} else if r0 := r[0]; r0 != (Reference{Location{0, 8}, Location{1*256 + 2, 1}}) {
		t.Error(r0)
	}
	if s := pointerClass.String(); s != "*bool" {
		t.Error(s)
	}
}

func TestPointerValue(t *testing.T) {
	storage := &DummyStorage{}
	boool, err := WriteBool(storage, false)
	if err != nil {
		t.Error(err)
	}
	if pointer, err := WritePointer(storage, boool); err != nil {
		t.Error(err)
	} else {
		if offset, err := pointer.Offset(storage); err != nil {
			t.Error(err)
		} else if offset != 0 {
			t.Error(offset)
		}
		if value, err := pointer.Value(storage); err != nil {
			t.Error(err)
		} else {
			if boool, ok := value.(Bool); !ok {
				t.Error(ok)
			} else if b, err := boool.Value(storage); err != nil {
				t.Error(err)
			} else if b {
				t.Error(b)
			}
		}
	}
	boool, err = WriteBool(storage, true)
	if err != nil {
		t.Error(err)
	}
	if pointer, err := WritePointer(storage, boool); err != nil {
		t.Error(err)
	} else {
		if offset, err := pointer.Offset(storage); err != nil {
			t.Error(err)
		} else if offset != 9 {
			t.Error(offset)
		}
		if value, err := pointer.Value(storage); err != nil {
			t.Error(err)
		} else {
			if boool, ok := value.(Bool); !ok {
				t.Error(ok)
			} else if b, err := boool.Value(storage); err != nil {
				t.Error(err)
			} else if !b {
				t.Error(b)
			}
		}
	}
}
