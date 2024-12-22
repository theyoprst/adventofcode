package main

import (
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

const steps = 2000

func SolvePart1(lines []string) any {
	sum := 0
	for _, line := range lines {
		num := must.Atoi(line)
		x := num
		for range steps {
			x = nextSecret(x)
		}
		sum += x
	}
	return sum
}

func SolvePart2(lines []string) any {
	// Idea: for each buyer, precalculate the price per each encountered for this buyer tuple of 4 diffs.
	// Store it in a map[diff -> price].
	// Then just summate all these maps. And find the max price in the resulting map.
	// Time Complexity: O(len(buyers) * len(steps) * len(diff)) ~= 2037 * 2000 * 4.
	// Space Complexity: O(20^len(diffs)) ~= 160_000 (in reality <41_000, because not all diffs are possible).

	// `totalPrice[tuple] = price` means buyers buy infromation after the `tuple` for `price` in total.
	totalPrice := make(map[int]int, 160000)
	hashKey := func(diffs []int) int {
		key := 0
		for _, diff := range diffs {
			// possible diffs are -9..9, shift by 9 to make them 0..18.
			key = key*32 + diff + 9
		}
		return key
	}

	for _, line := range lines {
		num := must.Atoi(line)
		var diffs []int

		// Have to track already seen tuples for each buyer because the monkey stops on the first seen tuple.
		seenByTheBuyer := make(map[int]struct{}, steps)

		x := num
		pricePrev := x % 10
		for range steps {
			x = nextSecret(x)
			price := x % 10
			if len(diffs) == 4 {
				diffs = slices.Delete(diffs, 0, 1)
			}
			diffs = append(diffs, price-pricePrev)
			if len(diffs) == 4 {
				key := hashKey(diffs)
				if _, seen := seenByTheBuyer[key]; !seen {
					seenByTheBuyer[key] = struct{}{}
					totalPrice[key] += price
				}
			}
			pricePrev = price
		}
	}

	maxPrice := 0
	for _, price := range totalPrice {
		maxPrice = max(maxPrice, price)
	}
	return maxPrice
}

func nextSecret(x int) int {
	const prune = 1<<24 - 1 // 0xFFFFFF
	x = (x<<6 ^ x) & prune
	x = (x>>5 ^ x) & prune
	x = (x<<11 ^ x) & prune
	return x
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
