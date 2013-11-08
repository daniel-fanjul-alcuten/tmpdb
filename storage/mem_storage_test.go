package tmpdb

import (
	. "github.com/daniel-fanjul-alcuten/tmpdb"
	"testing"
)

func TestMemStorageInterface(t *testing.T) {
	memFile := &MemFile{}
	memStorage := NewMemStorage(memFile)
	var _ Storage = memStorage
}
