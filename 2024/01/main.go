package main

import (
	"context"
	"sort"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(_ context.Context, lines []string) any {
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

func SolvePart2(_ context.Context, lines []string) any {
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

func SolvePart1AI(_ context.Context, lines []string) any {
	leftList := make([]int, 0, len(lines))
	rightList := make([]int, 0, len(lines))
	for _, line := range lines {
		fields := strings.Fields(line)
		leftList = append(leftList, must.Atoi(fields[0]))
		rightList = append(rightList, must.Atoi(fields[1]))
	}
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := range leftList {
		totalDistance += aoc.Abs(leftList[i] - rightList[i])
	}

	return totalDistance
}

func SolvePart2AI(_ context.Context, lines []string) any {
	leftList := make([]int, 0, len(lines))
	rightList := make([]int, 0, len(lines))
	for _, line := range lines {
		fields := strings.Fields(line)
		leftList = append(leftList, must.Atoi(fields[0]))
		rightList = append(rightList, must.Atoi(fields[1]))
	}
	sort.Ints(leftList)
	sort.Ints(rightList)

	rightCount := make(map[int]int)
	for _, num := range rightList {
		rightCount[num]++
	}

	similarityScore := 0
	for _, num := range leftList {
		similarityScore += num * rightCount[num]
	}

	return similarityScore
}

var (
	solvers1 = []aoc.Solver{SolvePart1, SolvePart1AI}
	solvers2 = []aoc.Solver{SolvePart2, SolvePart2AI}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
