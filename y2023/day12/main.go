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

func SolveGeneric(lines []string, dup int) int {
	var ans int
	for _, line := range lines {
		maskStr, numsStr := must.Split2(line, " ")
		maskStr = strings.Join(aoc.MakeSlice(maskStr, dup), "?") + "."
		numsStr = strings.Join(aoc.MakeSlice(numsStr, dup), ",")
		numsStrs := strings.Split(numsStr, ",")
		var nums []int
		for _, s := range numsStrs {
			nums = append(nums, must.Atoi(s))
		}
		mask := []byte(maskStr)
		numsSum := 0
		for _, n := range nums {
			numsSum += n
		}
		nDots := 0
		nBroken := 0
		for _, b := range mask {
			nDots += aoc.BoolToInt(b == '.')
			nBroken += aoc.BoolToInt(b == '#')
		}
		type CacheItem struct{ curN, maskIdx, numsIdx int }
		cache := map[CacheItem]int{}
		var bf func(curN int, maskIdx int, numsIds int) int
		bf = func(curN int, maskIdx int, numsIdx int) (result int) {
			cacheItem := CacheItem{curN, maskIdx, numsIdx}
			if result, ok := cache[cacheItem]; ok {
				return result
			}
			defer func() {
				cache[cacheItem] = result
			}()
			nums := nums[numsIdx:]
			mask := mask[maskIdx:]
			if len(mask) == 0 {
				if len(nums) == 0 {
					return 1
				}
				return 0
			}
			if curN > 0 {
				if len(nums) == 0 {
					return 0
				}
				if len(nums) > 0 && curN > nums[0] {
					return 0
				}
			}
			ch := mask[0]
			sum := 0
			if ch == '#' || ch == '?' {
				sum += bf(curN+1, maskIdx+1, numsIdx)
			}
			if ch == '.' || ch == '?' {
				if curN > 0 {
					if len(nums) > 0 && nums[0] == curN {
						sum += bf(0, maskIdx+1, numsIdx+1)
					}
				} else {
					sum += bf(0, maskIdx+1, numsIdx)
				}
			}
			return sum
		}
		arrangements := bf(0, 0, 0)
		ans += arrangements
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
