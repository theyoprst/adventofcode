package main

import (
	"fmt"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func main() {
	ans1, ans2 := 0, 0
	lines := aoc.ReadInputLines()
	cmd := lines[0]
	_ = cmd

	type Node struct {
		left, right string
	}
	g := map[string]*Node{}
	for _, line := range lines[2:] {
		must.Equal(len(line), 16)
		value := line[0:3]
		node := &Node{
			left:  line[7:10],
			right: line[12:15],
		}
		g[value] = node
	}
	if g["AAA"] != nil {
		value := "AAA"
		for ; value != "ZZZ"; ans1++ {
			if cmd[ans1%len(cmd)] == 'L' {
				value = g[value].left
			} else {
				value = g[value].right
			}
		}
		fmt.Println("Part 1:", ans1)
	}

	// Brute force.
	set := map[string]bool{}
	for value := range g {
		if strings.HasSuffix(value, "A") {
			set[value] = true
		}
	}
	isFinish := func(set map[string]bool) bool {
		for value := range set {
			if !strings.HasSuffix(value, "Z") {
				return false
			}
		}
		return true
	}
	newSet := map[string]bool{}
	for ; !isFinish(set) && ans2 <= len(cmd)*len(g); ans2++ {
		clear(newSet)
		isLeft := cmd[ans2%len(cmd)] == 'L'
		for value := range set {
			var newValue string
			if isLeft {
				newValue = g[value].left
			} else {
				newValue = g[value].right
			}
			newSet[newValue] = true
		}
		set, newSet = newSet, set
	}
	if isFinish(set) {
		fmt.Println("Part 2 (brute force):", ans2)
	} else {
		fmt.Println("Brute force lasts too long, need for find cycles.")
	}

	type Node2 struct {
		v string
		n int
	}

	ans2 = 1
	for v := range g {
		if v[2] != 'A' {
			continue
		}
		seen := map[Node2]int{}

		var loopLen int
		for n := 0; ; n++ {
			cn := n % len(cmd)
			if v[2] == 'Z' {
				if seen[Node2{v, cn}] > 0 {
					loopLen = n - seen[Node2{v, cn}]
					must.Equal(loopLen%len(cmd), 0)
					break
				}
				seen[Node2{v, cn}] = n
			}
			if cmd[n%len(cmd)] == 'L' {
				v = g[v].left
			} else {
				v = g[v].right
			}
		}
		// We ignore shifts (Node2.n) here because they are equal to loopLen in test data.
		// In general case it is more complex.
		ans2 = aoc.LCM(ans2, loopLen)
	}

	fmt.Println("Part 2:", ans2)
}
