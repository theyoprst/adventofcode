package main

import (
	"context"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

const space = -1

func SolvePart1(_ context.Context, lines []string) any {
	must.Equal(len(lines), 1)
	text := lines[0]
	var blocks []int
	for i, ch := range text {
		num := int(ch - '0')
		if i%2 == 0 {
			blocks = append(blocks, slices.Repeat([]int{i / 2}, num)...)
		} else {
			blocks = append(blocks, slices.Repeat([]int{space}, num)...)
		}
	}
	left := 0
	right := len(blocks) - 1
	for left < right {
		if blocks[left] != space {
			left++
			continue
		}
		if blocks[right] == space {
			right--
			continue
		}
		blocks[left], blocks[right] = blocks[right], blocks[left]
		left++
		right--
	}
	sum := 0
	for i, b := range blocks {
		if b != space {
			sum += i * b
		}
	}

	return sum
}

func SolvePart2(_ context.Context, lines []string) any {
	must.Equal(len(lines), 1)
	text := lines[0]
	var blocks []int
	for i, ch := range text {
		num := int(ch - '0')
		if i%2 == 0 {
			blocks = append(blocks, slices.Repeat([]int{i / 2}, num)...)
		} else {
			blocks = append(blocks, slices.Repeat([]int{space}, num)...)
		}
	}

	holeSizeToIdx := map[int]int{}
	findHole := func(blocks []int, size int) int {
		for i := holeSizeToIdx[size]; i < len(blocks)-size+1; i++ {
			if blocks[i] == space {
				ok := true
				for j := 1; j < size; j++ {
					if blocks[i+j] != space {
						ok = false
						break
					}
				}
				if ok {
					holeSizeToIdx[size] = i
					return i
				}
			}
		}
		holeSizeToIdx[size] = 1000000000
		return -1
	}

	rfirst := len(blocks)
	for rfirst > 0 {
		rlast := rfirst - 1
		for rlast >= 0 && blocks[rlast] == space {
			rlast--
		}
		if rlast < 0 {
			break
		}
		rfirst = rlast
		for rfirst >= 0 && blocks[rfirst] == blocks[rlast] {
			rfirst--
		}
		rfirst++
		if idx := findHole(blocks[:rfirst], rlast-rfirst+1); idx != -1 {
			src, dst := rfirst, idx
			for src < rlast+1 {
				blocks[src], blocks[dst] = blocks[dst], blocks[src]
				src++
				dst++
			}
		}
	}

	sum := 0
	for i, b := range blocks {
		if b != space {
			sum += i * b
		}
	}

	return sum
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
