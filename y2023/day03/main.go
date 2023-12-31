package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

func SolvePart1(lines []string) any {
	field := fld.NewByteField(lines).AddBorder('.')
	var ans int
	for i, row := range field {
		isN := false
		isAdj := false
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
					}
				}
			} else {
				if isN && isAdj {
					ans += n
				}
				n = 0
				isN = false
				isAdj = false
			}
		}
	}

	return ans
}

func SolvePart2(lines []string) any {
	field := fld.NewByteField(lines).AddBorder('.')
	allGears := map[fld.Pos][]int{}
	for i, row := range field {
		isN := false
		gears := containers.NewSet[fld.Pos]()
		var n int
		for j, ch := range row {
			if aoc.IsDigit(ch) {
				isN = true
				n = 10*n + int(ch-'0')
				for ni := i - 1; ni <= i+1; ni++ {
					for nj := j - 1; nj <= j+1; nj++ {
						pos := fld.NewPos(ni, nj)
						ch := field.Get(pos)
						if ch == '*' {
							gears.Add(pos)
						}
					}
				}
			} else {
				if isN {
					for g := range gears {
						allGears[g] = append(allGears[g], n)
					}
				}
				n = 0
				isN = false
				clear(gears)
			}
		}
	}

	var ans int
	for _, nn := range allGears {
		if len(nn) == 2 {
			ans += nn[0] * nn[1]
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
