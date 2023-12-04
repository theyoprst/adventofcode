// 6:00 - 6:12 - 6:40
package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/theyoprst/adventofcode/helpers"
	"github.com/theyoprst/adventofcode/must"
)

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
		r = r >> 1
		l = l >> 1
	}
}

func (t *STree) Get(i int) int {
	res := 0
	for i = i + t.n; i > 0; i = i >> 1 {
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

func main() {
	var ans1, ans2, ans2V2 int
	lines := helpers.ReadInputLines()
	copies := map[int]int{}
	tree := NewSTree(len(lines))
	tree.Inc(0, len(lines), 1)
	for i, line := range lines {
		_, line = must.Split2(line, ":")
		winsStr, haveStr := must.Split2(line, "|")
		wins := must.ParseInts(strings.TrimSpace(winsStr))
		have := must.ParseInts(strings.TrimSpace(haveStr))
		k := len(intersect(wins, have))
		copies[i]++
		ans2 += copies[i]
		ans2V2 += tree.Get(i)
		if k > 0 {
			ans1 += 1 << (k - 1)
		}
		for j := 0; j < k; j++ {
			copies[i+j+1] += copies[i]
		}
		tree.Inc(i+1, min(i+k+1, len(lines)), tree.Get(i))
	}
	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
	fmt.Println("Part 2 v2:", ans2V2)
}
