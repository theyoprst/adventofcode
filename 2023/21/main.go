package main

import (
	"bytes"
	"context"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/fld"
	"github.com/theyoprst/adventofcode/aoc/graphs"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(_ context.Context, lines []string) any {
	field := fld.NewByteField(lines)
	start := field.FindFirst('S')
	return CountReachable(field, start, 64)
}

var stepsPart2 = 26501365

func SolvePart2(_ context.Context, lines []string) any {
	return CountReachableInfiniteSmart(lines, stepsPart2)
}

func CountReachable(field fld.ByteField, start fld.Pos, steps int) int {
	outEdges := func(pos fld.Pos) (edges []graphs.OutEdge[fld.Pos]) {
		for _, dir := range fld.DirsSimple {
			npos := pos.Add(dir)
			if field.Inside(npos) && field.Get(npos) != '#' {
				edges = append(edges, graphs.OutEdge[fld.Pos]{To: npos, Cost: 1})
			}
		}
		return edges
	}
	res := 0
	for _, path := range graphs.DijkstraHeap(start, outEdges, &steps) {
		if path.MinCost <= steps && (path.MinCost+steps)%2 == 0 {
			res++
		}
	}
	return res
}

func CountReachableInfiniteNaive(lines []string, steps int) int {
	field := fld.NewByteField(lines)
	start := field.FindFirst('S')
	normPos := func(pos fld.Pos) fld.Pos {
		pos.Row %= field.Rows()
		if pos.Row < 0 {
			pos.Row += field.Rows()
		}
		pos.Col %= field.Cols()
		if pos.Col < 0 {
			pos.Col += field.Cols()
		}
		return pos
	}
	outEdges := func(pos fld.Pos) []graphs.OutEdge[fld.Pos] {
		var edges []graphs.OutEdge[fld.Pos]
		for _, dir := range fld.DirsSimple {
			npos := pos.Add(dir)
			if field.Get(normPos(npos)) != '#' {
				edges = append(edges, graphs.OutEdge[fld.Pos]{To: npos, Cost: 1})
			}
		}
		return edges
	}
	ans := 0
	paths := graphs.DijkstraHeap(start, outEdges, &steps)
	for _, path := range paths {
		must.LessOrEqual(path.MinCost, steps)
		if (path.MinCost+steps)%2 == 0 {
			ans++
		}
	}
	return ans
}

func CountReachableInfiniteSmart(lines []string, steps int) int {
	field := fld.NewByteField(lines)
	start := field.FindFirst('S')
	// 26501365: 65 steps to reach all 4 edges.
	// 26501300 / 131 = 202300 tiles are reachable in 4 directions, adjacent tiles has different parity.

	// N = 65

	// Original tile minimal distances on corners:
	//
	// 2N-1N-2N
	//  |  |  |
	//  N- 0 -N
	//  |  |  |
	// 2N-1N-2N
	//
	//
	// 3x3 tiles minimal distances in corners (+ one tile in each direction):
	//
	// 6N -- 4N-3N-4N -- 6N
	//  |     |  |  |     |
	//  |     |  |  |     |
	// 4N -- 2N-1N-2N -- 4N
	//  |     |  |  |     |
	// 3N --  N- 0 -N -- 3N
	//  |     |  |  |     |
	// 4N -- 2N-1N-2N -- 4N
	//  |     |  |  |     |
	//  |     |  |  |     |
	// 6N -- 4N-3N-4N -- 6N
	//
	//
	// 5x5 tiles minimal distances in corners (+ two tiles in each direction):
	// K = 2
	//
	//  5S --- 4S --- 3S----3S --- 4S ---  5S
	//   |      |      |  |  |      |       |
	//   |  e   |  O   |  E  |  O   |  e    |
	//   |      |      |  |  |      |       |
	//  4S --- 3S --- 2S----2S --- 3S ---  4S
	//   |      |      |  |  |      |       |
	//   |  O   |  E   |  o  |  E   |  O    |
	//   |      |      |  |  |      |       |
	//  3S --- 2S ---  S- N- S --- 2S ---  3S
	//   |      |      |  |  |      |       |
	//   | -E-  | -o-  N- 0 -N -O-  | -E-   |
	//   |      |      |  |  |      |       |
	//  3S --- 2S ---  S- N- S --- 2S ---  3S
	//   |      |      |  |  |      |       |
	//   |  O   |  E   |  o  |  E   |  O    |
	//   |      |      |  |  |      |       |
	//  4S --- 3S --- 2S----2S --- 3S ---  4S
	//   |      |      |  |  |      |       |
	//   |  e   |  O   |  E  |  O   |  e    |
	//   |      |      |  |  |      |       |
	//  5S --- 4S --- 3S----3S --- 4S ---  5S
	//
	// We need 202300 tiles in each direction (404601x404601 tiles).
	//
	size := field.Cols()
	must.Equal(size, field.Rows())
	must.Equal(size%2, 1)
	// steps := size/2 + size*10
	tiles := (steps - size/2) / size // Number of extra tiles available in one direction.
	must.Equal(start, fld.NewPos(size/2, size/2))
	must.Equal(size, field.Rows())
	must.Equal(size, field.Cols())
	ans := 0
	ans += CountReachable(field, fld.NewPos(size/2, size-1), size-1) // Leftmost
	ans += CountReachable(field, fld.NewPos(size/2, 0), size-1)      // Rightmost
	ans += CountReachable(field, fld.NewPos(size-1, size/2), size-1) // Upmost
	ans += CountReachable(field, fld.NewPos(0, size/2), size-1)      // Downmost
	// Left-Top side:
	ans += tiles * CountReachable(field, fld.NewPos(size-1, size-1), size/2-1)
	ans += (tiles - 1) * CountReachable(field, fld.NewPos(size-1, size-1), 3*(size/2))
	// Right-Top side:
	ans += tiles * CountReachable(field, fld.NewPos(size-1, 0), size/2-1)
	ans += (tiles - 1) * CountReachable(field, fld.NewPos(size-1, 0), 3*(size/2))
	// Left-Bottom side:
	ans += tiles * CountReachable(field, fld.NewPos(0, size-1), size/2-1)
	ans += (tiles - 1) * CountReachable(field, fld.NewPos(0, size-1), 3*(size/2))
	// Right-Bottom side:
	ans += tiles * CountReachable(field, fld.NewPos(0, 0), size/2-1)
	ans += (tiles - 1) * CountReachable(field, fld.NewPos(0, 0), 3*(size/2))
	// Not full tiles.
	// K = tiles
	// There are (K-1) full tiles to each direction.
	// k=2:
	//  1
	// 101
	//  1
	// k=4:
	//    1
	//   101
	//  10101
	// 1010101
	//  10101
	//   101
	//    1
	// n of odds: sum(1..k) + sum(1..k-1) = k*(k+1)/2 + k*(k-1)/2. odd(2) = 4, odd(4) = 16
	// Even nodes: odd - 2k + 1. even(2) = 1, even(4) = 9.
	odds := tiles * tiles
	// evens := odds - 2*tiles + 1
	evens := (tiles - 1) * (tiles - 1)
	if tiles%2 == 1 {
		odds, evens = evens, odds
	}
	ans += evens * CountReachable(field, start, steps)
	ans += odds * CountReachable(field, fld.NewPos(0, size/2), steps)
	return ans
}

func SolvePart2Naive(_ context.Context, lines []string) any {
	return CountReachableInfiniteNaive(lines, stepsPart2)
}

func SolvePart2Quadratic(_ context.Context, lines []string) any {
	return CountReachableInfiniteQuadratic(lines, stepsPart2)
}

func CountReachableInfiniteQuadratic(lines []string, steps int) int {
	field := fld.NewByteField(lines)
	start := field.FindFirst('S')
	size := field.Rows()
	factor := 5
	for row, line := range field {
		field[row] = bytes.Repeat(line, factor)
	}
	for i := 1; i < factor; i++ {
		field = append(field, field[:size]...)
	}
	start = start.Add(fld.NewPos((factor/2)*size, (factor/2)*size))
	// Assume that it's quadratic polynom:
	// count(t*size + size/2) = a*t*t + b*t + c
	p0 := CountReachable(field, start, size/2)
	p1 := CountReachable(field, start, 3*size/2)
	p2 := CountReachable(field, start, 5*size/2)
	c := p0
	a := (p0 - 2*p1 + p2) / 2
	b := p1 - p0 - a
	predict := func(t int) int {
		return a*t*t + b*t + c
	}
	// fmt.Println(CountReachable(field, start, size/2), predict(0))
	// fmt.Println(CountReachable(field, start, size+size/2), predict(1))
	// fmt.Println(CountReachable(field, start, 2*size+size/2), predict(2))
	// fmt.Println(CountReachable(field, start, 3*size+size/2), predict(3))
	// fmt.Println(CountReachable(field, start, 4*size+size/2), predict(4))
	return predict(steps / size)
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2, SolvePart2Quadratic}
)

func main() {
	aoc.Main(solvers1, append(solvers2, SolvePart2Naive))
}
