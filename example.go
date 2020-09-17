package main

import (
	"fmt"

	bitarray "gitlab.com/S4eed3sm/bit-array-in-golang/bitarray"
)

var testString string = "11110101010101011110110100010100101010101000001011110010101011101010111110101011111011010001010010101010100000101111001010101110101011111010101111111111010101001010000001101001010100111010101010111010010101001111111111111111111111110101001010101001111111111111001010101001"

func main() {

	b := bitarray.Bitarray{}
	b1 := bitarray.Bitarray{}
	b2 := bitarray.Bitarray{}
	b3 := bitarray.Bitarray{}
	b4 := bitarray.Bitarray{}
	bRev := bitarray.Bitarray{}
	result := bitarray.Bitarray{}

	b.InitializeByStrValue(testString)

	revStr := bitarray.ReverseString(testString)
	fmt.Println("revStr", revStr)
	for i, v := range revStr {
		b2.Set(uint64(i), uint8(v-'0'))
	}

	fmt.Printf("Compare: %v\n", b2.Compare(&b))

	b4.InitializeByBitarray(&b3)

	fmt.Printf("Compare: %v\n", b4.Compare(&b3))

	s := b.ToString()
	fmt.Println("testOfToString:", *s == testString)

	bRev.InitializeByStrValue(revStr)

	res := b.And(&bRev)
	resStr := "10010101010001011110110100010100000010101000001011110010101011100010101000001001010001010000000010000010000000000010001010101110100001011010000101110101010001000000000001000001000000001010001010010000010101000111010101001111010000010101000000101000101101111010001010101001"

	result.InitializeByStrValue(resStr)
	fmt.Println("Compare:", result.Compare(res))

	b.InitializeByValue(0b11010011)
	b1.InitializeByStrValue("10110010")
	fmt.Println("And:", *b.And(&b1).ToString())
	fmt.Println("Or:", *b.Or(&b1).ToString())
	fmt.Println("Xor:", *b.Xor(&b1).ToString())
	fmt.Println("Not:", *b.Not().ToString())

	b.InitializeByStrValue(testString)
	fmt.Println("ShiftLeft:", *b.ShiftLeft(127).ToString())
}
