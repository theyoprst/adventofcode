package main

import (
	"context"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
)

type OrderRule struct {
	Before, After int
}

func newOrderRules(lines []string) containers.Set[OrderRule] {
	orders := make(containers.Set[OrderRule], len(lines))
	for _, rule := range lines {
		r := aoc.Ints(rule)
		orders.Add(OrderRule{Before: r[0], After: r[1]})
	}
	return orders
}

// orderComparator is used both for checking and imposing order rules.
// Happily, order rules define total order, not just a partial one. There are even n*(n-1)/2 rules for sequences of n pages.
// If there were less rules, or the order was just partial, more sophisticated algorithms should be used,
// like topsort and check for topsorted.
func orderComparator(orderRules containers.Set[OrderRule]) func(int, int) int {
	return func(pageA, pageB int) int {
		if orderRules.Has(OrderRule{Before: pageA, After: pageB}) {
			return -1
		}
		if orderRules.Has(OrderRule{Before: pageB, After: pageA}) {
			return 1
		}
		return 0
	}
}

func SolvePart1(_ context.Context, lines []string) any {
	blocks := aoc.Blocks(lines)
	rulesBlock, updatesBlock := blocks[0], blocks[1]
	orderRules := newOrderRules(rulesBlock)
	cmp := orderComparator(orderRules)

	ans := 0
	for _, update := range updatesBlock {
		pages := aoc.Ints(update)
		if slices.IsSortedFunc(pages, cmp) {
			ans += pages[len(pages)/2]
		}
	}
	return ans
}

func SolvePart2(_ context.Context, lines []string) any {
	blocks := aoc.Blocks(lines)
	rulesBlock, updatesBlock := blocks[0], blocks[1]
	orderRules := newOrderRules(rulesBlock)
	cmp := orderComparator(orderRules)

	ans := 0
	for _, update := range updatesBlock {
		pages := aoc.Ints(update)
		if !slices.IsSortedFunc(pages, cmp) {
			slices.SortFunc(pages, cmp)
			ans += pages[len(pages)/2]
		}
	}
	return ans
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
