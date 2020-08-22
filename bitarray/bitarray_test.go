package bitarray

import (
	"testing"
)

func TestBitarray(t *testing.T) {
	checkInitBySize := func(t *testing.T, barr Bitarray, size int) {
		t.Helper()
		storage_size := len(barr.storage)
		if size != storage_size {
			t.Errorf("size = %d, but storage_size=%d", size, storage_size)
		}

	}
	t.Run("InitBySize", func(t *testing.T) {
		barr := Bitarray{}
		barr.InitilizeBySize(1000)
		checkInitBySize(t, barr, 16)
	})
}
