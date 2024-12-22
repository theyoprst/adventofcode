package main

import (
	"context"
	"log"
	"math"
	"math/rand"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/queues"
	"github.com/theyoprst/adventofcode/must"
)

const (
	maxFlow   = 3
	checkAllT = false // Check all t instead of early exit when found some s/t cut of size 3.
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
func SolvePart1FFA(_ context.Context, lines []string) any {
	type Edge struct{ from, to string }
	c := map[Edge]int{}
	graph := map[string]containers.Set[string]{}
	for _, line := range lines {
		first, rest := must.Split2(line, ": ")
		seconds := strings.Split(rest, " ")
		for _, second := range seconds {
			c[Edge{first, second}] = 1
			c[Edge{second, first}] = 1
			graph[first] = graph[first].Add(second)
			graph[second] = graph[second].Add(first)
		}
	}

	var s string // source node
	for v := range graph {
		s = v
		break
	}
	var t string // sink node

	flow := map[Edge]int{}
	seen := containers.NewSet[string]()

	var dfsFFA func(v string, inc int) int
	dfsFFA = func(v string, curMin int) int {
		if seen.Has(v) {
			return 0
		}
		seen.Add(v)
		if v == t {
			return curMin
		}
		for u := range graph[v] {
			edge := Edge{v, u}
			residualCap := c[edge] - flow[edge]
			if residualCap > 0 {
				totalMin := dfsFFA(u, min(curMin, residualCap))
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

	var c1, c2 int
	for t = range graph {
		if t == s {
			continue
		}
		log.Printf("Checking max flow from %q to %q", s, t)
		flowSize := 0
		clear(flow)
		for flowSize <= maxFlow {
			clear(seen)
			dFlow := dfsFFA(s, math.MaxInt)
			if dFlow == 0 {
				break
			}
			flowSize += dFlow
		}
		if flowSize == maxFlow {
			// Found max flow and min cut.
			c1 = len(seen)
			c2 = len(graph) - c1
			if !checkAllT {
				break
			}
		}
	}
	log.Printf("Found two components: %d, %d", c1, c2)
	return c1 * c2
}

// Same as FFA, but BFS instead of DFS was used. 20% faster than FFA because a more short paths are chosen.
func SolvePart1EdmondsKarp(_ context.Context, lines []string) any {
	type Edge struct{ from, to string }
	c := map[Edge]int{}
	graph := map[string]containers.Set[string]{}
	for _, line := range lines {
		first, rest := must.Split2(line, ": ")
		seconds := strings.Split(rest, " ")
		for _, second := range seconds {
			c[Edge{first, second}] = 1
			c[Edge{second, first}] = 1
			graph[first] = graph[first].Add(second)
			graph[second] = graph[second].Add(first)
		}
	}

	s := anyKey(graph)

	var c1, c2 int
	for t := range graph { // Sink node.
		if t == s {
			continue
		}
		log.Printf("Checking max flow from %q to %q", s, t)

		// Run Edmunds-Karp.
		flow := map[Edge]int{}
		flowSize := 0
		var prev map[string]string
		for flowSize <= maxFlow { // Early exit if flow is too large, i.e. bigger than 3.
			// Run BFS.
			prev = map[string]string{s: ""}
			queue := []string{s}
			for prev[t] == "" && len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]

				for next := range graph[cur] {
					edge := Edge{cur, next}
					if next != s && prev[next] == "" && c[edge] > flow[edge] {
						prev[next] = cur
						queue = append(queue, next)
					}
				}
			}
			if prev[t] == "" {
				break // No augmented path found by the BFS.
			}
			minResidualCap := math.MaxInt
			for v := t; v != s; v = prev[v] {
				edge := Edge{prev[v], v}
				minResidualCap = min(minResidualCap, c[edge]-flow[edge])
			}
			for v := t; v != s; v = prev[v] {
				flow[Edge{prev[v], v}] += minResidualCap
				flow[Edge{v, prev[v]}] -= minResidualCap
			}
			flowSize += minResidualCap
		}
		if flowSize == maxFlow {
			// Found max flow and min cut.
			c1 = len(prev)
			c2 = len(graph) - c1
			if !checkAllT {
				break
			}
		}
	}
	log.Printf("Found two components: %d, %d", c1, c2)
	return c1 * c2
}

type Edge struct {
	from, to string
}

func (e Edge) Reversed() Edge {
	return Edge{from: e.to, to: e.from}
}

// For some reason it's 5 times slower than Edmonds-Karp if check all sink nodes t.
// Maybe because BFS finds t on average much more faster than traversing all the graph for building layered network.
func SolvePart1Dinic(_ context.Context, lines []string) any {
	c := map[Edge]int{}
	graph := map[string]containers.Set[string]{}
	for _, line := range lines {
		first, rest := must.Split2(line, ": ")
		seconds := strings.Split(rest, " ")
		for _, second := range seconds {
			c[Edge{first, second}] = 1
			c[Edge{second, first}] = 1
			graph[first] = graph[first].Add(second)
			graph[second] = graph[second].Add(first)
		}
	}

	var comp1, comp2 int

	const inf = 1000000000

	s := anyKey(graph)
	for t := range graph {
		flow := map[Edge]int{}
		flowVal := 0
		if t == s {
			continue
		}
	phase:
		for {
			// Find minimal distance to each vertex.
			dist := map[string]int{}
			for v := range graph {
				dist[v] = inf
			}
			dist[s] = 0
			reachable := 1
			queue := []string{s}
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				for next := range graph[cur] {
					edge := Edge{cur, next}
					if flow[edge] < c[edge] && dist[next] == inf {
						dist[next] = dist[cur] + 1
						queue = append(queue, next)
						reachable++
					}
				}
			}
			if dist[t] == inf { // If t is not reachable, exit from phases loop.
				if flowVal == maxFlow {
					comp1 = reachable
					comp2 = len(graph) - comp1
				}
				break phase
			}
			// Build layered network.
			next := map[string][]string{}
			for u := range graph {
				for v := range graph[u] {
					if dist[u]+1 == dist[v] {
						next[u] = append(next[u], v)
					}
				}
			}

			// Find paths in layered network
			var dfs func(string, int) int
			dfs = func(v string, curMin int) int { // Returns added flow value.
				if curMin == 0 {
					return 0
				}
				if v == t {
					return curMin
				}
				for len(next[v]) > 0 {
					u := next[v][0]
					edge := Edge{v, u}
					totalMin := dfs(u, min(curMin, c[edge]-flow[edge]))
					if totalMin == 0 {
						next[v] = next[v][1:]
						continue
					}
					flow[edge] += totalMin
					flow[edge.Reversed()] -= totalMin
					return totalMin
				}
				return 0
			}

			for {
				diff := dfs(s, math.MaxInt)
				if diff == 0 {
					continue phase
				}
				flowVal += diff
				if flowVal > maxFlow {
					break phase
				}
			}
		} // Phase.
		if flowVal == maxFlow && !checkAllT {
			break
		}
	} // Iterate over t.
	return comp1 * comp2
}

// Stoer-Wagner algorithm: O(V) phases, each O(E log V) by using binary heap queue.
// Total complexity: O(V E log V).
// To find minimal cut = 3 it has to reduce graph by 80%, there is (probably) no way to reduce it.
// The slowest solution for this problem.
func SolveStoerWagner(_ context.Context, lines []string) any {
	graph := map[string]map[string]int{}
	for _, line := range lines {
		first, rest := must.Split2(line, ": ")
		seconds := strings.Split(rest, " ")
		if _, ok := graph[first]; !ok {
			graph[first] = map[string]int{}
		}
		for _, second := range seconds {
			graph[first][second] = 1
			if _, ok := graph[second]; !ok {
				graph[second] = map[string]int{}
			}
			graph[second][first] = 1
		}
	}
	n := len(graph)

	merged := map[string][]string{}
	for v := range graph {
		merged[v] = []string{v}
	}
	for len(graph) >= 2 { // phase: found s and t
		a := containers.NewSet[string]()
		queue := queues.NewPriorityQueue[string, int]()
		for v := range graph {
			queue.Insert(v, 0)
		}
		prev := ""
		for i := 0; i < len(graph)-1; i++ {
			sel, _ := queue.PopMin()
			must.NotEqual(sel, "")
			a.Add(sel)
			for v, cap := range graph[sel] {
				if queue.Has(v) {
					queue.Inc(v, -cap)
				}
			}
			prev = sel
		}
		sel, minusW := queue.PopMin()
		must.NotEqual(sel, "")
		must.Equal(queue.Len(), 0)
		if minusW == -3 {
			comp1 := len(merged[sel])
			comp2 := n - comp1
			log.Printf("Found components: %d * %d after %d (max=%d) iterations", comp1, comp2, n-len(graph)+1, n-1)
			return comp1 * comp2
		}
		// Merge sel and prev
		must.NotEqual(sel, prev)
		for v, cap := range graph[prev] {
			if v != sel {
				graph[sel][v] += cap
				graph[v][sel] += cap
			}
			delete(graph[v], prev)
		}
		delete(graph, prev)
		merged[sel] = append(merged[sel], merged[prev]...)
	} // phase
	panic("unreachable")
}

// SolveKarger uses probabilistic Karger algorithm of finding minimal global cut.
// Disjoint Set union is used to merge vertices.
// Single phase has time O(E), theoretical upper bound for number of phases is O(V^2).
// But for this graph number of phases were always less than V (after several dozens of runs).
// So looks like it is O(V*E) for the test graph.
func SolveKarger(_ context.Context, lines []string) any {
	var edges []Edge
	vertices := containers.NewSet[string]()
	for _, line := range lines {
		first, rest := must.Split2(line, ": ")
		vertices.Add(first)
		seconds := strings.Split(rest, " ")
		for _, second := range seconds {
			edges = append(edges, Edge{first, second})
			vertices.Add(second)
		}
	}

	randoms := 0
	for phase := 0; true; phase++ {
		edges := edges
		dsu := containers.NewDisjointSet[string]()
		for v := range vertices {
			dsu.Add(v)
		}
		for dsu.Components() > 2 {
			randoms++
			k := rand.Intn(len(edges))
			edges[0], edges[k] = edges[k], edges[0]
			edge := edges[0]
			edges = edges[1:]
			dsu.Union(edge.from, edge.to)
		}
		cutSize := 0
		var comp1, comp2 int
		for _, edge := range edges {
			if dsu.Root(edge.from) != dsu.Root(edge.to) {
				cutSize++
				if cutSize > 3 {
					break
				}
				comp1 = dsu.Size(edge.from)
				comp2 = dsu.Size(edge.to)
			}
		}
		if cutSize == 3 {
			log.Printf("Found two components after %d phases, %d randomizations: %d * %d.", phase+1, randoms, comp1, comp2)
			return comp1 * comp2
		}
	} // phase
	panic("Unreachable")
}

// Recursive optimization of Karger algorithm: it doubles execution tree after contraction to
// n/sqrt(2), n/2, n/(2sqrt(2)), n/4, et cetera.
// It works very slow this way. But it works better if:
// - divide by 2: beats original Karger in 33% cases
// - divide by 3: beats original Karger in 85% cases
// Possible reasons:
// - one of the min cut edges was removed in the beginning and we spent lots of time in a wrong branch of recursion.
// - randomization is slow.
func SolveKargerStein(_ context.Context, lines []string) any {
	var edges []Edge
	vertices := containers.NewSet[string]()
	for _, line := range lines {
		first, rest := must.Split2(line, ": ")
		vertices.Add(first)
		seconds := strings.Split(rest, " ")
		for _, second := range seconds {
			edges = append(edges, Edge{first, second})
			vertices.Add(second)
		}
	}

	randoms := 0
	contract := func(edges []Edge, dsu *containers.DisjointSet[string], wantSize int) []Edge {
		for dsu.Components() > wantSize {
			k := rand.Intn(len(edges))
			randoms++
			edges[0], edges[k] = edges[k], edges[0]
			edge := edges[0]
			edges = edges[1:]
			dsu.Union(edge.from, edge.to)
		}
		return edges
	}

	var comp1, comp2 int

	var findMinCut func(edges []Edge, dsu *containers.DisjointSet[string]) bool
	findMinCut = func(edges []Edge, dsu *containers.DisjointSet[string]) bool {
		if dsu.Components() <= 6 {
			edges = contract(edges, dsu, 2)
			cutSize := 0
			for _, edge := range edges {
				if dsu.Root(edge.from) != dsu.Root(edge.to) {
					cutSize++
					if cutSize > 3 {
						break
					}
					comp1 = dsu.Size(edge.from)
					comp2 = dsu.Size(edge.to)
				}
			}
			return cutSize == 3
		}
		// Note: In original Karger-Stein there is division by sqrt(2). But in this problem it works too slow.
		// nextComponents := int(float64(dsu.Components()) / math.Sqrt2)
		nextComponents := int(float64(dsu.Components()) / 3)
		dsu2 := dsu.Clone()
		return findMinCut(contract(edges, dsu2, nextComponents), dsu2) ||
			findMinCut(contract(edges, dsu, nextComponents), dsu)
	}

	for phase := 1; true; phase++ {
		edges := edges
		dsu := containers.NewDisjointSet[string]()
		for v := range vertices {
			dsu.Add(v)
		}
		if findMinCut(edges, dsu) {
			log.Printf("Found two components after %d phases, %d randomizations: %d * %d.", phase, randoms, comp1, comp2)
			return comp1 * comp2
		}
	} // phase
	panic("Unreachable")
}

func anyKey[K comparable, V any](m map[K]V) K {
	for k := range m {
		return k
	}
	panic("map is empty")
}

var (
	solvers1 = []aoc.Solver{SolvePart1FFA, SolvePart1EdmondsKarp, SolvePart1Dinic, SolveStoerWagner, SolveKarger, SolveKargerStein}
	solvers2 = []aoc.Solver{}
)

func main() {
	log.SetFlags(0)
	aoc.Main(solvers1, solvers2)
}
