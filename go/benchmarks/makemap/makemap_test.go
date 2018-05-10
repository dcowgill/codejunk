package makemap

import "testing"

var keys3 = []string{
	"main",
	"dessert",
	"drink",
}

var keys9 = []string{
	"main1",
	"dessert1",
	"drink1",
	"main2",
	"dessert2",
	"drink2",
	"main3",
	"dessert3",
	"drink3",
}

var keys18 = []string{
	"main1",
	"dessert1",
	"drink1",
	"main2",
	"dessert2",
	"drink2",
	"main3",
	"dessert3",
	"drink3",
	"main4",
	"dessert4",
	"drink4",
	"main5",
	"dessert5",
	"drink5",
	"main6",
	"dessert6",
	"drink6",
}

var keys50 = []string{
	"anniversary",
	"anybody",
	"architect",
	"back",
	"bank",
	"cart",
	"championship",
	"chemistry",
	"colonial",
	"compete",
	"confrontation",
	"consciousness",
	"coup",
	"eat",
	"electric",
	"emission",
	"feeling",
	"flower",
	"guitar",
	"hall",
	"harmony",
	"incredibly",
	"invite",
	"legally",
	"manner",
	"match",
	"mentally",
	"oversee",
	"participation",
	"pause",
	"placement",
	"plaintiff",
	"powder",
	"radiation",
	"rating",
	"reputation",
	"restaurant",
	"rival",
	"sentiment",
	"shadow",
	"shore",
	"sidewalk",
	"slice",
	"steam",
	"tragedy",
	"universal",
	"user",
	"when",
	"willingness",
	"wow",
}

func BenchmarkMakeMapNoSizeK3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoSize(keys3)
	}
}

func BenchmarkMakeMapWithSizeK3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithSize(keys3)
	}
}

func BenchmarkMakeMapNoSizeK9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoSize(keys9)
	}
}

func BenchmarkMakeMapWithSizeK9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithSize(keys9)
	}
}

func BenchmarkMakeMapNoSizeK18(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoSize(keys18)
	}
}

func BenchmarkMakeMapWithSizeK18(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithSize(keys18)
	}
}

func BenchmarkMakeMapNoSizeK50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoSize(keys50)
	}
}

func BenchmarkMakeMapWithSizeK50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithSize(keys50)
	}
}
