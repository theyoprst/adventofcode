package main

import (
	"slices"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func hash(s string) int {
	h := 0
	for _, ch := range s {
		h = (h + int(ch)) * 17 % 256
	}
	return h
}

func SolvePart1(lines []string) any {
	strs := strings.Split(lines[0], ",")
	ans := 0
	for _, s := range strs {
		ans += hash(s)
	}
	return ans
}

func SolvePart2(lines []string) any {
	strs := strings.Split(lines[0], ",")
	type Item struct {
		Label    string
		FocalLen int
	}
	var hashmap [256][]Item
	index := func(h int, label string) int {
		return slices.IndexFunc(hashmap[h], func(item Item) bool {
			return item.Label == label
		})
	}
	for _, s := range strs {
		if strings.HasSuffix(s, "-") {
			// Delete
			label := s[:len(s)-1]
			h := hash(label)
			if i := index(h, label); i != -1 {
				hashmap[h] = slices.Delete(hashmap[h], i, i+1)
			}
		} else {
			// Insert
			label, focalStr := must.Split2(s, "=")
			focal := must.Atoi(focalStr)
			h := hash(label)
			if i := index(h, label); i != -1 {
				hashmap[h][i].FocalLen = focal
			} else {
				hashmap[h] = append(hashmap[h], Item{label, focal})
			}
		}
	}
	ans := 0
	for boxI, box := range hashmap {
		for itemI, item := range box {
			ans += (boxI + 1) * (itemI + 1) * item.FocalLen
		}
	}
	return ans
}

var (
	solvers1 []aoc.Solver = []aoc.Solver{SolvePart1}
	solvers2 []aoc.Solver = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
