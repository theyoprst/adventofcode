package main

import (
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

func isOrderCorrect(rules containers.Set[OrderRule], pages []int) bool {
	for i, left := range pages {
		for _, right := range pages[i+1:] {
			reversed := OrderRule{Before: right, After: left}
			if rules.Has(reversed) {
				return false
			}
		}
	}
	return true
}

func SolvePart1(lines []string) any {
	blocks := aoc.Split(lines, "")
	rulesBlock, updatesBlock := blocks[0], blocks[1]
	orderRules := newOrderRules(rulesBlock)

	ans := 0
	for _, update := range updatesBlock {
		pages := aoc.Ints(update)
		if isOrderCorrect(orderRules, pages) {
			ans += pages[len(pages)/2]
		}
	}
	return ans
}

func SolvePart2(lines []string) any {
	blocks := aoc.Split(lines, "")
	rulesBlock, updatesBlock := blocks[0], blocks[1]
	orderRules := newOrderRules(rulesBlock)

	ans := 0
	for _, update := range updatesBlock {
		pages := aoc.Ints(update)
		if isOrderCorrect(orderRules, pages) {
			continue
		}
		slices.SortFunc(pages, func(pageA, pageB int) int {
			if orderRules.Has(OrderRule{Before: pageA, After: pageB}) {
				return -1
			}
			if orderRules.Has(OrderRule{Before: pageB, After: pageA}) {
				return 1
			}
			return 0
		})
		ans += pages[len(pages)/2]
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
