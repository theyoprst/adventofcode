package main

import (
	"context"
	"math"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

const stepsCount = 2000

func SolvePart1(_ context.Context, lines []string) any {
	sum := 0
	for _, line := range lines {
		secret := must.Atoi(line)
		for range stepsCount {
			secret = nextSecret(secret)
		}
		sum += secret
	}
	return sum
}

func SolvePart2(_ context.Context, lines []string) any {
	const diffsCount = 4
	// Idea: for each buyer, precalculate the price per each encountered for this buyer tuple of `diffsCount` diffs.
	// Store it in a map[diff -> price].
	// Then just summate all these maps. And find the max price in the resulting map.
	// Time Complexity: O(buyersCount * stepsCount * diffsCount) ~= 2037 * 2000 * 4.
	//   The `diffsCount` term can be dropped if polinomial hash is used AND this hash is still inside a machine word (64 bits).
	//   If module 10 is used for price, it requires 5 bits per a diff, so 64/5 = 12 diffs can be hashed in a single 64-bit word.
	// Space Complexity: minimun of
	//   - O(19^diffsCount) ~= 160_000 (in reality <41_000, because not all diffs are possible).
	//   - O(buyersCount * stepsCount) ~= 2037 * 2000 (under the condition that hash is a single machine word).

	nextPolynomialHash := func(hash int, diff int) int {
		hash <<= 5                    // 5 bits per diff (-9..9)
		hash += diff + 9              // diff (-9..9) -> (0..18)
		hash &= 1<<(5*diffsCount) - 1 // keep only last `diffsCount` diffs, 5 bits per diff
		return hash
	}

	// `totalPrice[h] = p` means buyers will buy the information after the sequence of 4 diffs with hash `h` for price `p` in total.
	totalPrice := make(map[int]int, int(math.Pow(19, diffsCount)))

	for _, line := range lines {
		secret := must.Atoi(line)
		seenByTheBuyer := make(map[int]struct{}, stepsCount) // to ignore the same diffs for the same buyer
		pricePrev := secret % 10
		hash := 0
		for step := range stepsCount {
			secret = nextSecret(secret)
			price := secret % 10
			hash = nextPolynomialHash(hash, price-pricePrev)
			if step >= diffsCount-1 { // the polynomial hash is ready to be used: it has `diffsCount` diffs
				if _, seen := seenByTheBuyer[hash]; !seen {
					seenByTheBuyer[hash] = struct{}{}
					totalPrice[hash] += price
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
