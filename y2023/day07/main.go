package main

import (
	"fmt"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func main() {
	lines := aoc.ReadInputLines()

	cardOrder := map[rune]int{
		'T': 11,
		'J': 12,
		'Q': 13,
		'K': 14,
		'A': 15,
	}
	for ch := '2'; ch <= '9'; ch++ {
		cardOrder[ch] = int(ch - '0')
	}

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

		for i, h := range hand {
			x := cardOrder[h]
			key = append(key, x)
			if hand[i] == 'J' {
				x = 0
			}
			keyJ = append(keyJ, x)
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
