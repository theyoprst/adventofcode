// 16:50 - 17:01 - 17:10.
package main

import (
	"context"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

type RuneSet map[rune]struct{}

func toSet(s string) RuneSet {
	set := map[rune]struct{}{}
	for _, r := range s {
		set[r] = struct{}{}
	}
	return set
}

func intersect(set1, set2 RuneSet) RuneSet {
	res := RuneSet{}
	for r := range set1 {
		if _, ok := set2[r]; ok {
			res[r] = struct{}{}
		}
	}
	return res
}

func priority(r rune) int {
	if 'a' <= r && r <= 'z' {
		return int(r - 'a' + 1)
	}
	must.GreaterOrEqual(r, 'A')
	must.LessOrEqual(r, 'Z')
	return int(r - 'A' + 27)
}

func SolvePart1(_ context.Context, lines []string) any {
	ans := 0
	for _, line := range lines {
		lines = append(lines, line)
		must.Equal(len(line)%2, 0)
		s1, s2 := line[:len(line)/2], line[len(line)/2:]
		m1, m2 := toSet(s1), toSet(s2)
		i := intersect(m1, m2)
		must.Equal(len(i), 1)
		for r := range i {
			ans += priority(r)
		}
	}
	return ans
}

func SolvePart2(_ context.Context, lines []string) any {
	ans := 0
	must.Equal(len(lines)%3, 0)
	for i := 0; i < len(lines)/3; i++ {
		ll := lines[i*3 : i*3+3]
		m1, m2, m3 := toSet(ll[0]), toSet(ll[1]), toSet(ll[2])
		i := intersect(intersect(m1, m2), m3)
		must.Equal(len(i), 1)
		for r := range i {
			ans += priority(r)
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
