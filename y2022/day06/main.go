package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
)

func SolveGeneric(lines []string, limit int) []int {
	var res []int
	for _, line := range lines {
		for i := 0; i < len(line)-limit; i++ {
			sub := []byte(line[i : i+limit])
			slices.Sort(sub)
			dup := false
			for j := 0; j < len(sub)-1; j++ {
				if sub[j+1] == sub[j] {
					dup = true
				}
			}
			if !dup {
				res = append(res, i+limit)
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
