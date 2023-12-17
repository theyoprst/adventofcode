package graphs

import (
	"github.com/theyoprst/adventofcode/aoc/queues"
)

type OutEdge[V comparable] struct {
	To   V
	Cost int
}

type MinPath[V comparable] struct {
	MinCost int
	Prev    V
}

// DijkstraHeap returns minimal costs from startV to other vertices in implicit graph given by function outEdges().
// DijkstraHeap has time complexity O(E*log(V)) which is perfect for sparse graphs like planar graphs, e.g. E=O(V).
// If graph is dense, i.e. E=O(V^2), it has complexity O(V^2 log(V)) which is not perfect and implementation with
// naive priority queue should be used. Although for advent of code puzzles it should be OK.
func DijkstraHeap[V comparable](startV V, outEdges func(v V) []OutEdge[V]) map[V]MinPath[V] {
	pq := queues.NewPriorityQueue[V, int]()
	pq.Insert(startV, 0)
	res := map[V]MinPath[V]{}
	from := map[V]V{}
	for pq.Len() > 0 {
		minV, cost := pq.PopMin()
		res[minV] = MinPath[V]{
			MinCost: cost,
			Prev:    from[minV],
		}
		for _, edge := range outEdges(minV) {
			v := edge.To
			if _, ok := res[v]; ok {
				continue
			}
			newCost := cost + edge.Cost
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
