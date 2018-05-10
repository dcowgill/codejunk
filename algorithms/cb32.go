package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

//
// Encode
//

const cb32chars = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

func encode(n uint64) string {
	var bytes [13]byte
	chars := []byte(cb32chars)
	i := 0
	for n > 0 {
		// fmt.Printf("%b & 0x1F = %v (%c)\n", n, n&0x1F, chars[n&0x1F])
		bytes[i] = chars[n&0x1F]
		i++
		n >>= 5
	}
	fmt.Println(string(bytes[:i]))
	return string(bytes[:i])
}

//
// Decode
//

var decodeTable [256]int8

type CorruptInputError int64

func (e CorruptInputError) Error() string {
	return "illegal base32 data at input byte " + strconv.FormatInt(int64(e), 10)
}

func decode(s string) (uint64, error) {
	var r uint64
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}
		n := decodeTable[s[i]]
		if n < 0 {
			return 0, CorruptInputError(i)
		}
		r <<= 5
		r |= uint64(n)
	}
	return r, nil
}

var tests = []uint64{
	9223372036854775807,  // 2^63-1
	18446744073709551615, // 2^64-1
}

func runTests() {
	for i := -100; i <= 100; i++ {
		test(uint64(i))
	}
	for _, n := range tests {
		test(n)
	}
	for i := 0; i < 100; i++ {
		test(uint64(rand.Int()))
	}
}

func test(n uint64) {
	e := encode(n)
	r, err := decode(e)
	if err != nil {
		log.Fatalf("decode(encode(%d)) => %s\n", n, err)
	} else if r != n {
		log.Fatalf("decode(encode(%d)) failed; got %v\n", n, r)
	}
}

func main() {
	// var n uint64 = 9223372036854775807
	// s := encode(n)
	// fmt.Println(s)
	// r, err := decode(s)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Decoded %v\n", r)
	// if n != r {
	// 	log.Fatal("%v != %v\n", n, r)
	// }

	runTests()
}

func init() {
	for i := 0; i < len(decodeTable); i++ {
		decodeTable[i] = -1
	}
	for i, b := range "0123456789" {
		decodeTable[b] = int8(i)
	}
	for i, b := range "ABCDEFGHJKMNPQRSTVWXYZ" {
		decodeTable[b] = int8(i + 10)
	}
	for i, b := range "abcdefghjkmnpqrstvwxyz" {
		decodeTable[b] = int8(i + 10)
	}
	for _, b := range "Oo" {
		decodeTable[b] = decodeTable['0']
	}
	for _, b := range "IiLl" {
		decodeTable[b] = decodeTable['1']
	}
	// for i, n := range decodeTable {
	// 	fmt.Printf("%c => %d\n", i, n)
	// }
}
