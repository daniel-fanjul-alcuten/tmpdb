package tmpdb

import (
	"testing"
)

func TestDummyStorage(t *testing.T) {
	dummy := &DummyStorage{}
	var storage Storage = dummy
	if off := storage.Allocate(1, nil); off != 0 {
		t.Error(off)
	}
	if off := storage.Allocate(1, nil); off != 1 {
		t.Error(off)
	}
	if off := storage.Allocate(2, nil); off != 2 {
		t.Error(off)
	}
	if off := storage.Allocate(3, nil); off != 4 {
		t.Error(off)
	}
	if off := storage.Allocate(5, nil); off != 7 {
		t.Error(off)
	}
	if off := storage.Allocate(8, nil); off != 12 {
		t.Error(off)
	}
	if off := storage.Allocate(13, nil); off != 20 {
		t.Error(off)
	}
	if off := storage.Allocate(21, nil); off != 33 {
		t.Error(off)
	}
}
