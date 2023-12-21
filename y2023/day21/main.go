package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/fld"
	"github.com/theyoprst/adventofcode/aoc/graphs"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(lines []string) any {
	field := fld.NewByteField(lines)
	start := field.FindFirst('S')
	return CountReachable(field, start, 64)
}

var stepsPart2 int = 26501365

func SolvePart2(lines []string) any {
	field := fld.NewByteField(lines)
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
	S := field.Cols()
	must.Equal(S, field.Rows())
	N := S / 2
	// steps := N + S*10
	const steps = 26501365
	K := (steps - N) / (2*N + 1)
	start := field.FindFirst('S')
	must.Equal(start, fld.NewPos(N, N))
	must.Equal(S, field.Rows())
	must.Equal(S, field.Cols())
	ans := 0
	ans += CountReachable(field, fld.NewPos(N, 2*N), 2*N) // Leftmost
	ans += CountReachable(field, fld.NewPos(N, 0), 2*N)   // Rightmost
	ans += CountReachable(field, fld.NewPos(2*N, N), 2*N) // Upmost
	ans += CountReachable(field, fld.NewPos(0, N), 2*N)   // Downmost
	// Left-Top side:
	ans += K * CountReachable(field, fld.NewPos(2*N, 2*N), N-1)
	ans += (K - 1) * CountReachable(field, fld.NewPos(2*N, 2*N), 3*N)
	// Right-Top side:
	ans += K * CountReachable(field, fld.NewPos(2*N, 0), N-1)
	ans += (K - 1) * CountReachable(field, fld.NewPos(2*N, 0), 3*N)
	// Left-Bottom side:
	ans += K * CountReachable(field, fld.NewPos(0, 2*N), N-1)
	ans += (K - 1) * CountReachable(field, fld.NewPos(0, 2*N), 3*N)
	// Right-Bottom side:
	ans += K * CountReachable(field, fld.NewPos(0, 0), N-1)
	ans += (K - 1) * CountReachable(field, fld.NewPos(0, 0), 3*N)
	// Not full tiles.
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
	odds := K*(K+1)/2 + K*(K-1)/2
	evens := odds - 2*K + 1
	ans += evens * CountReachable(field, start, steps)
	ans += odds * CountReachable(field, fld.NewPos(0, N), steps)
	return ans
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

func CountReachableInfiniteNaive(field fld.ByteField, start fld.Pos, steps int) int {
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
	paths := graphs.DijkstraHeap(start, outEdges, &stepsPart2)
	for _, path := range paths {
		must.LessOrEqual(path.MinCost, stepsPart2)
		if path.MinCost%2 == stepsPart2%2 {
			ans++
		}
	}
	return ans
}

func SolvePart2Naive(lines []string) any {
	field := fld.NewByteField(lines)
	start := field.FindFirst('S')
	return CountReachableInfiniteNaive(field, start, stepsPart2)
}

var (
	solvers1 []aoc.Solver = []aoc.Solver{SolvePart1}
	solvers2 []aoc.Solver = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, append(solvers2, SolvePart2Naive))
}
