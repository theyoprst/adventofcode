package main

import (
	"sort"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

type Monkey struct {
	items   []int
	op      string
	opValue int
	test    int
	ifTrue  int
	ifFalse int
	count   int
}

func parseMonkeys(lines []string) []*Monkey {
	monkeys := make([]*Monkey, 0)
	for i := 0; i < len(lines); i += 7 {
		items := strings.Split(strings.Split(lines[i+1], ": ")[1], ", ")
		itemsList := make([]int, len(items))
		for j, item := range items {
			itemsList[j] = must.Atoi(item)
		}
		operation := strings.Split(lines[i+2], "old ")[1]
		value := 0
		if strings.Contains(operation, "old") {
			value = 0
		} else {
			value = must.Atoi(strings.Split(operation, " ")[1])
		}
		test := must.Atoi(strings.Split(lines[i+3], "by ")[1])
		ifTrue := must.Atoi(strings.Split(lines[i+4], "monkey ")[1])
		ifFalse := must.Atoi(strings.Split(lines[i+5], "monkey ")[1])
		monkeys = append(monkeys, &Monkey{items: itemsList, op: operation, opValue: value, test: test, ifTrue: ifTrue, ifFalse: ifFalse})
	}
	return monkeys
}

func SolvePart1(lines []string) any {
	monkeys := parseMonkeys(lines)
	for round := 0; round < 20; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				var newItem int
				if monkey.opValue == 0 { // operation old + old or old * old.
					if strings.Contains(monkey.op, "+") {
						newItem = item + item
					} else if strings.Contains(monkey.op, "*") {
						newItem = item * item
					}
				} else {
					if strings.Contains(monkey.op, "+") {
						newItem = item + monkey.opValue
					} else if strings.Contains(monkey.op, "*") {
						newItem = item * monkey.opValue
					}
				}
				newItem /= 3
				if newItem%monkey.test == 0 {
					monkeys[monkey.ifTrue].items = append(monkeys[monkey.ifTrue].items, newItem)
				} else {
					monkeys[monkey.ifFalse].items = append(monkeys[monkey.ifFalse].items, newItem)
				}
				monkey.count++
			}
			monkey.items = make([]int, 0)
		}
	}

	counts := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		counts[i] = monkey.count
	}
	sort.Ints(counts)
	return counts[len(counts)-1] * counts[len(counts)-2]
}

func SolvePart2(lines []string) any {
	monkeys := parseMonkeys(lines)

	lcm := 1
	for _, monkey := range monkeys {
		lcm = aoc.LCM(lcm, monkey.test)
	}

	for round := 0; round < 10000; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				var newItem int
				if monkey.opValue == 0 {
					if strings.Contains(monkey.op, "+") {
						newItem = item + item
					} else if strings.Contains(monkey.op, "*") {
						newItem = item * item
					}
				} else {
					if strings.Contains(monkey.op, "+") {
						newItem = item + monkey.opValue
					} else if strings.Contains(monkey.op, "*") {
						newItem = item * monkey.opValue
					}
				}
				newItem %= lcm
				if newItem%monkey.test == 0 {
					monkeys[monkey.ifTrue].items = append(monkeys[monkey.ifTrue].items, newItem)
				} else {
					monkeys[monkey.ifFalse].items = append(monkeys[monkey.ifFalse].items, newItem)
				}
				monkey.count++
			}
			monkey.items = make([]int, 0)
		}
	}

	counts := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		counts[i] = monkey.count
	}
	sort.Ints(counts)
	return int64(counts[len(counts)-1]) * int64(counts[len(counts)-2])
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
