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
	storageSize := int64(math.Ceil(float64(size) / 64))
	b.storage = make([]uint64, storageSize)
	return nil
}
