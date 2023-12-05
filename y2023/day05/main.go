// 6:00, 6:32, 7:12
package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/theyoprst/adventofcode/helpers"
	"github.com/theyoprst/adventofcode/must"
)

func splitByEmpty(lines []string) [][]string {
	var g []string
	var gg [][]string
	for _, line := range append(lines, "") {
		if line == "" {
			gg = append(gg, g)
			g = []string{}
		} else {
			g = append(g, line)
		}
	}
	return gg
}

type MapItem struct {
	dst, src, size int
}

type Seg struct {
	start, after int
}

func main() {
	lines := helpers.ReadInputLines()
	_, seedsStr := must.Split2(lines[0], ": ")
	seeds := must.ParseInts(seedsStr)
	must.Equal(len(seeds)%2, 0)
	segs := make([]Seg, len(seeds)/2)
	for i := range segs {
		segs[i] = Seg{seeds[2*i], seeds[2*i] + seeds[2*i+1]}
	}
	lines = lines[2:]
	for _, g := range splitByEmpty(lines) {
		var items []MapItem
		for _, line := range g[1:] {
			ints := must.ParseInts(line)
			items = append(items, MapItem{dst: ints[0], src: ints[1], size: ints[2]})
		}
		for i, seed := range seeds {
			for _, item := range items {
				if item.src <= seed && seed < item.src+item.size {
					seed += item.dst - item.src
					break
				}
			}
			seeds[i] = seed
		}
		slices.SortFunc(items, func(a, b MapItem) int {
			return a.src - b.src
		})
		var newSegs []Seg
		for _, seg := range segs {
			var split []Seg
			for _, item := range items {
				in := Seg{
					start: max(seg.start, item.src),
					after: min(seg.after, item.src+item.size),
				}
				if in.start < in.after {
					// seg.start .. in.start .. in.after .. seg.after.
					if seg.start < in.start {
						split = append(split, Seg{seg.start, in.start})
					}
					mapped := in
					mapped.start += item.dst - item.src
					mapped.after += item.dst - item.src
					split = append(split, mapped)
					seg = Seg{in.after, seg.after}
				}
			}
			if seg.start < seg.after {
				split = append(split, seg)
			}
			newSegs = append(newSegs, split...)
		}
		segs = newSegs
	}
	ans1 := math.MaxInt
	for _, seed := range seeds {
		ans1 = min(ans1, seed)
	}
	ans2 := math.MaxInt
	for _, seg := range segs {
		ans2 = min(ans2, seg.start)
	}

	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}
