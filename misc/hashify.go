// hashify renames every file specified on the command line to its
// hex-encoded SHA-1 hash value, keeping its original extension.
//
// For example, a file named "hello.txt" might be renamed to
// "e644c0c418ce2b1560efe7c2bb46af2f858a4585.txt".
//
package main

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func main() {
	// Parse command-line arguments.
	printOnly := flag.Bool("printonly", false, "dry run")
	flag.Parse()

	// For each argument...
	for _, oldName := range flag.Args() {
		// Attempt to open the file for reading.
		fp, err := os.Open(oldName)
		if err != nil {
			log.Fatalf("error opening %q: %s", oldName, err)
		}

		// Compute the SHA-1 hash value of the file's contents.
		h := sha1.New()
		if _, err := io.Copy(h, fp); err != nil {
			log.Fatalf("error reading from %q: %s", oldName, err)
		}
		hashVal := hex.EncodeToString(h.Sum(nil))

		// New name: <sha1>.<ext>, using the original extension.
		newName := hashVal + path.Ext(oldName)

		// The names differ, print the transform, then rename.
		if oldName != newName {
			fmt.Printf("%s => %s\n", oldName, newName)
			if !*printOnly {
				if err := os.Rename(oldName, newName); err != nil {
					log.Fatalf("error renaming %q: %s", oldName, err)
				}
			}
		}
	}
}
