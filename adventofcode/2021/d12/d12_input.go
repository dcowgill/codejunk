package d12

var (
	exampleInput1 = [][]string{
		{"start", "A"},
		{"start", "b"},
		{"A", "c"},
		{"A", "b"},
		{"b", "d"},
		{"A", "end"},
		{"b", "end"},
	}

	exampleInput2 = [][]string{
		{"dc", "end"},
		{"HN", "start"},
		{"start", "kj"},
		{"dc", "start"},
		{"dc", "HN"},
		{"LN", "dc"},
		{"HN", "end"},
		{"kj", "sa"},
		{"kj", "HN"},
		{"kj", "dc"},
	}

	exampleInput3 = [][]string{
		{"fs", "end"},
		{"he", "DX"},
		{"fs", "he"},
		{"start", "DX"},
		{"pj", "DX"},
		{"end", "zg"},
		{"zg", "sl"},
		{"zg", "pj"},
		{"pj", "he"},
		{"RW", "he"},
		{"fs", "DX"},
		{"pj", "RW"},
		{"zg", "RW"},
		{"start", "pj"},
		{"he", "WI"},
		{"zg", "he"},
		{"pj", "fs"},
		{"start", "RW"},
	}

	realInput = [][]string{
		{"cz", "end"},
		{"cz", "WR"},
		{"TD", "end"},
		{"TD", "cz"},
		{"start", "UM"},
		{"end", "pz"},
		{"kb", "UM"},
		{"mj", "UM"},
		{"cz", "kb"},
		{"WR", "start"},
		{"WR", "pz"},
		{"kb", "WR"},
		{"TD", "kb"},
		{"mj", "kb"},
		{"TD", "pz"},
		{"UM", "pz"},
		{"kb", "start"},
		{"pz", "mj"},
		{"WX", "cz"},
		{"sp", "WR"},
		{"mj", "WR"},
	}
)
