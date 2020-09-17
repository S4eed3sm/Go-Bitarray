package bitarray

import (
	"testing"
)

var testString string = "11110101010101011110110100010100101010101000001011110010101011101010111110101011111011010001010010101010100000101111001010101110101011111010101111111111010101001010000001101001010100111010101010111010010101001111111111111111111111110101001010101001111111111111001010101001"

func TestBitarray(t *testing.T) {
	checkInitBySize := func(t *testing.T, b *Bitarray, size int) {
		t.Helper()
		storagesize := len(b.storage)
		if size != storagesize {
			t.Errorf("size = %d, but storagesize=%d\n", size, storagesize)
		}

	}

	checkInitByStrValue := func(t *testing.T, b *Bitarray, storage []uint64) {
		t.Helper()
		size := len(b.storage)
		if len(storage) != size {
			t.Errorf("size = %d, but storagesize=%d\n", size, len(storage))
		}
		for i := 0; i < size; i++ {
			if b.storage[i] != storage[i] {
				t.Errorf("b.storage[%d]=%d != storage=%d\n", i, b.storage[i], storage[i])
			}
		}
	}

	checkInitByValue := func(t *testing.T, b *Bitarray, val uint64) {
		t.Helper()
		if len(b.storage) != 1 {
			t.Errorf("in InitByValue size is alway equal to 1, but now is[%d]", len(b.storage))
		}

		if val != b.storage[0] {
			t.Errorf("val[%d] != b.storage[0]{%d}", val, b.storage[0])
		}
	}

	checkGet := func(t *testing.T, b *Bitarray, index uint64, val uint8) {
		t.Helper()
		bVal, err := b.Get(index)
		if err != nil {
			t.Errorf("%s\n", err.Error())
		}
		if bVal != val {
			t.Errorf("b.storage[%d]=%d != val=%d\n", index, bVal, val)
		}

	}

	checkSet := func(t *testing.T, b *Bitarray, index uint64, val uint8) {
		t.Helper()
		err := b.Set(index, val)
		if err != nil {
			t.Errorf("%s\n", err.Error())
		}

		bVal, err1 := b.Get(index)

		if err1 != nil {
			t.Errorf("%s\n", err1.Error())
		}
		if bVal != val {
			t.Errorf("index=%d: setVal=%d != getVal=%d", index, val, bVal)
		}
	}

	checkCompare := func(t *testing.T, b1 *Bitarray, b2 *Bitarray, result bool) {
		t.Helper()
		r := b1.Compare(b2)
		if r != result {
			t.Errorf("resultOfCompare[%v] != givenResult[%v]\n", r, result)
		}
	}

	checkToString := func(t *testing.T, b *Bitarray, val string) {
		t.Helper()
		s := b.ToString()
		if *s != val {
			t.Errorf("ToString[%s] != %s", *s, val)
		}
	}

	checkAnd := func(t *testing.T, b1 *Bitarray, b2 *Bitarray, result *Bitarray) {
		t.Helper()
		res := b1.And(b2)
		if res.Compare(result) == false {
			t.Errorf("AndRes[%s] != result[%s]", *res.ToString(), *result.ToString())
		}
	}

	checkOr := func(t *testing.T, b1 *Bitarray, b2 *Bitarray, result *Bitarray) {
		t.Helper()
		res := b1.Or(b2)
		if res.Compare(result) == false {
			t.Errorf("OrRes[%s] != result[%s]", *res.ToString(), *result.ToString())
		}
	}

	checkXor := func(t *testing.T, b1 *Bitarray, b2 *Bitarray, result *Bitarray) {
		t.Helper()
		res := b1.Xor(b2)
		if res.Compare(result) == false {
			t.Errorf("XorRes[%s] != result[%s]", *res.ToString(), *result.ToString())
		}
	}

	checkNot := func(t *testing.T, b *Bitarray, result *Bitarray) {
		t.Helper()
		res := b.Not()
		resAnd := b.And(res)
		if *resAnd.ToString() != *result.ToString() {
			t.Errorf("NotRes[%s] != result[%s]", *res.ToString(), *result.ToString())
		}
	}

	checkShiftLeft := func(t *testing.T, b *Bitarray, shiftCount int, result *Bitarray) {
		t.Helper()
		res := b.ShiftLeft(shiftCount)
		if res.Compare(result) == false {
			t.Errorf("res[%s] != result[%s]", *res.ToString(), *result.ToString())
		}
	}

	t.Run("TestInitBySize", func(t *testing.T) {
		b := Bitarray{}
		b.InitializeBySize(1000)
		checkInitBySize(t, &b, 16)
	})

	t.Run("TestInitByeStrValue", func(t *testing.T) {
		b := Bitarray{}
		b.InitializeByStrValue(testString)
		checkInitByStrValue(t, &b, []uint64{0b1111010101010101, 0b1110110100010100101010101000001011110010101011101010111110101011, 0b1110110100010100101010101000001011110010101011101010111110101011,
			0b1111111101010100101000000110100101010011101010101011101001010100, 0b1111111111111111111111110101001010101001111111111111001010101001})

	})

	t.Run("TestInitByValue", func(t *testing.T) {
		b := Bitarray{}
		value := uint64(0b1100010101011001)
		b.InitializeByValue(value)
		checkInitByValue(t, &b, value)

	})

	t.Run("TestGet", func(t *testing.T) {
		b := Bitarray{}
		b.InitializeByStrValue(testString)
		testStringRev := ReverseString(testString)
		checkGet(t, &b, 27, 1)
		checkGet(t, &b, 87, 1)
		checkGet(t, &b, 86, 0)
		for i, v := range testStringRev {
			checkGet(t, &b, uint64(i), uint8(v-'0'))
		}

	})

	t.Run("TestSet", func(t *testing.T) {
		b := Bitarray{}
		testStringRev := ReverseString(testString)
		for i, v := range testStringRev {
			checkSet(t, &b, uint64(i), uint8(v-'0'))
		}
	})

	t.Run("TestCompare", func(t *testing.T) {
		b := Bitarray{}
		b2 := Bitarray{}
		b3 := Bitarray{}
		b4 := Bitarray{}

		testStringRev := ReverseString(testString)
		for i, v := range testStringRev {
			b.Set(uint64(i), uint8(v-'0'))
		}

		b2.InitializeByStrValue(testString)

		b3.InitializeByStrValue(testString)

		b4.InitializeByBitarray(&b)

		checkCompare(t, &b, &b2, true)
		checkCompare(t, &b, &b3, true)
		checkCompare(t, &b, &b4, true)
		b2.Set(2, 1)
		checkCompare(t, &b, &b2, false)
	})

	t.Run("TestToString", func(t *testing.T) {
		b := Bitarray{}
		b.InitializeByStrValue(testString)
		checkToString(t, &b, testString)
	})

	t.Run("TestAnd", func(t *testing.T) {
		b1 := Bitarray{}
		b2 := Bitarray{}
		result := Bitarray{}
		b1.InitializeByStrValue(testString)
		b2.InitializeByStrValue(ReverseString(testString))
		resStr := "10010101010001011110110100010100000010101000001011110010101011100010101000001001010001010000000010000010000000000010001010101110100001011010000101110101010001000000000001000001000000001010001010010000010101000111010101001111010000010101000000101000101101111010001010101001"
		result.InitializeByStrValue(resStr)
		checkAnd(t, &b1, &b2, &result)
	})

	t.Run("TestOr", func(t *testing.T) {
		b1 := Bitarray{}
		b2 := Bitarray{}
		result := Bitarray{}
		b1.InitializeByStrValue(testString)
		b2.InitializeByStrValue(ReverseString(testString))
		resStr := "11110101010111111111111110010101111010101111111111111111111111111010111111111111111111011101111010111110100001111111101011111111111111111111111111111111010111111110000101111101011110111011111111111111111101011111111111111111111111110101011110101001111111111111101010101111"
		result.InitializeByStrValue(resStr)
		checkOr(t, &b1, &b2, &result)
	})

	t.Run("TestXor", func(t *testing.T) {
		b1 := Bitarray{}
		b2 := Bitarray{}
		result := Bitarray{}
		b1.InitializeByStrValue(testString)
		b2.InitializeByStrValue(ReverseString(testString))
		resStr := "1100000000110100001001010000001111000000111110100001101010100011000010111110110101110001101111000111100100001111101100001010001011110100101111010001010000110111110000100111100011110110001110101101111101000011000101010110000101111100000011110000001010010000101100000000110"
		result.InitializeByStrValue(resStr)
		checkXor(t, &b1, &b2, &result)
	})

	t.Run("TestNot", func(t *testing.T) {
		b1 := Bitarray{}
		b1.InitializeByStrValue(testString)
		result := Bitarray{}
		result.InitializeBySize(1)
		checkNot(t, &b1, &result)
	})

	t.Run("TestShiftLeft", func(t *testing.T) {
		b := Bitarray{}
		result := Bitarray{}

		b.InitializeByStrValue(testString)
		//testString << 127
		resStr := "111101010101010111101101000101001010101010000010111100101010111010101111101010111110110100010100101010101000001011110010101011101010111110101011111111110101010010100000011010010101001110101010101110100101010011111111111111111111111101010010101010011111111111110010101010010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
		result.InitializeByStrValue(resStr)
		checkShiftLeft(t, &b, 127, &result)
	})
}
