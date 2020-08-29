package bitarray

import (
	"testing"

	utils "gitlab.com/S4eed3sm/bit-array-in-golang/utils"
)

var testString string = "11110101010101011110110100010100101010101000001011110010101011101010111110101011111011010001010010101010100000101111001010101110101011111010101111111111010101001010000001101001010100111010101010111010010101001111111111111111111111110101001010101001111111111111001010101001"

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

	checkGet := func(t *testing.T, barr Bitarray, index uint64, val uint8) {
		t.Helper()
		barrVal, err := barr.Get(index)
		if err != nil {
			t.Errorf("%s\n", err.Error())
		}
		if barrVal != val {
			t.Errorf("barr.storage[%d]=%d != val=%d\n", index, barrVal, val)
		}

	}

	checkSet := func(t *testing.T, barr Bitarray, index uint64, val uint8) {
		t.Helper()
		err := barr.Set(index, val)
		if err != nil {
			t.Errorf("%s\n", err.Error())
		}

		barrVal, err1 := barr.Get(index)

		if err1 != nil {
			t.Errorf("%s\n", err1.Error())
		}
		if barrVal != val {
			t.Errorf("index=%d: setVal=%d != getVal=%d", index, val, barrVal)
		}
	}

	checkCompare := func(t *testing.T, barr1 Bitarray, barr2 Bitarray, result bool) {
		t.Helper()
		r := barr1.Compare(&barr2)
		if r != result {
			t.Errorf("resultOfCompare[%v] != givenResult[%v]\n", r, result)
		}
	}

	checkToString := func(t *testing.T, barr Bitarray, val string) {
		t.Helper()
		s := barr.ToString()
		if *s != val {
			t.Errorf("ToString[%s] != %s", *s, val)
		}
	}

	checkAnd := func(t *testing.T, barr1 Bitarray, barr2 Bitarray, result Bitarray) {
		t.Helper()
		res := barr1.And(&barr2)
		if res.Compare(&result) == false {
			t.Errorf("AndRes[%s] != result[%s]", *res.ToString(), *result.ToString())
		}
	}

	checkOr := func(t *testing.T, barr1 Bitarray, barr2 Bitarray, result Bitarray) {
		t.Helper()
		res := barr1.Or(&barr2)
		if res.Compare(&result) == false {
			t.Errorf("OrRes[%s] != result[%s]", *res.ToString(), *result.ToString())
		}
	}

	checkXor := func(t *testing.T, barr1 Bitarray, barr2 Bitarray, result Bitarray) {
		t.Helper()
		res := barr1.Xor(&barr2)
		if res.Compare(&result) == false {
			t.Errorf("XorRes[%s] != result[%s]", *res.ToString(), *result.ToString())
		}
	}

	checkNot := func(t *testing.T, barr1 Bitarray, result Bitarray) {
		t.Helper()
		res := barr1.Not()
		resAnd := barr1.And(res)
		if *resAnd.ToString() != *result.ToString() {
			t.Errorf("NotRes[%s] != result[%s]", *res.ToString(), *result.ToString())
		}
	}

	t.Run("TestInitBySize", func(t *testing.T) {
		barr := Bitarray{}
		barr.InitilizeBySize(1000)
		checkInitBySize(t, barr, 16)
	})

	t.Run("TestInitByeStrValue", func(t *testing.T) {
		barr := Bitarray{}
		barr.InitilizeByStrValue(testString)
		checkInitByStrValue(t, barr, []uint64{0b1111010101010101, 0b1110110100010100101010101000001011110010101011101010111110101011, 0b1110110100010100101010101000001011110010101011101010111110101011,
			0b1111111101010100101000000110100101010011101010101011101001010100, 0b1111111111111111111111110101001010101001111111111111001010101001})

	})

	t.Run("TestGet", func(t *testing.T) {
		barr := Bitarray{}
		barr.InitilizeByStrValue(testString)
		testStringRev := utils.ReverseString(testString)
		checkGet(t, barr, 27, 1)
		checkGet(t, barr, 87, 1)
		checkGet(t, barr, 86, 0)
		for i, v := range testStringRev {
			checkGet(t, barr, uint64(i), uint8(v-'0'))
		}

	})

	t.Run("TestSet", func(t *testing.T) {
		barr := Bitarray{}
		testStringRev := utils.ReverseString(testString)
		for i, v := range testStringRev {
			checkSet(t, barr, uint64(i), uint8(v-'0'))
		}
	})

	t.Run("TestCompare", func(t *testing.T) {
		barr := Bitarray{}
		barr2 := Bitarray{}
		barr3 := Bitarray{}
		barr4 := Bitarray{}

		testStringRev := utils.ReverseString(testString)
		for i, v := range testStringRev {
			barr.Set(uint64(i), uint8(v-'0'))
		}

		barr2.InitilizeByStrValue(testString)

		barr3.InitilizeByStrValue(testString)

		barr4.InitilizeByBitarray(&barr)

		checkCompare(t, barr, barr2, true)
		checkCompare(t, barr, barr3, true)
		checkCompare(t, barr, barr4, true)
		barr2.Set(2, 1)
		checkCompare(t, barr, barr2, false)
	})

	t.Run("TestToString", func(t *testing.T) {
		barr := Bitarray{}
		barr.InitilizeByStrValue(testString)
		checkToString(t, barr, testString)
	})

	t.Run("TestAnd", func(t *testing.T) {
		barr1 := Bitarray{}
		barr2 := Bitarray{}
		result := Bitarray{}
		barr1.InitilizeByStrValue(testString)
		barr2.InitilizeByStrValue(utils.ReverseString(testString))
		resStr := "10010101010001011110110100010100000010101000001011110010101011100010101000001001010001010000000010000010000000000010001010101110100001011010000101110101010001000000000001000001000000001010001010010000010101000111010101001111010000010101000000101000101101111010001010101001"
		result.InitilizeByStrValue(resStr)
		checkAnd(t, barr1, barr2, result)
	})

	t.Run("TestOr", func(t *testing.T) {
		barr1 := Bitarray{}
		barr2 := Bitarray{}
		result := Bitarray{}
		barr1.InitilizeByStrValue(testString)
		barr2.InitilizeByStrValue(utils.ReverseString(testString))
		resStr := "11110101010111111111111110010101111010101111111111111111111111111010111111111111111111011101111010111110100001111111101011111111111111111111111111111111010111111110000101111101011110111011111111111111111101011111111111111111111111110101011110101001111111111111101010101111"
		result.InitilizeByStrValue(resStr)
		checkOr(t, barr1, barr2, result)
	})

	t.Run("TestXor", func(t *testing.T) {
		barr1 := Bitarray{}
		barr2 := Bitarray{}
		result := Bitarray{}
		barr1.InitilizeByStrValue(testString)
		barr2.InitilizeByStrValue(utils.ReverseString(testString))
		resStr := "1100000000110100001001010000001111000000111110100001101010100011000010111110110101110001101111000111100100001111101100001010001011110100101111010001010000110111110000100111100011110110001110101101111101000011000101010110000101111100000011110000001010010000101100000000110"
		result.InitilizeByStrValue(resStr)
		checkXor(t, barr1, barr2, result)
	})

	t.Run("TestNot", func(t *testing.T) {
		barr1 := Bitarray{}
		barr1.InitilizeByStrValue(testString)
		result := Bitarray{}
		result.InitilizeBySize(1)
		checkNot(t, barr1, result)
	})
}
