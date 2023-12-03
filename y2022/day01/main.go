// 12:16 - 12:21 - 12:24
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/theyoprst/adventofcode/must"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var ans1, ans2 int
	var sum int
	var sums []int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			ans1 = max(ans1, sum)
			sums = append(sums, sum)
			sum = 0
		} else {
			sum += must.Atoi(line)
		}
	}
	ans1 = max(ans1, sum)
	sums = append(sums, sum)
	slices.Sort(sums)
	slices.Reverse(sums)
	for _, s := range sums[:3] {
		ans2 += s
	}
	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}
