package main

import (
	"context"
	"log"
	"slices"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(_ context.Context, lines []string) any {
	var ans int
	for _, line := range lines {
		_, line = must.Split2(line, ":")
		winsStr, haveStr := must.Split2(line, "|")
		wins := aoc.Ints(strings.TrimSpace(winsStr))
		have := aoc.Ints(strings.TrimSpace(haveStr))
		k := len(intersect(wins, have))
		if k > 0 {
			ans += 1 << (k - 1)
		}
	}
	return ans
}

func SolvePart2(_ context.Context, lines []string) any {
	var ans int
	copies := map[int]int{}
	for i, line := range lines {
		_, line = must.Split2(line, ":")
		winsStr, haveStr := must.Split2(line, "|")
		wins := aoc.Ints(strings.TrimSpace(winsStr))
		have := aoc.Ints(strings.TrimSpace(haveStr))
		k := len(intersect(wins, have))
		copies[i]++
		ans += copies[i]
		for j := 0; j < k; j++ {
			copies[i+j+1] += copies[i]
		}
	}
	return ans
}

func SolvePart2SegmentTree(_ context.Context, lines []string) any {
	var ans int
	tree := NewSTree(len(lines))
	tree.Inc(0, len(lines), 1)
	for i, line := range lines {
		_, line = must.Split2(line, ":")
		winsStr, haveStr := must.Split2(line, "|")
		wins := aoc.Ints(strings.TrimSpace(winsStr))
		have := aoc.Ints(strings.TrimSpace(haveStr))
		k := len(intersect(wins, have))
		ans += tree.Get(i)
		tree.Inc(i+1, min(i+k+1, len(lines)), tree.Get(i))
	}
	return ans
}

type STree struct {
	n int
	t []int
}

func NewSTree(n int) STree {
	return STree{
		n: n,
		t: make([]int, 2*n),
	}
}

func (t *STree) Inc(l, r int, value int) {
	l += t.n
	r += t.n
	for l < r {
		if l&1 != 0 {
			t.t[l] += value
			l++
		}
		if r&1 != 0 {
			r--
			t.t[r] += value
		}
		r >>= 1
		l >>= 1
	}
}

func (t *STree) Get(i int) int {
	res := 0
	for i += t.n; i > 0; i >>= 1 {
		res += t.t[i]
	}
	return res
}

func intersect(a, b []int) []int {
	var res []int
	for _, x := range a {
		if slices.Contains(b, x) {
			res = append(res, x)
		}
	}
	return res
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2, SolvePart2SegmentTree}
)

func main() {
	log.SetFlags(0)
	aoc.Main(solvers1, solvers2)
}
