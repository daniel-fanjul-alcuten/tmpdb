package tmpdb

import (
	"io"
	"testing"
)

func TestMemFile(t *testing.T) {
	file := &MemFile{}
	file.Data = []byte{1, 2}
	var reader io.ReaderAt = file
	var writer io.WriterAt = file

	p := make([]byte, 2)
	if n, err := reader.ReadAt(p, 1); err != io.ErrUnexpectedEOF {
		t.Error(err)
	} else {
		if n != 1 {
			t.Error(n)
		}
		if p[0] != 2 {
			t.Error(p[1])
		}
	}

	if n, err := writer.WriteAt([]byte{3, 4}, 1); err != nil {
		t.Error(err)
	} else if n != 2 {
		t.Error(n)
	}

	p = make([]byte, 2)
	if n, err := reader.ReadAt(p, 0); err != nil {
		t.Error(err)
	} else if n != 2 {
		t.Error(n)
	} else if p[0] != 1 {
		t.Error(p[0])
	} else if p[1] != 3 {
		t.Error(p[1])
	}

	p = make([]byte, 2)
	if n, err := reader.ReadAt(p, 1); err != nil {
		t.Error(err)
	} else if n != 2 {
		t.Error(n)
	} else if p[0] != 3 {
		t.Error(p[0])
	} else if p[1] != 4 {
		t.Error(p[1])
	}
}
