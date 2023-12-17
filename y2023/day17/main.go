package main

import (
	"math"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/fld"
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
	found := map[Vertex]bool{}
	type Node struct {
		Vertex
		cost int
	}
	nodes := []Node{{
		Vertex: Vertex{steps: minSteps},
	}}
	from := map[Vertex]Vertex{}
	for len(nodes) > 0 {
		minNode := Node{cost: math.MaxInt}
		minI := -1
		for i, node := range nodes {
			if node.cost < minNode.cost {
				minNode = node
				minI = i
			}
		}
		nodes = slices.Delete(nodes, minI, minI+1)

		found[minNode.Vertex] = true
		if minNode.pos == fld.NewPos(field.Rows()-1, field.Cols()-1) {
			// v := minNode.Vertex
			// for v.pos != fld.NewPos(0, 0) {
			// 	// fmt.Println("Back to start:", v)
			// 	must.Equal(from[v].pos.Add(v.dir), v.pos)
			// 	v = from[v]
			// 	switch v.dir {
			// 	case fld.Right:
			// 		field.Set(v.pos, '>')
			// 	case fld.Left:
			// 		field.Set(v.pos, '<')
			// 	case fld.Up:
			// 		field.Set(v.pos, '^')
			// 	case fld.Down:
			// 		field.Set(v.pos, 'v')
			// 	}
			// }
			// fmt.Println(fld.ToString(field))
			return minNode.cost
		}
		for _, dir := range []fld.Pos{fld.Left, fld.Right, fld.Up, fld.Down} {
			pos := minNode.pos.Add(dir)
			if !field.Inside(pos) {
				continue
			}
			if dir == minNode.dir.Mult(-1) {
				continue // Forbid turn-over.
			}
			if dir != minNode.dir && minNode.steps < minSteps {
				continue
			}
			steps := 1
			if dir == minNode.dir {
				steps += minNode.steps
				if steps > maxSteps {
					continue
				}
			}
			v := Vertex{
				pos:   pos,
				dir:   dir,
				steps: steps,
			}
			must.LessOrEqual(steps, maxSteps)
			if found[v] {
				continue
			}
			newCost := minNode.cost + int(field.Get(pos)-'0')
			nodeI := slices.IndexFunc(nodes, func(node Node) bool {
				return node.Vertex == v
			})
			if nodeI != -1 {
				if nodes[nodeI].cost > newCost {
					nodes[nodeI].cost = newCost
					from[v] = minNode.Vertex
				}
			} else {
				nodes = append(nodes, Node{
					Vertex: v,
					cost:   newCost,
				})
				from[v] = minNode.Vertex
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
