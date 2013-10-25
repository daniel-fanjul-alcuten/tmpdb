package tmpdb

import (
	"io"
	"testing"
)

func TestMemFile(t *testing.T) {
	memFile := &MemFile{}
	memFile.data = []byte{1, 2}
	var file File = memFile

	p := make([]byte, 2)
	if n, err := file.ReadAt(p, 1); err != io.ErrUnexpectedEOF {
		t.Error(err)
	} else {
		if n != 1 {
			t.Error(n)
		}
		if p[0] != 2 {
			t.Error(p[1])
		}
	}

	if n, err := file.WriteAt([]byte{3, 4}, 1); err != nil {
		t.Error(err)
	} else if n != 2 {
		t.Error(n)
	}

	p = make([]byte, 2)
	if n, err := file.ReadAt(p, 0); err != nil {
		t.Error(err)
	} else if n != 2 {
		t.Error(n)
	} else if p[0] != 1 {
		t.Error(p[0])
	} else if p[1] != 3 {
		t.Error(p[1])
	}

	p = make([]byte, 2)
	if n, err := file.ReadAt(p, 1); err != nil {
		t.Error(err)
	} else if n != 2 {
		t.Error(n)
	} else if p[0] != 3 {
		t.Error(p[0])
	} else if p[1] != 4 {
		t.Error(p[1])
	}
}
