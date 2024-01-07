package main

import (
	"log"
	"math"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

type MapItem struct {
	dst, src, size int
}

type Seg struct {
	first, after int
}

func ParseInput(lines []string) (seeds []int, maps [][]MapItem) {
	seeds = aoc.Ints(lines[0])
	must.Equal(len(seeds)%2, 0)
	segs := make([]Seg, len(seeds)/2)
	for i := range segs {
		segs[i] = Seg{seeds[2*i], seeds[2*i] + seeds[2*i+1]}
	}
	lines = lines[2:]
	for _, g := range aoc.Split(lines, "") {
		var items []MapItem
		for _, line := range g[1:] {
			ints := aoc.Ints(line)
			items = append(items, MapItem{dst: ints[0], src: ints[1], size: ints[2]})
		}
		maps = append(maps, items)
	}
	return seeds, maps
}

func SolvePart1(lines []string) any {
	seeds, maps := ParseInput(lines)
	for _, items := range maps {
		for i, seed := range seeds {
			for _, item := range items {
				if item.src <= seed && seed < item.src+item.size {
					seed += item.dst - item.src
					break
				}
			}
			seeds[i] = seed
		}
	}
	return slices.Min(seeds)
}

// Start with 10 segments. Filter each segment through mappings splitting by subsegments.
// Happily, only 170 subsegments are at the end. But in theory there could be huge number of them.
func SolvePart2(lines []string) any {
	seeds, maps := ParseInput(lines)
	segs := make([]Seg, len(seeds)/2)
	for i := range segs {
		segs[i] = Seg{seeds[2*i], seeds[2*i] + seeds[2*i+1]}
	}
	for _, items := range maps {
		slices.SortFunc(items, func(a, b MapItem) int {
			return a.src - b.src
		})
		var newSegs []Seg
		for _, seg := range segs {
			var split []Seg
			for _, item := range items {
				in := Seg{ // Intersection of current segment and the mapping segment.
					first: max(seg.first, item.src),
					after: min(seg.after, item.src+item.size),
				}
				if in.first < in.after { // If intersection is not empty.
					// seg.first .. in.first .. in.after .. seg.after.
					if seg.first < in.first {
						// Prefix of current segment not covered by the mapping segment.
						split = append(split, Seg{seg.first, in.first})
					}
					mapped := in
					mapped.first += item.dst - item.src
					mapped.after += item.dst - item.src
					split = append(split, mapped)  // "Middle" of current segment mapped to the new place.
					seg = Seg{in.after, seg.after} // Contract current segment, the suffix will be processed by next mapping segments.
				}
			}
			if seg.first < seg.after { // If some current segment's suffix left, don't forget it.
				split = append(split, seg)
			}
			newSegs = append(newSegs, split...)
		}
		segs = newSegs
	}
	ans := math.MaxInt
	log.Printf("Final count of segments is %d", len(segs))
	for _, seg := range segs {
		ans = min(ans, seg.first)
	}

	return ans
}

func SolvePart2BruteForce(lines []string) any {
	seeds, maps := ParseInput(lines)
	segs := make([]Seg, len(seeds)/2)
	for i := range segs {
		segs[i] = Seg{seeds[2*i], seeds[2*i] + seeds[2*i+1]}
	}
	mins := make(chan int)
	for _, seg := range segs {
		seg := seg
		go func() {
			minSeed := math.MaxInt
			for i := seg.first; i < seg.after; i++ {
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
	ans := math.MaxInt
	for range segs {
		ans = min(ans, <-mins)
	}

	return ans
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2 /*, SolvePart2BruteForce*/}
)

func main() {
	log.SetFlags(0)
	aoc.Main(solvers1, solvers2)
}
