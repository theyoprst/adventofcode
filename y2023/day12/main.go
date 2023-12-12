package main

import (
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(lines []string) any {
	return SolveGeneric(lines, 1)
}

func SolvePart2(lines []string) any {
	return SolveGeneric(lines, 5)
}

const (
	Block = '#'
	Empty = '.'
	Any   = '?'
)

func SolveGeneric(lines []string, dup int) int {
	var ans int
	for _, line := range lines {
		maskStr, blocksStr := must.Split2(line, " ")
		maskStr = strings.Join(aoc.MakeSlice(maskStr, dup), "?") + string(Empty)
		blocksStr = strings.Join(aoc.MakeSlice(blocksStr, dup), ",")
		blocks := aoc.Ints(blocksStr)
		mask := []byte(maskStr)

		type CacheItem struct{ curBlock, maskIdx, blocksIdx int }
		cache := map[CacheItem]int{}

		var dp func(curN int, maskIdx int, blocksIdx int) int
		dp = func(curBlock int, maskIdx int, blocksIdx int) (result int) {
			cacheItem := CacheItem{curBlock, maskIdx, blocksIdx}
			if result, ok := cache[cacheItem]; ok {
				return result
			}
			defer func() {
				cache[cacheItem] = result
			}()

			blocks := blocks[blocksIdx:]
			mask := mask[maskIdx:]
			if len(mask) == 0 {
				return aoc.BoolToInt(len(blocks) == 0)
			}
			ch := mask[0]
			sum := 0
			if ch == Block || ch == Any {
				sum += dp(curBlock+1, maskIdx+1, blocksIdx)
			}
			if ch == Empty || ch == Any {
				if curBlock > 0 {
					if len(blocks) > 0 && blocks[0] == curBlock {
						sum += dp(0, maskIdx+1, blocksIdx+1)
					}
				} else {
					sum += dp(0, maskIdx+1, blocksIdx)
				}
			}
			return sum
		}
		ans += dp(0, 0, 0)
	}
	return ans
}

var solversPart1 []aoc.Solver = []aoc.Solver{
	SolvePart1,
}

var solversPart2 []aoc.Solver = []aoc.Solver{
	SolvePart2,
}

func main() {
	aoc.Main([]aoc.Solver{SolvePart1}, []aoc.Solver{SolvePart2})
}
