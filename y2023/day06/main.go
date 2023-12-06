package main

import (
	"fmt"
	"math"

	"github.com/theyoprst/adventofcode/aoc"
)

func main() {
	ans := 1
	ansV2 := 1
	lines := aoc.ReadInputLines()
	tt := aoc.Ints(lines[0])
	rr := aoc.Ints(lines[1])
	for i := range tt {
		t := tt[i]
		r := rr[i]
		wins := 0
		for x := 0; x <= t; x++ {
			y := tt[i] - x
			wins += aoc.BoolToInt(x*y > r)
		}
		ans *= wins

		// x * (t - x) > r <=> -x*x +tx -r > 0
		x1, x2 := aoc.SolveQuadratic(-1, t, -r)
		const eps = 1e-10
		p1 := int(math.Ceil(*x1 + eps))
		p2 := int(math.Floor(*x2 - eps))
		ansV2 *= p2 - p1 + 1
	}

	fmt.Println("Answer:", ans)
	fmt.Println("Answer analytical:", ansV2)
}
