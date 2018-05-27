/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Suppose we represent our file system by a string in the following manner:

The string "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext" represents:

dir
    subdir1
    subdir2
        file.ext

The directory dir contains an empty sub-directory subdir1 and a sub-directory
subdir2 containing a file file.ext.

The string
"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
represents:

dir
    subdir1
        file1.ext
        subsubdir1
    subdir2
        subsubdir2
            file2.ext

The directory dir contains two sub-directories subdir1 and subdir2. subdir1
contains a file file1.ext and an empty second-level sub-directory subsubdir1.
subdir2 contains a second-level sub-directory subsubdir2 containing a file
file2.ext.

We are interested in finding the longest (number of characters) absolute path to
a file within our file system. For example, in the second example above, the
longest absolute path is "dir/subdir2/subsubdir2/file2.ext", and its length is
32 (not including the double quotes).

Given a string representing the file system in the above format, return the
length of the longest absolute path to a file in the abstracted file system. If
there is no file in the system, return 0.

Note:

The name of a file contains at least a period and an extension.

The name of a directory or sub-directory will not contain a period.
*/
package dcp017

import (
	"strings"
)

// Represents a file or directory.
type file struct {
	name  string
	files []*file
}

// Files must have an extension; directories never do.
func (f *file) isDir() bool {
	return !strings.ContainsRune(f.name, '.')
}

// Parses the encoded filesystem string and returns a set of files.
func parse(filesystem string) []*file {
	files, _ := parseRec(filesystem, 0, 0)
	return files
}

// Recursive implementation of parse().
func parseRec(fs string, pos int, depth int) ([]*file, int) {
	var files []*file
	for pos < len(fs) {
		tabs := leadingTabs(fs, pos)
		if tabs != depth {
			break
		}
		pos += tabs
		i := pos
		for ; i < len(fs); i++ {
			if fs[i] == '\n' {
				break
			}
		}
		f := file{name: fs[pos:i]}
		f.files, pos = parseRec(fs, i+1, depth+1)
		files = append(files, &f)
	}
	return files, pos
}

// Reports the number of consecutive tab runes starting at pos in s.
func leadingTabs(s string, pos int) int {
	i := pos
	for ; i < len(s); i++ {
		if s[i] != '\t' {
			break
		}
	}
	return i - pos
}

// Reports the longest absolute path in the filesystem.
func longestAbsPath(fs []*file) string {
	var max string
	for _, f := range fs {
		if f.isDir() {
			s := f.name + "/" + longestAbsPath(f.files)
			if len(s) > len(max) {
				max = s
			}
		} else if len(f.name) > len(max) {
			max = f.name
		}
	}
	return max
}
