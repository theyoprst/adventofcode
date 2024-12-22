package main

import (
	"context"
	"math"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
	"github.com/theyoprst/adventofcode/aoc/graphs"
	"github.com/theyoprst/adventofcode/must"
)

// Cheatsheet:
//
// Human readable regex:
//   rex.New(rex.Common.RawVerbose(``)).MustCompile()
//

type Node struct {
	Pos fld.Pos
	Dir fld.Pos
}

func SolvePart1(_ context.Context, lines []string) any {
	field := fld.NewByteField(lines)
	start := field.FindFirst('S')
	end := field.FindFirst('E')

	startNode := Node{Pos: start, Dir: fld.East}
	minPaths := graphs.DijkstraHeap(startNode, func(v Node) []graphs.OutEdge[Node] {
		var edges []graphs.OutEdge[Node]
		for _, dir := range fld.DirsSimple {
			next := Node{
				Pos: v.Pos.Add(dir),
				Dir: dir,
			}
			if field.Get(next.Pos) == '#' {
				continue
			}
			edges = append(edges, graphs.OutEdge[Node]{
				To:   next,
				Cost: cost(v, next),
			})
		}
		return edges
	}, nil)
	shortest := math.MaxInt
	for _, dir := range fld.DirsSimple {
		node := Node{Pos: end, Dir: dir}
		if _, ok := minPaths[node]; ok {
			shortest = min(shortest, minPaths[node].MinCost)
		}
	}
	must.NotEqual(shortest, math.MaxInt)
	return shortest
}

func SolvePart2(_ context.Context, lines []string) any {
	field := fld.NewByteField(lines)
	start := field.FindFirst('S')
	end := field.FindFirst('E')

	startNode := Node{Pos: start, Dir: fld.East}
	minPaths := graphs.DijkstraHeap(startNode, func(v Node) []graphs.OutEdge[Node] {
		var edges []graphs.OutEdge[Node]
		for _, dir := range fld.DirsSimple {
			next := Node{
				Pos: v.Pos.Add(dir),
				Dir: dir,
			}
			if field.Get(next.Pos) == '#' {
				continue
			}
			edges = append(edges, graphs.OutEdge[Node]{
				To:   next,
				Cost: cost(v, next),
			})
		}
		return edges
	}, nil)

	shortest := math.MaxInt
	for _, dir := range fld.DirsSimple {
		node := Node{Pos: end, Dir: dir}
		if _, ok := minPaths[node]; ok {
			shortest = min(shortest, minPaths[node].MinCost)
		}
	}
	must.NotEqual(shortest, math.MaxInt)

	visited := containers.Set[Node]{}
	onBestPaths := containers.Set[fld.Pos]{}

	var dfs func(node Node, wantMinDist int)
	dfs = func(node Node, wantMinDist int) {
		if visited.Has(node) {
			return
		}
		visited.Add(node)
		if _, ok := minPaths[node]; !ok {
			return
		}
		if minPaths[node].MinCost != wantMinDist {
			return
		}
		onBestPaths.Add(node.Pos)
		prev := Node{Pos: node.Pos.Sub(node.Dir)}
		for _, dir := range fld.DirsSimple {
			prev.Dir = dir
			dfs(prev, wantMinDist-cost(prev, node))
		}
	}

	for _, dir := range fld.DirsSimple {
		node := Node{Pos: end, Dir: dir}
		dfs(node, shortest)
	}
	return len(onBestPaths)
}

func cost(from, to Node) int {
	diff := from.Pos.Sub(to.Pos)
	must.Equal(aoc.Abs(diff.Row)+aoc.Abs(diff.Col), 1)
	if from.Dir == to.Dir {
		return 1
	}
	return 1001
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
