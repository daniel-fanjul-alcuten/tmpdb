package tmpdb

import (
	. "github.com/daniel-fanjul-alcuten/tmpdb"
	"testing"
)

func TestBoolClass(t *testing.T) {
	boolClass := BoolClass{}
	var class Class = boolClass
	if l := class.Length(); l != 1 {
		t.Error(l)
	}
	if r := class.References(nil); r != nil {
		t.Error(r)
	}
	if s := boolClass.String(); s != "bool" {
		t.Error(s)
	}
}

func TestBoolValue(t *testing.T) {
	storage := &DummyStorage{}
	if boool, err := WriteBool(storage, false); err != nil {
		t.Error(err)
	} else {
		if value, err := boool.Value(storage); err != nil {
			t.Error(err)
		} else if value {
			t.Error(value)
		}
	}
	if boool, err := WriteBool(storage, true); err != nil {
		t.Error(err)
	} else {
		if value, err := boool.Value(storage); err != nil {
			t.Error(err)
		} else if !value {
			t.Error(value)
		}
	}
}
