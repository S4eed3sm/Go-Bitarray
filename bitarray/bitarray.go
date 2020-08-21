package bitarray

import (
	"math"
)

//Bitarray array struct
type Bitarray struct {
	storage []uint64
}

//InitilizeBySize init bitarray with given size
func (b *Bitarray) InitilizeBySize(size int) error {
	powOf2 := math.Log2(float64(size))
	listCount := int64(math.Ceil(powOf2 / 64))
	b.storage = make([]uint64, listCount)

	return nil
}
