package main

import (
	"context"
	"maps"
	"slices"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(_ context.Context, lines []string) any {
	graph := make(map[string]containers.Set[string])
	for _, line := range lines {
		n1, n2 := must.Split2(line, "-")
		graph[n1] = graph[n1].Add(n2)
		graph[n2] = graph[n2].Add(n1)
	}

	triangles := containers.NewSet[string]()
	for node, neighbors := range graph {
		if node[0] != 't' {
			continue
		}
		for second := range neighbors {
			thirds := graph[second].Intersection(neighbors)
			for third := range thirds {
				triangle := []string{node, second, third}
				slices.Sort(triangle)
				triangles.Add(strings.Join(triangle, ","))
			}
		}
	}
	triList := triangles.Slice()
	slices.Sort(triList)
	return len(triList)
}

func SolvePart2(_ context.Context, lines []string) any {
	graph := make(map[string]containers.Set[string])
	for _, line := range lines {
		n1, n2 := must.Split2(line, "-")
		graph[n1] = graph[n1].Add(n2)
		graph[n2] = graph[n2].Add(n1)
	}

	maxCliqueSize := 0
	var maxClique string

	// In original Bron-Kerbosch algorithm, they also use "not" set, but they need it for tracking of all maximal cliques,
	// that is the cliques which cannot be extended by adding more vertices, they are not necessarily the largest cliques.
	// In this problem, we need only the largest clique, so we can omit this, saving a bit of compute power.
	var bronKerbosch func(yes, maybe containers.Set[string])
	bronKerbosch = func(yes, maybe containers.Set[string]) {
		if len(maybe) == 0 {
			if len(yes) > maxCliqueSize {
				maxCliqueSize = len(yes)
				items := yes.Slice()
				slices.Sort(items)
				maxClique = strings.Join(items, ",")
			}
			return
		}
		// Pivoting (skip) speeds up ~9 times (9ms -> 1ms).
		skip := graph[maybe.Any()]
		for node := range maybe.Difference(skip) {
			bronKerbosch(yes.Add(node), maybe.Intersection(graph[node]))
			yes.Remove(node)
			maybe.Remove(node)
		}
	}

	vertices := slices.Collect(maps.Keys(graph))
	bronKerbosch(containers.NewSet[string](), containers.NewSet(vertices...))

	return maxClique
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
