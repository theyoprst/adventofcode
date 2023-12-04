// 6:12 - 6:40
package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/theyoprst/adventofcode/helpers"
	"github.com/theyoprst/adventofcode/must"
)

func mustParseList(s string) []int {
	split := strings.Fields(s)
	var res []int
	for _, x := range split {
		res = append(res, must.Atoi(x))
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
	var ans1, ans2 int
	lines := helpers.ReadInputLines()
	copies := map[int]int{}
	for i, line := range lines {
		_, line = must.Split2(line, ":")
		winsStr, haveStr := must.Split2(line, "|")
		wins := mustParseList(strings.TrimSpace(winsStr))
		have := mustParseList(strings.TrimSpace(haveStr))
		in := intersect(wins, have)
		if len(in) > 0 {
			ans1 += 1 << (len(in) - 1)
		}
		inc := copies[i] + 1
		ans2 += inc
		if len(in) > 0 {
			for j := 0; j < len(in); j++ {
				copies[i+j+1] += copies[i] + 1
			}
		}
	}
	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}
