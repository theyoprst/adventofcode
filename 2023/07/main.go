package main

import (
	"context"
	"slices"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

type Hand struct {
	key []int
	bid int
}

func SolvePart1(_ context.Context, lines []string) any {
	cardOrder := "23456789TJQKA"
	var hands []Hand
	for _, line := range lines {
		hand, bidStr := must.Split2(line, " ")
		bid := must.Atoi(bidStr)

		m := map[byte]int{}
		for i := range hand {
			m[hand[i]]++
		}
		key := aoc.MapSortedValues(m)
		slices.Reverse(key)

		for _, h := range hand {
			key = append(key, strings.Index(cardOrder, string(h)))
		}

		hands = append(hands, Hand{
			bid: bid,
			key: key,
		})
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		return slices.Compare(a.key, b.key)
	})
	ans := 0
	for i, hand := range hands {
		ans += (i + 1) * hand.bid
	}

	return ans
}

func SolvePart2(_ context.Context, lines []string) any {
	cardOrder := "J23456789TQKA"
	var hands []Hand
	for _, line := range lines {
		hand, bidStr := must.Split2(line, " ")
		bid := must.Atoi(bidStr)

		m := map[byte]int{}
		for i := range hand {
			m[hand[i]]++
		}
		j := m['J']
		delete(m, 'J')
		key := aoc.MapSortedValues(m)
		slices.Reverse(key)
		if len(key) > 0 {
			key[0] += j
		} else {
			key = []int{j}
		}

		for _, h := range hand {
			key = append(key, strings.Index(cardOrder, string(h)))
		}

		hands = append(hands, Hand{
			bid: bid,
			key: key,
		})
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		return slices.Compare(a.key, b.key)
	})
	ans := 0
	for i, hand := range hands {
		ans += (i + 1) * hand.bid
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
