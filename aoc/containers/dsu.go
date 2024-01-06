package containers

import "maps"

type DisjointSet[T comparable] struct {
	parent map[T]T
	count  map[T]int
	n      int
}

func NewDisjointSet[T comparable]() DisjointSet[T] {
	return DisjointSet[T]{
		parent: map[T]T{},
		count:  map[T]int{},
		n:      0,
	}
}

func (ds *DisjointSet[T]) Clone() DisjointSet[T] {
	return DisjointSet[T]{
		parent: maps.Clone(ds.parent),
		count:  maps.Clone(ds.count),
		n:      ds.n,
	}
}

func (ds *DisjointSet[T]) Components() int {
	return ds.n
}

func (ds *DisjointSet[T]) Has(value T) bool {
	_, ok := ds.parent[value]
	return ok
}

func (ds *DisjointSet[T]) Add(value T) {
	ds.parent[value] = value
	ds.count[value] = 1
	ds.n++
}

func (ds *DisjointSet[T]) Root(value T) T {
	if ds.parent[value] != value {
		ds.parent[value] = ds.Root(ds.parent[value])
	}
	return ds.parent[value]
}

func (ds *DisjointSet[T]) Union(first, second T) {
	big := ds.Root(first)
	small := ds.Root(second)
	if big == small {
		return
	}
	if ds.count[big] < ds.count[small] {
		big, small = small, big
	}
	ds.parent[small] = big
	ds.count[big] += ds.count[small]
	ds.n--
}

func (ds *DisjointSet[T]) Size(value T) int {
	return ds.count[ds.Root(value)]
}
