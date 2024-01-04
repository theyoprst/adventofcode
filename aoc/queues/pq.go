package queues

import (
	"github.com/theyoprst/adventofcode/must"
	"golang.org/x/exp/constraints"
)

type PriorityType interface {
	constraints.Integer | constraints.Float
}

type Item[T comparable, P PriorityType] struct {
	value    T
	priority P
}

type PriorityQueue[T comparable, P PriorityType] struct {
	items  []Item[T, P]
	lookup map[T]int
}

func NewPriorityQueue[T comparable, P PriorityType]() PriorityQueue[T, P] {
	return PriorityQueue[T, P]{
		lookup: map[T]int{},
	}
}

func (q PriorityQueue[T, P]) Len() int {
	return len(q.items)
}

func (q PriorityQueue[T, P]) Swap(i, j int) {
	vi, vj := q.items[i], q.items[j]
	q.items[i], q.items[j] = vj, vi
	q.lookup[vi.value], q.lookup[vj.value] = q.lookup[vj.value], q.lookup[vi.value]
}

func (q *PriorityQueue[T, P]) Insert(item T, priority P) {
	q.items = append(q.items, Item[T, P]{value: item, priority: priority})
	q.lookup[item] = len(q.items) - 1
	q.siftUp(len(q.items) - 1)
}

func (q *PriorityQueue[T, P]) PopMin() (value T, priority P) {
	item := q.items[0]
	q.RemoveByIndex(0)
	return item.value, item.priority
}

func (q *PriorityQueue[T, P]) RemoveByIndex(i int) {
	item := q.items[i]
	n := q.Len()
	q.Swap(i, n-1)

	q.items = q.items[:n-1]
	delete(q.lookup, item.value)

	if i < n-1 {
		q.siftDown(i)
	}
}

func (q *PriorityQueue[T, P]) Lookup(item T) (index int, priority P) {
	must.Equal(len(q.items), len(q.lookup))
	index, ok := q.lookup[item]
	if !ok {
		return -1, 0
	}
	return index, q.items[index].priority
}

func (q *PriorityQueue[T, P]) Has(item T) bool {
	_, has := q.lookup[item]
	return has
}

func (q *PriorityQueue[T, P]) Inc(item T, inc P) {
	idx, p := q.Lookup(item)
	q.SetByIndex(idx, p+inc)
}

func (q *PriorityQueue[T, P]) SetByIndex(index int, priority P) {
	q.items[index].priority = priority
	q.siftUp(index)
	q.siftDown(index)
}

func (q *PriorityQueue[T, P]) siftUp(i int) {
	for q.items[i].priority < q.items[parent(i)].priority {
		q.Swap(i, parent(i))
		i = parent(i)
	}
}

func (q *PriorityQueue[T, P]) siftDown(i int) {
	for leftChild(i) < q.Len() {
		left := leftChild(i)
		right := rightChild(i)
		minChild := left
		if right < q.Len() && q.items[right].priority < q.items[left].priority {
			minChild = right
		}
		if q.items[i].priority <= q.items[minChild].priority {
			break
		}
		q.Swap(i, minChild)
		i = minChild
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return i<<1 + 1
}

func rightChild(i int) int {
	return i<<1 + 2
}
