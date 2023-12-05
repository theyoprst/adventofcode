package main

import (
	"fmt"

	"github.com/theyoprst/adventofcode/aoc"
)

func main() {
	var ans1, ans2 int
	lines := aoc.ReadInputLines()
	for i, line := range lines {
		_, _ = i, line
	}

	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}
