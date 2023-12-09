// Time so solve: 28min.
package main

import (
	"fmt"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func Pop[T any](a []T) ([]T, T) {
	must.Greater(len(a), 0)
	return a[:len(a)-1], a[len(a)-1]
}

func main() {
	lines := aoc.ReadInputLines()
	emptyIdx := slices.Index(lines, "")
	n := len(aoc.Ints(lines[emptyIdx-1]))
	stacks := make([][]byte, n+1)
	stacks2 := make([][]byte, n+1)
	for i := emptyIdx - 2; i >= 0; i-- {
		line := lines[i]
		for j := 1; j <= n; j++ {
			letter := line[1+(j-1)*4]
			if letter != ' ' {
				stacks[j] = append(stacks[j], letter)
				stacks2[j] = append(stacks2[j], letter)
			}
		}
	}
	for i := emptyIdx + 1; i < len(lines); i++ {
		nums := aoc.Ints(lines[i])
		count, src, dst := nums[0], nums[1], nums[2]
		for j := 0; j < count; j++ {
			var x byte
			stacks[src], x = Pop(stacks[src])
			stacks[dst] = append(stacks[dst], x)
		}
		stacks2[dst] = append(stacks2[dst], stacks2[src][len(stacks2[src])-count:]...)
		stacks2[src] = stacks2[src][:len(stacks2[src])-count]
	}

	var ans []byte
	for i := 1; i <= n; i++ {
		ans = append(ans, stacks[i][len(stacks[i])-1])
	}
	fmt.Println("Part 1:", string(ans))

	var ans2 []byte
	for i := 1; i <= n; i++ {
		ans2 = append(ans2, stacks2[i][len(stacks2[i])-1])
	}
	fmt.Println("Part 2:", string(ans2))
}
