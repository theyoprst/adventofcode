// Time so solve: 28min.
package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(lines []string) any {
	blocks := aoc.Split(lines, "")
	must.Equal(len(blocks), 2)
	stacks := parseStacks(blocks[0])
	steps := blocks[1]
	for _, step := range steps {
		nums := aoc.Ints(step)
		count, src, dst := nums[0], nums[1], nums[2]
		for j := 0; j < count; j++ {
			var x byte
			stacks[src], x = pop(stacks[src])
			stacks[dst] = append(stacks[dst], x)
		}
	}

	var ans []byte
	for i := 1; i < len(stacks); i++ {
		ans = append(ans, stacks[i][len(stacks[i])-1])
	}

	return string(ans)
}

func SolvePart2(lines []string) any {
	blocks := aoc.Split(lines, "")
	must.Equal(len(blocks), 2)
	stacks := parseStacks(blocks[0])
	steps := blocks[1]
	for _, step := range steps {
		nums := aoc.Ints(step)
		count, src, dst := nums[0], nums[1], nums[2]
		stacks[dst] = append(stacks[dst], stacks[src][len(stacks[src])-count:]...)
		stacks[src] = stacks[src][:len(stacks[src])-count]
	}

	var ans []byte
	for i := 1; i < len(stacks); i++ {
		ans = append(ans, stacks[i][len(stacks[i])-1])
	}
	return string(ans)
}

func parseStacks(lines []string) [][]byte {
	n := len(aoc.Ints(lines[len(lines)-1]))
	stacks := make([][]byte, n+1)
	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i]
		for j := 1; j <= n; j++ {
			letter := line[1+(j-1)*4]
			if letter != ' ' {
				stacks[j] = append(stacks[j], letter)
			}
		}
	}
	return stacks
}

func pop[T any](a []T) ([]T, T) {
	must.Greater(len(a), 0)
	return a[:len(a)-1], a[len(a)-1]
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
