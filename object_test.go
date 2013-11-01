package tmpdb

import (
	"bytes"
	"io"
	"testing"
)

type TestClass struct{}

func (c TestClass) Length() int64 {
	return 1
}

func (c TestClass) References(data []byte) []Reference {
	return nil
}

func (c TestClass) NewObject(blob *Blob) Object {
	return nil
}

func TestBlobNewReader(t *testing.T) {
	file := &MemFile{}
	file.Data = []byte{2, 3}
	blob := &Blob{1, TestClass{}}
	reader := blob.NewReader(file)
	data := []byte{0}
	if n, err := io.ReadFull(reader, data); err != nil {
		t.Error(err)
	} else {
		if n != 1 {
			t.Error(n)
		}
		if !bytes.Equal(data, []byte{3}) {
			t.Error(data)
		}
	}
}

func TestNewBlob(t *testing.T) {
	class, storage := TestClass{}, &DummyStorage{}
	blob := NewBlob(storage, 1, class)
	if offset := blob.Offset; offset != 1 {
		t.Error(offset)
	}
	if oclass := blob.Class; oclass != class {
		t.Error(oclass)
	}
}

func TestWriteBlob(t *testing.T) {
	class, storage := TestClass{}, &DummyStorage{Offset: 2}
	if blob, n, err := WriteBlob(storage, []byte{3}, class); err != nil {
		t.Error(err)
	} else {
		if offset := blob.Offset; offset != 2 {
			t.Error(offset)
		}
		if oclass := blob.Class; oclass != class {
			t.Error(oclass)
		}
		if n != 1 {
			t.Error(n)
		}
	}
}
