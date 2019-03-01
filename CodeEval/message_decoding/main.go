package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// MESSAGE DECODING
// CHALLENGE DESCRIPTION:
//
// Credits: This challenge has appeared in a past ACM competition.
//
// Some message encoding schemes require that an encoded message be sent
// in two parts. The first part, called the header, contains the
// characters of the message. The second part contains a pattern that
// represents the message.You must write a program that can decode
// messages under such a scheme.
//
// The heart of the encoding scheme for your program is a sequence of
// "key" strings of 0's and 1's as follows:
//
// 0,00,01,10,000,001,010,011,100,101,110,0000,0001,...,1011,1110,00000,...
//
// The first key in the sequence is of length 1, the next 3 are of
// length 2, the next 7 of length 3, the next 15 of length 4, etc. If
// two adjacent keys have the same length, the second can be obtained
// from the first by adding 1 (base 2). Notice that there are no keys in
// the sequence that consist only of 1's.
//
// The keys are mapped to the characters in the header in order. That
// is, the first key (0) is mapped to the first character in the header,
// the second key (00) to the second character in the header, the kth
// key is mapped to the kth character in the header. For example,
// suppose the header is:
//
//		AB#TANCnrtXc
//
// Then 0 is mapped to A, 00 to B, 01 to #, 10 to T, 000 to A, ..., 110
// to X, and 0000 to c.
//
// The encoded message contains only 0's and 1's and possibly carriage
// returns, which are to be ignored. The message is divided into
// segments. The first 3 digits of a segment give the binary
// representation of the length of the keys in the segment. For example,
// if the first 3 digits are 010, then the remainder of the segment
// consists of keys of length 2 (00, 01, or 10). The end of the segment
// is a string of 1's which is the same length as the length of the keys
// in the segment. So a segment of keys of length 2 is terminated by 11.
// The entire encoded message is terminated by 000 (which would signify
// a segment in which the keys have length 0). The message is decoded by
// translating the keys in the segments one-at-a-time into the header
// characters to which they have been mapped.

type message struct {
	header string
	keys   []string
}

func parseTestCase(s string) message {
	if len(s) < 3 || s[len(s)-3:] != "000" {
		log.Fatalf("bad message: %#v", s)
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '0' && s[i] != '1' {
			return message{
				header: s[:i+1],
				keys:   parseSegments(s[i+1:]),
			}
		}
	}
	return message{
		header: s,
		keys:   []string{},
	}
}

func mustParseBinaryInt(s string) int {
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(n)
}

func isSegmentTerminator(s string) bool {
	for _, c := range s {
		if c != '1' {
			return false
		}
	}
	return true
}

func parseSegments(msg string) []string {
	keys := make([]string, 0)
	for i := 0; i < len(msg)-3; {
		klen_str := msg[i : i+3]
		klen := mustParseBinaryInt(klen_str)
		if klen_str == "000" || klen < 1 {
			log.Fatalf("bad key length %#v (%d) at offset %d", klen_str, klen, i)
		}
		for i += 3; i < len(msg)-klen; i += klen {
			k := msg[i : i+klen]
			if isSegmentTerminator(k) {
				break
			}
			keys = append(keys, k)
		}
		i += klen
	}
	return keys
}

type keyMap map[string]rune

func buildKeyMap(header string) keyMap {
	// this simply computes 2**len - 1
	calcMax := func(len int) int {
		n := 1
		for i := 0; i < len; i++ {
			n *= 2
		}
		return n - 1
	}
	// convert integer 'k' to base-2 string, padded to 'len' chars
	makeKey := func(k, len int) string {
		key := make([]byte, len)
		for i := len - 1; i >= 0; i-- {
			if k%2 == 0 {
				key[i] = '0'
			} else {
				key[i] = '1'
			}
			k /= 2
		}
		return string(key)
	}
	var (
		key    = 0
		keylen = 1
		max    = calcMax(keylen)
		keymap = make(keyMap, len(header))
	)
	for _, r := range header {
		keymap[makeKey(key, keylen)] = r
		key++
		if key >= max {
			keylen++
			max = calcMax(keylen)
			key = 0
		}
	}
	return keymap
}

func decodeMessage(msg message, km keyMap) string {
	runes := make([]rune, len(msg.keys))
	for i, k := range msg.keys {
		r, ok := km[k]
		if !ok {
			log.Fatalf("bad key %#v", k)
		}
		runes[i] = r
	}
	return string(runes)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		testcase := scanner.Text()
		message := parseTestCase(testcase)
		keymap := buildKeyMap(message.header)
		fmt.Println(decodeMessage(message, keymap))
	}
}
