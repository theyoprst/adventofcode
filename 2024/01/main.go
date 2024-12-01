package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
)

func SolvePart1(lines []string) any {
	var x, y []int
	for _, line := range lines {
		ii := aoc.Ints(line)
		x = append(x, ii[0])
		y = append(y, ii[1])
	}
	sort.Ints(x)
	sort.Ints(y)
	diff := 0
	for i := range x {
		diff += aoc.Abs(x[i] - y[i])
	}
	return diff
}

func SolvePart2(lines []string) any {
	var x []int
	counter := map[int]int{}
	for _, line := range lines {
		ii := aoc.Ints(line)
		x = append(x, ii[0])
		counter[ii[1]]++
	}
	score := 0
	for i := range x {
		score += counter[x[i]] * x[i]
	}
	return score
}

func SolvePart1Qwen32B(lines []string) any {
	left := make([]int, 0)
	right := make([]int, 0)
	for _, line := range lines {
		parts := strings.Fields(line)
		left = append(left, atoi(parts[0]))
		right = append(right, atoi(parts[1]))
	}
	sort.Ints(left)
	sort.Ints(right)
	totalDistance := 0
	for i := range left {
		totalDistance += abs(left[i] - right[i])
	}
	return totalDistance
}

func SolvePart2Qwen32B(lines []string) any {
	left := make([]int, 0)
	rightCount := make(map[int]int)
	for _, line := range lines {
		parts := strings.Fields(line)
		left = append(left, atoi(parts[0]))
		rightCount[atoi(parts[1])]++
	}
	similarityScore := 0
	for _, num := range left {
		similarityScore += num * rightCount[num]
	}
	return similarityScore
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var (
	solvers1 = []aoc.Solver{SolvePart1, SolvePart1Qwen32B}
	solvers2 = []aoc.Solver{SolvePart2, SolvePart2Qwen32B}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
