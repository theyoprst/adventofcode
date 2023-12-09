package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
)

func SolvePart1(lines []string) any {
	return 0
}

func SolvePart2(lines []string) any {
	return 0
}

func main() {
	lines1 := aoc.ReadInputLines()
	lines2 := slices.Clone(lines1)
	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	if cmd != "part2" || cmd == "part1" {
		fmt.Println("Part 1:", SolvePart1(lines1))
	}
	if cmd != "part1" || cmd == "part2" {
		fmt.Println("Part 2:", SolvePart2(lines2))
	}
}
