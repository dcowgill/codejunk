package main

func main() {
	part := flag.Int("part", 1, "which part")
	flag.Parse()
	switch *part {
	case 1:
		part1()
	case 2:
		part2()
	}
}

func part1() {
	fmt.Println("part1")
}

func part2() {
	fmt.Println("part2")
}
