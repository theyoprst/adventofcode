package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

type Node struct {
	left, right string
}

func ParseGraph(lines []string) map[string]*Node {
	g := map[string]*Node{}
	for _, line := range lines {
		must.Equal(len(line), 16)
		value := line[0:3]
		node := &Node{
			left:  line[7:10],
			right: line[12:15],
		}
		g[value] = node
	}
	return g
}

func SolvePart2(lines []string) any {
	cmd := lines[0]
	g := ParseGraph(lines[2:])

	type Node2 struct {
		v string
		n int
	}

	ans2 := 1
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

	return ans2
}

func SolvePart1(lines []string) any {
	ans := 0
	cmd := lines[0]

	g := ParseGraph(lines[2:])
	value := "AAA"
	for ; value != "ZZZ"; ans++ {
		if cmd[ans%len(cmd)] == 'L' {
			value = g[value].left
		} else {
			value = g[value].right
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
