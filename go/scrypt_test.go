package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/crypto/scrypt"
)

// Default encryption settings.
const (
	saltLen      = 64
	scryptN      = 32768
	scryptR      = 16
	scryptP      = 2
	scryptKeyLen = 64
)

type ScryptedPassword struct {
	SaltLen int
	N       int
	R       int
	P       int
	KeyLen  int
	Salt    []byte
	Key     []byte
}

// Return a slice of n random bytes.
func makeSalt(n int) []byte {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return b
}

// Encrypt 'password'.
func hashPassword(password string) *ScryptedPassword {
	salt := makeSalt(saltLen)
	key, err := scrypt.Key([]byte(password), salt, scryptN, scryptR, scryptP, scryptKeyLen)
	if err != nil {
		panic(err)
	}
	return &ScryptedPassword{
		SaltLen: saltLen,
		N:       scryptN,
		R:       scryptR,
		P:       scryptP,
		KeyLen:  scryptKeyLen,
		Salt:    salt,
		Key:     key,
	}
}

// Test if 'password' is the string originally used to create 's', via
// the hashPassword function.
func (s ScryptedPassword) Verify(password string) bool {
	key, err := scrypt.Key([]byte(password), s.Salt, s.N, s.R, s.P, s.KeyLen)
	if err != nil {
		panic(err)
	}
	if len(s.Key) != len(key) {
		return false
	}
	for i, b := range s.Key {
		if b != key[i] {
			return false
		}
	}
	return true
}

func main() {
	s1 := hashPassword("owueropwiuqrepoqwujflksjhflksajflkasjhdf")

	bytes, err := json.Marshal(s1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
	fmt.Println(len(string(bytes)))

	var s2 ScryptedPassword
	if err := json.Unmarshal(bytes, &s2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(s1.Verify("owueropwiuqrepoqwujflksjhflksajflkasjhdf"))
	fmt.Println(s2.Verify("owueropwiuqrepoqwujflksjhflksajflkasjhdf"))
}
