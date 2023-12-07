package main

import (
	"fmt"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

type Hand struct {
	hand string
	typ  []int
	jtyp []int
	bid  int
}

func NewHand(hand string, bid int) Hand {
	m := map[byte]int{}
	for i := range hand {
		m[hand[i]]++
	}
	var a []int
	for _, v := range m {
		a = append(a, v)
	}
	slices.Sort(a)
	slices.Reverse(a)

	aj := a
	j, ok := m['J']
	if ok {
		delete(m, 'J')
		aj = []int{}
		for _, v := range m {
			aj = append(aj, v)
		}
		slices.Sort(aj)
		slices.Reverse(aj)
		if len(aj) > 0 {
			aj[0] += j
		} else {
			aj = []int{j}
		}
	}
	return Hand{
		hand: hand,
		typ:  a,
		jtyp: aj,
		bid:  bid,
	}
}

var cardOrder = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func main() {
	ans1, ans2 := 0, 0
	lines := aoc.ReadInputLines()
	var hands []Hand
	for i, line := range lines {
		_, _ = i, line
		hand, bidStr := must.Split2(line, " ")
		bid := must.Atoi(bidStr)
		hands = append(hands, NewHand(hand, bid))
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		res := slices.Compare(a.typ, b.typ)
		if res != 0 {
			return res
		}
		for i := range a.hand {
			ca := cardOrder[a.hand[i]]
			cb := cardOrder[b.hand[i]]
			if ca != cb {
				return ca - cb
			}
		}
		return 0
	})
	for i, hand := range hands {
		ans1 += (i + 1) * hand.bid
	}

	cardOrder['J'] = -1
	slices.SortFunc(hands, func(a, b Hand) int {
		res := slices.Compare(a.jtyp, b.jtyp)
		if res != 0 {
			return res
		}
		for i := range a.hand {
			ca := cardOrder[a.hand[i]]
			cb := cardOrder[b.hand[i]]
			if ca != cb {
				return ca - cb
			}
		}
		return 0
	})
	for i, hand := range hands {
		ans2 += (i + 1) * hand.bid
	}

	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}
