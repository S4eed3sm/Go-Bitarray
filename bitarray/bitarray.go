package bitarray

import (
	"fmt"
	"math"
	"strconv"

	"gitlab.com/S4eed3sm/bit-array-in-golang/utils"
)

var blockLen int = 64

//Bitarray array struct
type Bitarray struct {
	storage []uint64
}

func (b *Bitarray) getStrgIdxInnerIdx(i uint64) (strgIdx uint64, innerIdx uint64, err error) {

	if uint64(len(b.storage)*64) <= i {
		return 0, 0, fmt.Errorf("b.storage have not enough lenght")
	}
	strgIdx = uint64(len(b.storage)) - i/64 - 1
	innerIdx = 1 << (i % 64)

	return strgIdx, innerIdx, err
}

//InitilizeBySize init bitarray with given size
func (b *Bitarray) InitilizeBySize(size uint64) error {
	storageSize := uint64(math.Ceil(float64(size) / 64))
	b.storage = make([]uint64, storageSize)
	return nil
}

//InitilizeByStrValue take a value like "0110111001"
func (b *Bitarray) InitilizeByStrValue(val string) error {

	val = utils.RemoveRightZeros(val)

	size := uint64(len(val))
	err := b.InitilizeBySize(size)
	var tmp uint64

	if err == nil {
		i := len(b.storage) - 1
		for len(val) >= blockLen && i >= 0 {
			tmp, err = strconv.ParseUint(val[len(val)-blockLen:], 2, 64)
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

//
func (b *Bitarray) InitilizeByBitarray(ba *Bitarray) error {
	b.storage = ba.storage
	return nil
}

//Get i'th index of bitarray
func (b *Bitarray) Get(i uint64) (uint8, error) {
	strgIdx, innerIdx, err := b.getStrgIdxInnerIdx(i)
	if err != nil {
		return 0, err
	}
	val := b.storage[strgIdx] & innerIdx
	if val == 0 {
		return 0, nil
	}
	return 1, nil
}

func (b *Bitarray) resize(i uint64) error {
	appendSize := uint64(math.Ceil(float64(i-uint64(len(b.storage)*64-1)) / 64))
	b.storage = append(make([]uint64, appendSize), b.storage...)
	return nil
}

func (b *Bitarray) Set(i uint64, v uint8) error {
	strgIdx, innerIdx, err := b.getStrgIdxInnerIdx(i)
	if err != nil {
		err = b.resize(i)
		if err == nil {
			return b.Set(i, v)
		}
		return err
	}

	if v == 0 {
		b.storage[strgIdx] &= ^(innerIdx)
	} else if v == 1 {
		b.storage[strgIdx] |= innerIdx
	} else {
		return fmt.Errorf("v must be 0 or 1, but you pass[%d]", v)
	}
	return nil
}

func (b *Bitarray) Compare(ba *Bitarray) bool {
	if len(b.storage) != len(ba.storage) {
		return false
	}

	for i, v := range b.storage {
		if v != ba.storage[i] {
			return false
		}
	}
	return true
}

func (b *Bitarray) And(ba *Bitarray) (res *Bitarray) {
	tmp := &Bitarray{
		storage: []uint64{},
	}
	if len(b.storage) <= len(ba.storage) {
		res = b
		tmp = ba
	} else {
		res = ba
		tmp = b
	}
	for i := len(res.storage) - 1; i >= 0; i-- {
		res.storage[i] &= tmp.storage[i]
	}
	return res
}

func (b *Bitarray) Or(ba *Bitarray) (res *Bitarray) {
	tmp := &Bitarray{
		storage: []uint64{},
	}
	if len(b.storage) <= len(ba.storage) {
		res = ba
		tmp = b
	} else {
		res = b
		tmp = ba
	}

	for i := len(tmp.storage) - 1; i >= 0; i-- {
		res.storage[i] |= tmp.storage[i]
	}
	return res
}

func (b *Bitarray) Xor(ba *Bitarray) (res *Bitarray) {
	tmp := &Bitarray{
		storage: []uint64{},
	}
	if len(b.storage) <= len(ba.storage) {
		res = ba
		tmp = b
	} else {
		res = b
		tmp = ba
	}
	for i := len(tmp.storage) - 1; i >= 0; i-- {
		res.storage[i] ^= tmp.storage[i]
	}
	return res
}

func (b *Bitarray) Not() (res *Bitarray) {
	res = &Bitarray{
		storage: make([]uint64, len(b.storage)),
	}
	for i := 0; i < len(b.storage); i++ {
		res.storage[i] = ^b.storage[i]
	}
	return res
}

func (b *Bitarray) ToString() *string {
	res := ""
	for _, v := range b.storage {
		res += fmt.Sprintf("%b", v)
	}
	res = utils.RemoveRightZeros(res)
	return &res
}
