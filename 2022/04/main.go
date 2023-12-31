// 16:43 - 16:53 - 17:00
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/theyoprst/adventofcode/must"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var ans1, ans2 int
	for scanner.Scan() {
		line := scanner.Text()
		a, b := must.Split2(line, ",")
		a1Str, a2Str := must.Split2(a, "-")
		b1Str, b2Str := must.Split2(b, "-")
		a1 := must.Atoi(a1Str)
		a2 := must.Atoi(a2Str)
		b1 := must.Atoi(b1Str)
		b2 := must.Atoi(b2Str)
		must.LessOrEqual(a1, a2)
		must.LessOrEqual(b1, b2)
		// Intersect:
		i1 := max(a1, b1)
		i2 := min(a2, b2)
		if i1 == a1 && i2 == a2 || i1 == b1 && i2 == b2 {
			ans1++
		}
		if i1 <= i2 {
			ans2++
		}
	}
	must.NoError(scanner.Err())

	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}
