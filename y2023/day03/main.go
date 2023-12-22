package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

type Point struct {
	i, j int
}

func Solve(lines []string) (ans1, ans2 int) {
	field := fld.NewByteField(lines).AddBorder('.')
	allGears := map[Point][]int{}
	for i, row := range field {
		isN := false
		isAdj := false
		gears := containers.NewSet[Point]()
		var n int
		for j, ch := range row {
			if aoc.IsDigit(ch) {
				isN = true
				n = 10*n + int(ch-'0')
				for ni := i - 1; ni <= i+1; ni++ {
					for nj := j - 1; nj <= j+1; nj++ {
						ch := field[ni][nj]
						if ch != '.' && !aoc.IsDigit(ch) {
							isAdj = true
						}
						if ch == '*' {
							gears.Add(Point{ni, nj})
						}
					}
				}
			} else {
				if isN {
					for g := range gears {
						allGears[g] = append(allGears[g], n)
					}
				}
				if isN && isAdj {
					ans1 += n
				}
				n = 0
				isN = false
				isAdj = false
				clear(gears)
			}
		}
	}

	for _, nn := range allGears {
		if len(nn) == 2 {
			ans2 += nn[0] * nn[1]
		}
	}

	return ans1, ans2
}

func SolvePart1(lines []string) any {
	ans1, _ := Solve(lines)
	return ans1
}

func SolvePart2(lines []string) any {
	_, ans2 := Solve(lines)
	return ans2
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
