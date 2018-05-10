// https://en.wikipedia.org/wiki/Rendezvous_hashing
package main

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"log"
	"sort"
)

func urandom(n int) []byte {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return b
}

func sha1sum(b []byte) string {
	s := sha1.Sum(b)
	return hex.EncodeToString(s[:])
}

func hash(a ...string) string {
	var s string
	for _, t := range a {
		s += t
	}
	return sha1sum([]byte(s))
}

func randstr() string { return sha1sum(urandom(20)) }

func main() {
	var nservers, nkeys int
	flag.IntVar(&nservers, "servers", 5, "number of servers")
	flag.IntVar(&nkeys, "keys", 100, "number of keys")
	flag.Parse()
	log.Printf("creating %d keys", nkeys)
	keys := make([]string, nkeys)
	for i := range keys {
		keys[i] = randstr()
	}
	log.Printf("creating %d servers", nservers)
	servers := make([]string, nservers)
	for i := range servers {
		servers[i] = randstr()
	}
	log.Print("hashing keys")
	table1 := make(map[string]string)
	for _, key := range keys {
		table1[key] = assign(key, servers)
	}
	for i := 0; i < 3; i++ {
		log.Print("adding 1 server")
		servers = append(servers, randstr())
		log.Print("re-hashing all keys")
		table2 := make(map[string]string)
		for _, key := range keys {
			table2[key] = assign(key, servers)
		}
		log.Print("comparing hash tables")
		nmoved := 0
		for key, server := range table1 {
			if server != table2[key] {
				nmoved++
			}
		}
		log.Printf("number of keys moved: %d -- %.1f%%", nmoved,
			100*float64(nmoved)/float64(len(keys)))
		table1 = table2
	}
}

func assign(key string, servers []string) string {
	type hashed struct {
		hash string
		sid  int
	}
	hashes := make([]hashed, len(servers))
	for i, server := range servers {
		hashes[i] = hashed{hash(server, key), i}
	}
	sort.Slice(hashes, func(i, j int) bool {
		return hashes[i].hash < hashes[j].hash
	})
	return servers[hashes[0].sid]
}
