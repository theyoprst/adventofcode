package main

import (
	"fmt"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

// TODO: implement input params in tests.yaml
const gridSize = 71
const firstBytes = 1024

func SolvePart1(lines []string) any {
	bytesPositions := parseBytePositions(lines)
	corrupted := containers.NewSet(bytesPositions[:firstBytes]...)
	return startToFinishDist(corrupted, gridSize)
}

func SolvePart2(lines []string) any {
	bytePositions := parseBytePositions(lines)

	left := firstBytes - 1          // invariant: left index doesn't block
	right := len(bytePositions) - 1 // invariant: right index blocks
	for left+1 < right {
		mid := (left + right) / 2
		corrupted := containers.NewSet(bytePositions[:mid+1]...)
		if startToFinishDist(corrupted, gridSize) == -1 {
			right = mid
		} else {
			left = mid
		}
	}
	pos := bytePositions[right]
	return fmt.Sprintf("%d,%d", pos.Col, pos.Row)
}

func parseBytePositions(lines []string) []fld.Pos {
	positions := make([]fld.Pos, len(lines))
	for i, line := range lines {
		a := aoc.Ints(line)
		positions[i] = fld.Pos{Col: a[0], Row: a[1]}
	}
	return positions
}

func startToFinishDist(corrupted containers.Set[fld.Pos], gridSize int) int {
	start := fld.Pos{Col: 0, Row: 0}
	finish := fld.Pos{Col: gridSize - 1, Row: gridSize - 1}

	var queue []fld.Pos
	queue = append(queue, start)
	dist := make(map[fld.Pos]int)
	dist[start] = 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, dir := range fld.DirsSimple {
			next := cur.Add(dir)
			if next.Col < 0 || next.Col >= gridSize || next.Row < 0 || next.Row >= gridSize {
				continue
			}
			if corrupted.Has(next) {
				continue
			}
			if _, ok := dist[next]; ok {
				continue
			}
			dist[next] = dist[cur] + 1
			queue = append(queue, next)
			if next == finish {
				return dist[next]
			}
		}
	}

	return -1
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
