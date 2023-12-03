package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/theyoprst/adventofcode/helpers"
)

type Point struct {
	i, j int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var ans1, ans2 int
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	field := helpers.Border2D(lines, '.')

	adjacentSymbols := func(i, j int) map[byte][]Point {
		res := map[byte][]Point{}
		for di := -1; di <= 1; di++ {
			for dj := -1; dj <= 1; dj++ {
				i0 := i + di
				j0 := j + dj
				ch := field[i0][j0]
				if ch != '.' && !helpers.IsDigit(ch) {
					res[ch] = append(res[ch], Point{i0, j0})
				}
			}
		}
		return res
	}
	allGears := map[Point][]int{}
	for i, row := range field {
		isN := false
		isAdj := false
		gears := map[Point]bool{}
		var n int
		for j, ch := range row {
			if helpers.IsDigit(ch) {
				isN = true
				n = 10*n + int(ch-'0')
				symbols := adjacentSymbols(i, j)
				if len(symbols) > 0 {
					isAdj = true
				}
				for _, g := range symbols['*'] {
					gears[g] = true
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
	fmt.Println("Part 1:", ans1)

	for _, nn := range allGears {
		if len(nn) == 2 {
			ans2 += nn[0] * nn[1]
		}
	}
	fmt.Println("Part 2:", ans2)
}
