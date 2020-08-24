package bitarray

import (
	"fmt"
	"math"
	"strconv"
)

var blockLen int = 64

//Bitarray array struct
type Bitarray struct {
	storage []uint64
}

//InitilizeBySize init bitarray with given size
func (b *Bitarray) InitilizeBySize(size int) error {
	storageSize := int64(math.Ceil(float64(size) / 64))
	b.storage = make([]uint64, storageSize)
	return nil
}

//InitilizeByStrValue take a value like "0110111001"
func (b *Bitarray) InitilizeByStrValue(val string) error {

	for val[0] == '0' {
		val = val[1:]
	}

	size := len(val)
	err := b.InitilizeBySize(size)
	var tmp uint64

	if err == nil {
		i := len(b.storage) - 1
		for len(val) >= blockLen && i >= 0 {
			x := val[len(val)-blockLen:]
			tmp, err = strconv.ParseUint(val[len(val)-blockLen:], 2, 64)
			fmt.Println(x, tmp)
			b.storage[i] = tmp
			val = val[:len(val)-blockLen]
			i--
		}
		if len(val) > 0 && i == 0 {
			tmp, err = strconv.ParseUint(val, 2, 64)
			b.storage[i] = tmp
		} else {
			err = fmt.Errorf("len(val)=%d, i=%d", len(val), i)
		}
	}

	return err
}

//InitilizeByValue take a uint like 0b111010
func (b *Bitarray) InitilizeByValue(val uint64) error {

	err := b.InitilizeBySize(64)
	b.storage[0] = val
	return err
}
