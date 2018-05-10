package stringeq

import "testing"

var strings = []string{"", "1", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec faucibus sodales felis, eu efficitur ex faucibus quis. Phasellus quis dui elit. Etiam eget turpis in tortor mattis vestibulum sed ut sapien. Nunc consectetur pellentesque sapien vitae mollis. Integer vitae porttitor nisi. Quisque commodo velit nec tempus vestibulum. Morbi ac justo quis diam hendrerit mollis. Vestibulum volutpat porttitor turpis eu tincidunt. Morbi vel finibus urna. Vestibulum tempus magna eu sapien pellentesque condimentum. Sed vestibulum pulvinar nunc vel porttitor. Nam ac neque at lorem malesuada eleifend a non lacus. Aliquam erat volutpat.", "In hac habitasse platea dictumst."}

func BenchmarkIsEmpty1(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		for _, s := range strings {
			if len(s) == 0 {
				n++
			}
		}
	}
}

func BenchmarkIsEmpty2(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		for _, s := range strings {
			if s == "" {
				n++
			}
		}
	}
}

func BenchmarkIsNotEmpty1(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		for _, s := range strings {
			if len(s) > 0 {
				n++
			}
		}
	}
}

func BenchmarkIsNotEmpty2(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		for _, s := range strings {
			if s != "" {
				n++
			}
		}
	}
}
