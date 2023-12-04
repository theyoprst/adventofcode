// 6:00 - 6:12 - 6:40
package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/theyoprst/adventofcode/helpers"
	"github.com/theyoprst/adventofcode/must"
)

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
		wins := must.ParseInts(strings.TrimSpace(winsStr))
		have := must.ParseInts(strings.TrimSpace(haveStr))
		k := len(intersect(wins, have))
		copies[i]++
		ans2 += copies[i]
		if k > 0 {
			ans1 += 1 << (k - 1)
			for j := 0; j < k; j++ {
				copies[i+j+1] += copies[i]
			}
		}
	}
	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}
