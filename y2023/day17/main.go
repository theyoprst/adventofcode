package main

import (
	"fmt"
	"math"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/fld"
	"github.com/theyoprst/adventofcode/aoc/graphs"
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

type Edge = graphs.OutEdge[Vertex]

func SolveGeneric(lines []string, minSteps, maxSteps int) any {
	field := fld.NewByteField(lines)
	outEdges := func(v Vertex) []Edge {
		var edges []Edge
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
			edges = append(edges, Edge{
				To: Vertex{
					pos:   pos,
					dir:   dir,
					steps: steps,
				},
				Cost: int(field.Get(pos) - '0'),
			})
		}
		return edges
	}

	minPaths := graphs.DijkstraHeap(Vertex{steps: minSteps}, outEdges)

	ans := math.MaxInt
	var minV Vertex
	bottomRight := fld.NewPos(field.Rows()-1, field.Cols()-1)
	for v, path := range minPaths {
		if v.pos == bottomRight && minSteps <= v.steps && v.steps <= maxSteps {
			if path.MinCost < ans {
				ans = path.MinCost
				minV = v
			}
			ans = min(ans, path.MinCost)
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
		v = minPaths[v].Prev
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
