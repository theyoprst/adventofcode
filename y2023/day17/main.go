package main

import (
	"fmt"
	"math"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/fld"
	"github.com/theyoprst/adventofcode/aoc/queues"
)

func SolvePart1(lines []string) any {
	return SolveGeneric(lines, 0, 3)
}

func SolvePart2(lines []string) any {
	return SolveGeneric(lines, 4, 10)
}

type Vertex struct {
	pos   fld.Pos
	dir   fld.Pos
	steps int // steps done before in that dir
}

type OutEdge struct {
	to   Vertex
	cost int
}

type MinPath struct {
	cost int
	prev Vertex
}

func Dijkstra(startV Vertex, outEdges func(v Vertex) []OutEdge) map[Vertex]MinPath {
	pq := queues.NewPriorityQueue[Vertex, int]()
	pq.Insert(startV, 0)
	res := map[Vertex]MinPath{}
	from := map[Vertex]Vertex{}
	for pq.Len() > 0 {
		minV, cost := pq.PopMin()
		res[minV] = MinPath{
			cost: cost,
			prev: from[minV],
		}
		for _, edge := range outEdges(minV) {
			v := edge.to
			if _, ok := res[v]; ok {
				continue
			}
			newCost := cost + edge.cost
			nodeI, curCost := pq.Lookup(v)
			if nodeI != -1 {
				if newCost < curCost {
					pq.SetByIndex(nodeI, newCost)
					from[v] = minV
				}
			} else {
				pq.Insert(v, newCost)
				from[v] = minV
			}
		}
	}
	return res
}

func SolveGeneric(lines []string, minSteps, maxSteps int) any {
	field := fld.NewByteField(lines)
	outEdges := func(v Vertex) []OutEdge {
		var edges []OutEdge
		for _, dir := range []fld.Pos{fld.Left, fld.Right, fld.Up, fld.Down} {
			pos := v.pos.Add(dir)
			if !field.Inside(pos) {
				continue
			}
			if dir == v.dir.Mult(-1) {
				continue // Forbid turn-over.
			}
			if dir != v.dir && v.steps < minSteps {
				continue
			}
			steps := 1
			if dir == v.dir {
				steps += v.steps
				if steps > maxSteps {
					continue
				}
			}
			edges = append(edges, OutEdge{
				to: Vertex{
					pos:   pos,
					dir:   dir,
					steps: steps,
				},
				cost: int(field.Get(pos) - '0'),
			})
		}
		return edges
	}

	minPaths := Dijkstra(Vertex{steps: minSteps}, outEdges)

	ans := math.MaxInt
	var minV Vertex
	bottomRight := fld.NewPos(field.Rows()-1, field.Cols()-1)
	for v, path := range minPaths {
		if v.pos == bottomRight && minSteps <= v.steps && v.steps <= maxSteps {
			if path.cost < ans {
				ans = path.cost
				minV = v
			}
			ans = min(ans, path.cost)
		}
	}

	// Optional: print path.
	v := minV
	for v.pos != fld.NewPos(0, 0) {
		switch v.dir {
		case fld.Right:
			field.Set(v.pos, '>')
		case fld.Left:
			field.Set(v.pos, '<')
		case fld.Up:
			field.Set(v.pos, '^')
		case fld.Down:
			field.Set(v.pos, 'v')
		}
		v = minPaths[v].prev
	}
	fmt.Println()
	fmt.Println(fld.ToString(field))

	return ans
}

var (
	solvers1 []aoc.Solver = []aoc.Solver{SolvePart1}
	solvers2 []aoc.Solver = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
