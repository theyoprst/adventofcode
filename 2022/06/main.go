package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
)

func SolveGeneric(lines []string, limit int) []int {
	var res []int
	for _, line := range lines {
		for i := limit; i < len(line); i++ {
			sub := []byte(line[i-limit : i])
			if len(containers.NewSet[byte](sub...)) == len(sub) {
				res = append(res, i)
				break
			}
		}
	}
	return res
}

func SolvePart1(lines []string) []int {
	return SolveGeneric(lines, 4)
}

func SolvePart2(lines []string) []int {
	return SolveGeneric(lines, 14)
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
