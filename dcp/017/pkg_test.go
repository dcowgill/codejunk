package dcp017

import "testing"

func TestParse(t *testing.T) {
	var tests = []struct {
		fs   string
		path string
	}{
		{
			"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext",
			"dir/subdir2/file.ext",
		},
		{
			"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
			"dir/subdir2/subsubdir2/file2.ext",
		},
	}
	for _, tt := range tests {
		t.Run(tt.fs, func(t *testing.T) {
			path := longestAbsPath(parse(tt.fs))
			if path != tt.path {
				t.Fatalf("got %q, want %q", path, tt.path)
			}
		})
	}
}
