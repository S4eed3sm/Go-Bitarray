package bitarray

import (
	"fmt"
	"math"
	"strconv"
)

var blockLen int = 64
var msb uint64 = 0x8000000000000000

//Bitarray array struct
type Bitarray struct {
	storage []uint64
}

func (b *Bitarray) getStrgIdxInnerIdx(i uint64) (strgIdx uint64, innerIdx uint64, err error) {

	if uint64(len(b.storage)*64) <= i {
		return 0, 0, fmt.Errorf("b.storage have not enough length")
	}
	strgIdx = uint64(len(b.storage)) - i/64 - 1
	innerIdx = 1 << (i % 64)

	return strgIdx, innerIdx, err
}

//InitializeBySize init bitarray by given size
func (b *Bitarray) InitializeBySize(size uint64) error {
	storageSize := uint64(math.Ceil(float64(size) / 64))
	b.storage = make([]uint64, storageSize)
	return nil
}

//InitializeByStrValue take a value like "0110111001"
func (b *Bitarray) InitializeByStrValue(val string) error {

	val = removeLeftZeros(val)

	size := uint64(len(val))
	err := b.InitializeBySize(size)
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

//InitializeByValue take a uint like 0b111010
func (b *Bitarray) InitializeByValue(val uint64) error {

	err := b.InitializeBySize(64)
	b.storage[0] = val
	return err
}

//InitializeByBitarray: clone given bitarray
func (b *Bitarray) InitializeByBitarray(bSrc *Bitarray) error {
	b.storage = make([]uint64, len(bSrc.storage))
	copy(b.storage, bSrc.storage)
	return nil
}

//Get i-th index of bitarray
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

//Set: set given index to 0 or 1
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

//Compare: compare given bitarray with caller bitarray and return true or false
//'0001101' is equal to '01101' is equal to '1101'
func (b *Bitarray) Compare(ba *Bitarray) bool {
	if len(b.storage) != len(ba.storage) {
		if *b.ToString() == *ba.ToString() {
			return true
		}
		return false
	}

	for i, v := range b.storage {
		if v != ba.storage[i] {
			return false
		}
	}
	return true
}

//And: Bitwise And given bitarray with caller bitarray and return bitarray result
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

//Or: Bitwise Or given bitarray with caller bitarray and return bitarray result
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

//Xor: Bitwise Xor given bitarray with caller bitarray and return bitarray result
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

//Not: Bitwise Not caller bitarray and return bitarray result
func (b *Bitarray) Not() (res *Bitarray) {
	res = &Bitarray{
		storage: make([]uint64, len(b.storage)),
	}
	for i := 0; i < len(b.storage); i++ {
		res.storage[i] = ^b.storage[i]
	}
	return res
}

//ToString: return string representation of bitarray
func (b *Bitarray) ToString() *string {
	res := ""
	for _, v := range b.storage {
		tmp := fmt.Sprintf("%064b", v)
		_ = tmp
		res += fmt.Sprintf("%064b", v)
	}
	res = removeLeftZeros(res)
	return &res
}

func (b *Bitarray) ShiftLeft(n int) (res *Bitarray) {
	if n == 0 {
		res = &Bitarray{}
		res.InitializeByBitarray(b)
		return res
	}

	size := len(b.storage)
	sr := 64 - (n % 64)
	sl := n % 64
	if b.storage[0]>>(sr) != 0 {
		size++
	}
	res = &Bitarray{
		storage: make([]uint64, size),
	}
	for i, v := range b.storage {
		res.storage[i+1] = v << sl
		res.storage[i] += v >> sr
	}
	for i := 0; i < n/64; i++ {
		res.storage = append(res.storage, 0)
	}
	return res
}

func (b *Bitarray) ShiftRight(n int) (res *Bitarray) {
	res = &Bitarray{}
	if n == 0 {
		res.InitializeByBitarray(b)
		return res
	}

	size := len(b.storage)
	dr := n / 64

	if dr >= size {
		res.InitializeBySize(1)
		return res
	}

	sr := n % 64
	r := uint64(math.Pow(2, float64(sr))) - 1
	res.storage = make([]uint64, len(b.storage)-dr)
	copy(res.storage, b.storage[:len(b.storage)-dr])

	var shiftCarry uint64 = 0
	for i, v := range res.storage {
		res.storage[i] = res.storage[i] >> sr
		res.storage[i] += shiftCarry
		shiftCarry = (v & r)
		shiftCarry = shiftCarry << (64 - sr)
	}
	for res.storage[0] == 0 {
		res.storage = res.storage[1:]
	}
	return res
}

func (b *Bitarray) Flip(i uint64) error {
	v, err := b.Get(i)
	if err != nil {
		return err
	}
	b.Set(i, 1-v)
	return nil
}
