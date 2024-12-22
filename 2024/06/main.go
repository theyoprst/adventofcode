package main

import (
	"context"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

const (
	guardCh    = '^'
	obstacleCh = '#'
	freeCh     = '.'
)

var dirs = []fld.Pos{fld.Up, fld.Right, fld.Down, fld.Left}

func SolvePart1(_ context.Context, lines []string) any {
	field := fld.NewByteField(lines)
	guardPos := field.FindFirst(guardCh)
	dirIdx := 0
	visited := containers.NewSet(guardPos)
	for {
		npos := guardPos.Add(dirs[dirIdx])
		if !field.Inside(npos) {
			break
		}
		if field.Get(npos) == obstacleCh {
			dirIdx = (dirIdx + 1) % len(dirs) // Turn right.
		} else {
			guardPos = npos
			visited.Add(guardPos)
		}
	}
	return len(visited)
}

func SolvePart2(_ context.Context, lines []string) any {
	// Brute force. Put obstacles on the path and check if the guard will loop.
	// Optimizations:
	// 1. Put obstacles only on the path from part1, not on the whole field (speedup x4)
	// 2. Remember visited states only on turns, not on every step (speedup x3)
	// 3. Start loop detection from the new obstacle, not from the initial guard position (speedup x3)
	// 4. Jump to the next obstable using precomputed obstacles positions (speedup x3)
	// 5. Track only visited states on turns from UP directions (speedup x2)
	// Total: 0.02s for the input (from 5s without optimizations).
	field := fld.NewByteField(lines)
	guardPos := field.FindFirst(guardCh)

	scanRowObstacles := func(row int) []int {
		var cols []int
		for col := range field.Cols() {
			if field.Get(fld.NewPos(row, col)) == obstacleCh {
				cols = append(cols, col)
			}
		}
		return cols
	}

	scanColObstacles := func(col int) []int {
		var rows []int
		for row := range field.Rows() {
			if field.Get(fld.NewPos(row, col)) == obstacleCh {
				rows = append(rows, row)
			}
		}
		return rows
	}

	rowObstacles := make([][]int, field.Rows())
	for row := range field.Rows() {
		rowObstacles[row] = scanRowObstacles(row)
	}

	colObstacles := make([][]int, field.Cols())
	for col := range field.Cols() {
		colObstacles[col] = scanColObstacles(col)
	}

	ans := 0
	dirIdx := 0
	visited := containers.NewSet(guardPos)
	for {
		npos := guardPos.Add(dirs[dirIdx])
		if !field.Inside(npos) {
			break
		}
		if field.Get(npos) == obstacleCh {
			dirIdx = (dirIdx + 1) % len(dirs) // Turn right.
		} else {
			if !visited.Has(npos) {
				field.Set(npos, obstacleCh)
				rowObstacles[npos.Row] = scanRowObstacles(npos.Row)
				colObstacles[npos.Col] = scanColObstacles(npos.Col)
				if isLooped(rowObstacles, colObstacles, guardPos, dirIdx) {
					ans++
				}
				field.Set(npos, freeCh)
				rowObstacles[npos.Row] = scanRowObstacles(npos.Row)
				colObstacles[npos.Col] = scanColObstacles(npos.Col)
				visited.Add(npos)
			}
			guardPos = npos
		}
	}
	return ans
}

func isLooped(rowObstacles, colObstacles [][]int, guardPos fld.Pos, dirIdx int) bool {
	type State struct {
		pos    fld.Pos
		dirIdx int
	}
	state := State{pos: guardPos, dirIdx: dirIdx}
	// Only track visited states on turns from UP directions, to speed up overall execution.
	visitedUp := containers.NewSet[fld.Pos]()
	for state.dirIdx != 0 || !visitedUp.Has(state.pos) {
		switch state.dirIdx {
		case 0: // UP
			visitedUp.Add(state.pos)
			// Find obstable in this col above the guard (row value less than guard's row).
			state.pos.Row = lowerBound(colObstacles[state.pos.Col], state.pos.Row, -100) + 1
		case 1: // RIGHT
			// Find obstable in this row to the right of the guard (col value greater than guard's col).
			state.pos.Col = upperBound(rowObstacles[state.pos.Row], state.pos.Col, -100) - 1
		case 2: // DOWN
			// Find obstable in this col below the guard (row value greater than guard's row).
			state.pos.Row = upperBound(colObstacles[state.pos.Col], state.pos.Row, -100) - 1
		case 3: // LEFT
			// Find obstable in this row to the left of the guard (col value less than guard's col).
			state.pos.Col = lowerBound(rowObstacles[state.pos.Row], state.pos.Col, -100) + 1
		default:
			panic("invalid direction")
		}
		if state.pos.Row < 0 || state.pos.Col < 0 {
			return false
		}
		// Turn right
		state.dirIdx = (state.dirIdx + 1) % len(dirs)
	}
	return true
}

// lowerBound returns the maximum value in `arr` which is less than `value`.
// `arr` is not empty and sorted in ascending order.
func lowerBound(arr []int, value int, defaultValue int) int {
	// Binary search doesn't make sense here because the array is small.
	ans := defaultValue
	for _, x := range arr {
		if x >= value {
			break
		}
		ans = x
	}
	return ans
}

// upperBound returns the minimun value in `arr` which is greater than `value`.
// `arr` is not empty and sorted in ascending order.
func upperBound(arr []int, value int, defaultValue int) int {
	// Binary search doesn't make sense here because the array is small.
	for _, x := range arr {
		if x > value {
			return x
		}
	}
	return defaultValue
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
