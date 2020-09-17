Pure-Golang and memory efficient implementation of bitarray. The array size is limited to 2^64(uint64). Bitwise {And, Or, Xor, Not} functions are supported.

### First, initialize a large bitarray string:
```go
var testString string = "11110101010101011110110100010100101010101000001011110010101011101010111110101011111011010001010010101010100000101111001010101110101011111010101111111111010101001010000001101001010100111010101010111010010101001111111111111111111111110101001010101001111111111111001010101001"
b := bitarray.Bitarray{}
b.InitializeByStrValue(testString)
b.Get(0)//returns 1
b.Get(1)//returns 0
b.Get(2)//returns 0
b.Get(3)//returns 1
```
### Compare two bitarray:
```go
b := bitarray.Bitarray{}
b1 := bitarray.Bitarray{}
v := uint64(0b1100010101011001)
b.InitializeByValue(v)
b1.InitializeByValue(v)
b.Compare(b1)//returns true
```

### Bitwise {And, Or, Xor and Not} functions are also be used:
```go
b := bitarray.Bitarray{}
b1 := bitarray.Bitarray{}
b.InitializeByValue(   0b11010011)
b1.InitializeByStrValue("10110010")
*b.And(&b1).ToString()//returns "10010010"
*b.Or(&b1).ToString()//returns "10110010" 
```