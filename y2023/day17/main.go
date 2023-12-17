package main

import (
	"fmt"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/fld"
	"github.com/theyoprst/adventofcode/aoc/queues"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(lines []string) any {
	return SolveGeneric(lines, 0, 3)
}

func SolvePart2(lines []string) any {
	return SolveGeneric(lines, 4, 10)
}

func SolveGeneric(lines []string, minSteps, maxSteps int) any {
	field := fld.NewByteField(lines)
	type Vertex struct {
		pos   fld.Pos
		dir   fld.Pos
		steps int // steps done before in that dir
	}
	minCosts := map[Vertex]int{}
	pq := queues.NewPriorityQueue[Vertex, int]()
	pq.Insert(Vertex{steps: minSteps}, 0)
	from := map[Vertex]Vertex{}
	for pq.Len() > 0 {
		minV, cost := pq.PopMin()
		minCosts[minV] = cost
		if minV.pos == fld.NewPos(field.Rows()-1, field.Cols()-1) && minSteps <= minV.steps && minV.steps <= maxSteps {
			v := minV
			for v.pos != fld.NewPos(0, 0) {
				// fmt.Println("Back to start:", v)
				must.Equal(from[v].pos.Add(v.dir), v.pos)
				v = from[v]
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
			}
			fmt.Println(fld.ToString(field))
			return cost
		}
		for _, dir := range []fld.Pos{fld.Left, fld.Right, fld.Up, fld.Down} {
			pos := minV.pos.Add(dir)
			if !field.Inside(pos) {
				continue
			}
			if dir == minV.dir.Mult(-1) {
				continue // Forbid turn-over.
			}
			if dir != minV.dir && minV.steps < minSteps {
				continue
			}
			steps := 1
			if dir == minV.dir {
				steps += minV.steps
				if steps > maxSteps {
					continue
				}
			}
			v := Vertex{
				pos:   pos,
				dir:   dir,
				steps: steps,
			}
			if minCosts[v] > 0 {
				continue
			}
			newCost := cost + int(field.Get(pos)-'0')
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
	panic("unreachable")
}

var (
	solvers1 []aoc.Solver = []aoc.Solver{SolvePart1}
	solvers2 []aoc.Solver = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
