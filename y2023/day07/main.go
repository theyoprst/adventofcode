package main

import (
	"fmt"
	"maps"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func main() {
	lines := aoc.ReadInputLines()

	cardOrder := map[rune]int{}
	for i, r := range "23456789TJQKA" {
		cardOrder[r] = i
	}
	cardOrderJ := maps.Clone(cardOrder)
	cardOrderJ['J'] = -1

	type Hand struct {
		key  []int
		keyJ []int
		bid  int
	}
	var hands []Hand
	for i, line := range lines {
		_, _ = i, line
		hand, bidStr := must.Split2(line, " ")
		bid := must.Atoi(bidStr)

		m := map[byte]int{}
		for i := range hand {
			m[hand[i]]++
		}
		key := aoc.MapSortedValues(m)
		slices.Reverse(key)

		j := m['J']
		delete(m, 'J')
		keyJ := aoc.MapSortedValues(m)
		slices.Reverse(keyJ)
		if len(keyJ) > 0 {
			keyJ[0] += j
		} else {
			keyJ = []int{j}
		}

		for _, h := range hand {
			key = append(key, cardOrder[h])
			keyJ = append(keyJ, cardOrderJ[h])
		}

		hands = append(hands, Hand{
			bid:  bid,
			key:  key,
			keyJ: keyJ,
		})
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		return slices.Compare(a.key, b.key)
	})
	ans1 := 0
	for i, hand := range hands {
		ans1 += (i + 1) * hand.bid
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		return slices.Compare(a.keyJ, b.keyJ)
	})
	ans2 := 0
	for i, hand := range hands {
		ans2 += (i + 1) * hand.bid
	}

	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}
