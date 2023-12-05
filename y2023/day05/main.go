// 6:00, 6:32, 7:12
package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

type MapItem struct {
	dst, src, size int
}

type Seg struct {
	start, after int
}

func main() {
	lines := aoc.ReadInputLines()
	seeds := aoc.Ints(lines[0])
	must.Equal(len(seeds)%2, 0)
	segs := make([]Seg, len(seeds)/2)
	for i := range segs {
		segs[i] = Seg{seeds[2*i], seeds[2*i] + seeds[2*i+1]}
	}
	origSegs := slices.Clone(segs)
	lines = lines[2:]
	var maps [][]MapItem
	for _, g := range aoc.Split(lines, "") {
		var items []MapItem
		for _, line := range g[1:] {
			ints := aoc.Ints(line)
			items = append(items, MapItem{dst: ints[0], src: ints[1], size: ints[2]})
		}
		maps = append(maps, items)
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
	ans1 := slices.Min(seeds)
	ans2 := math.MaxInt
	for _, seg := range segs {
		ans2 = min(ans2, seg.start)
	}

	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)

	// Brute-force:
	mins := make(chan int)
	for _, seg := range origSegs {
		seg := seg
		go func() {
			minSeed := math.MaxInt
			for i := seg.start; i < seg.after; i++ {
				seed := i
				for _, items := range maps {
					for _, item := range items {
						if item.src <= seed && seed < item.src+item.size {
							seed += item.dst - item.src
							break
						}
					}
				}
				minSeed = min(minSeed, seed)
			}
			mins <- minSeed
		}()
	}
	ans2bf := math.MaxInt
	for range origSegs {
		ans2bf = min(ans2bf, <-mins)
	}

	fmt.Println("Part 2 (brute force):", ans2bf)
}
