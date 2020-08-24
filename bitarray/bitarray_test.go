package bitarray

import (
	"testing"
)

func TestBitarray(t *testing.T) {
	checkInitBySize := func(t *testing.T, barr Bitarray, size int) {
		t.Helper()
		storagesize := len(barr.storage)
		if size != storagesize {
			t.Errorf("size = %d, but storagesize=%d\n", size, storagesize)
		}

	}

	checkInitByStrValue := func(t *testing.T, barr Bitarray, storage []uint64) {
		t.Helper()
		size := len(barr.storage)
		if len(storage) != size {
			t.Errorf("size = %d, but storagesize=%d\n", size, len(storage))
		}
		for i := 0; i < size; i++ {
			if barr.storage[i] != storage[i] {
				t.Errorf("barr.storage[%d]=%d != storage=%d\n", i, barr.storage[i], storage[i])
			}
		}
	}

	t.Run("InitBySize", func(t *testing.T) {
		barr := Bitarray{}
		barr.InitilizeBySize(1000)
		checkInitBySize(t, barr, 16)
	})

	t.Run("InitByeStrValue", func(t *testing.T) {
		barr := Bitarray{}
		barr.InitilizeByStrValue("11110101010101011110110100010100101010101000001011110010101011101010111110101011111011010001010010101010100000101111001010101110101011111010101111111111010101001010000001101001010100111010101010111010010101001111111111111111111111110101001010101001111111111111001010101001")
		checkInitByStrValue(t, barr, []uint64{0b1111010101010101, 0b1110110100010100101010101000001011110010101011101010111110101011, 0b1110110100010100101010101000001011110010101011101010111110101011,
			0b1111111101010100101000000110100101010011101010101011101001010100, 0b1111111111111111111111110101001010101001111111111111001010101001})

	})
}
