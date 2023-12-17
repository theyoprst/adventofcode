package queues

import (
	"slices"
	"testing"
)

func TestPQWorkflow(t *testing.T) {
	q := NewPriorityQueue[string, int]()
	index, _ := q.Lookup("")
	if index != -1 {
		t.Errorf("Unexpected index %d found on empty priority queue, want -1", index)
	}
	q.Insert("A", 10)
	q.Insert("B", 9)
	q.Insert("C", 8)
	q.Insert("D", 7)
	q.Insert("E", 6)

	item, priority := q.PopMin()
	if item != "E" {
		t.Errorf("Unexpected item: got %q, want %q", item, "E")
	}
	if priority != 6 {
		t.Errorf("Unexpected priority: got %d, want %d", priority, 6)
	}

	bIndex, bPriority := q.Lookup("B")
	if bPriority != 9 {
		t.Errorf("Unexpected priority: got %d, want %d", bPriority, 9)
	}
	q.SetByIndex(bIndex, 0)

	item, priority = q.PopMin()
	if item != "B" {
		t.Errorf("Unexpected item: got %q, want %q", item, "B")
	}
	if priority != 0 {
		t.Errorf("Unexpected priority: got %d, want %d", priority, 0)
	}

	var items []string
	var priorities []int
	for q.Len() > 0 {
		item, priority := q.PopMin()
		items = append(items, item)
		priorities = append(priorities, priority)
	}
	wantItems := []string{"D", "C", "A"}
	wantPriorities := []int{7, 8, 10}
	if slices.Compare(items, wantItems) != 0 {
		t.Errorf("Unexpected fetched items: got %v, want %v", items, wantItems)
	}
	if slices.Compare(priorities, wantPriorities) != 0 {
		t.Errorf("Unexpected fetched priorities: got %v, want %v", priorities, wantPriorities)
	}
}
