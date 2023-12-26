package main

import (
	"log"
	"math"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/must"
)

// General considerations:
// We know that minimum cut of that graph is 3 (otherwise we would have more than 1 answer which is not common in AOC).
// And we need to find this cut, at least we need sizes of two vertices sets.
// It is also well-known that Min Cut == Max Flow for any graph.
// We have bunch of algorithms finding Max Flow, although most of them are for oriented graphs and for designated nodes
// s (source) and t (sink).
// Our graph is bidirectional and capacity of each edge C(e) = 1.
// Also our graph is sparse: |E| ~= 3 * |V|.

// SolvePart1FFA solves part 1 using Ford-Fulkerson Algorithm / method (FFA).
// FFA finds Max Flow from s to t running DFS at maximum sum(C) times, each time it increases flow value at least by 1.
// But we are interested only in that pairs (s, t) which have max flow 3 between them. Which means that if inside FFA we reached flow value > 3,
// we can forget about (s, t) and move on to another pair of vertices.
// The final algorithm is:
// - fix an arbitrary vertex `s`
// - iterate over all the rest vertices `t`:
// - check if flow between `s` and `t` is less than 4, running DFS not more than 4 times
// - if max flow is 3, FFA have found min cut with value 3, vertices from the first component (reachable from s) are marked as "seen" after last DFS.
// - return |seen| * (|V| - |seen|)
// Complexity is `O(|E| * |V|)`.
// But in test data graph is split in two roughly equal parts. Which means that FFA will be run on averate 2 times, so
// Expected running time is `O(|E|)`.
func SolvePart1FFA(lines []string) any {
	type Edge struct{ from, to string }
	cap := map[Edge]int{}
	graph := map[string]containers.Set[string]{}
	for _, line := range lines {
		first, rest := must.Split2(line, ": ")
		seconds := strings.Split(rest, " ")
		for _, second := range seconds {
			cap[Edge{first, second}] = 1
			cap[Edge{second, first}] = 1
			graph[first] = graph[first].Add(second)
			graph[second] = graph[second].Add(first)
		}
	}
	var s string
	for v := range graph {
		s = v
		break
	}
	flow := map[Edge]int{}
	seen := containers.NewSet[string]()
	var t string
	var dfs func(v string, inc int) int
	dfs = func(v string, curMin int) int {
		if seen.Has(v) {
			return 0
		}
		seen.Add(v)
		if v == t {
			return curMin
		}
		for u := range graph[v] {
			edge := Edge{v, u}
			residualCap := cap[edge] - flow[edge]
			if residualCap > 0 {
				totalMin := dfs(u, min(curMin, residualCap))
				if totalMin > 0 {
					flow[edge] += totalMin
					flow[Edge{u, v}] -= totalMin
					return totalMin
				}
			}
		}
		return 0
	}
	const maxFlow = 3
	for t = range graph {
		if t == s {
			continue
		}
		log.Printf("Checking max flow from %q to %q", s, t)
		flowSize := 0
		clear(flow)
		for {
			clear(seen)
			dFlow := dfs(s, math.MaxInt)
			if dFlow == 0 {
				break
			}
			flowSize += dFlow
			if flowSize > maxFlow {
				break
			}
		}
		if flowSize == maxFlow {
			// Found max flow and min cut.
			sLen := len(seen)
			tLen := len(graph) - sLen
			log.Printf("Found two components: %d, %d", sLen, tLen)
			return sLen * tLen
		}
	}
	panic("Unreachable")
}

var (
	solvers1 = []aoc.Solver{SolvePart1FFA}
	solvers2 = []aoc.Solver{}
)

func main() {
	log.SetFlags(0)
	aoc.Main(solvers1, solvers2)
}
