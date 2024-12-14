package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/fld"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(lines []string) any {
	const steps = 100
	// const width = 11
	// const height = 7
	const width = 101
	const height = 103
	quadrants := [][]int{
		{0, 0},
		{0, 0},
	}
	for _, line := range lines {
		p, v := parsePV(line)
		p = p.Add(v.Mult(steps))
		p.Row = (p.Row + steps*height) % height
		p.Col = (p.Col + steps*width) % width
		if p.Col == width/2 || p.Row == height/2 {
			continue
		}
		quadrants[aoc.BoolToInt(p.Row > height/2)][aoc.BoolToInt(p.Col >= width/2)]++
	}
	prod := 1
	for _, row := range quadrants {
		for _, q := range row {
			prod *= q
		}
	}
	return prod
}

func SolvePart2(lines []string) any {
	const width = 101
	const height = 103

	printGrid := func(robots []fld.Pos) {
		grid := make([][]byte, height)
		for i := range grid {
			grid[i] = make([]byte, width)
		}
		for _, line := range grid {
			for col := range line {
				line[col] = '.'
			}
		}
		for _, pos := range robots {
			grid[pos.Row][pos.Col] = '#'
		}
		var sb strings.Builder
		for _, line := range grid {
			for _, cell := range line {
				sb.WriteByte(cell)
			}
			sb.WriteByte('\n')
		}
		sb.WriteByte('\n')
		fmt.Print(sb.String())
	}
	_ = printGrid

	var robots []fld.Pos
	var velocities []fld.Pos
	for _, line := range lines {
		p, v := parsePV(line)
		robots = append(robots, p)
		velocities = append(velocities, v)
	}

	for time := 1; ; time++ {
		x := make([]int, 0, len(robots))
		y := make([]int, 0, len(robots))
		for j := range robots {
			robots[j] = robots[j].Add(velocities[j])
			robots[j].Row = (robots[j].Row + height) % height
			robots[j].Col = (robots[j].Col + width) % width
			x = append(x, robots[j].Col)
			y = append(y, robots[j].Row)
		}
		// Check that std dev for both x and y are less < 20.
		// Which means that most of the robots are in a small area, a square ~40x40.
		// Christmas tree is 33x31 (height x width), and std dev for this state is <19 for both x and y.
		// Kudos: https://www.reddit.com/r/adventofcode/comments/1hdvhvu/comment/m1zgdsh/
		if stddev(x) < 20 && stddev(y) < 20 {
			// printGrid(robots)
			return time
		}
	}
}

func SolvePart2Fast(lines []string) any {
	const width = 101
	const height = 103

	var x []int
	var y []int
	var dx []int
	var dy []int
	for _, line := range lines {
		a := aoc.Ints(line)
		x = append(x, a[0])
		y = append(y, a[1])
		dx = append(dx, a[2])
		dy = append(dy, a[3])
	}

	// Use the fact that situations in the grid have cycle 101*103: because 101 and 103 are primes,
	// and each robot cycles it's x coordinate in 101 steps, and y in 103 steps).
	// So we can iterate over 101*103 steps and find the minimum std dev / entropy / whatever relevane metric.
	// But we can only do 103 iteractions to find the best entropies for both x and y.
	// And then find the right time knowing the modulos by the width and the height.
	// Kudos: https://www.reddit.com/r/adventofcode/comments/1hdvhvu/comment/m1zws1g/

	var bestTimeX, bestTimeY int
	bestVarX := math.MaxInt64
	for i := range width {
		for j := range x {
			x[j] = (x[j] + dx[j] + width) % width
		}
		varX := varianceN(x)
		if varX < bestVarX {
			bestVarX = varX
			bestTimeX = i
		}
	}

	bestVarY := math.MaxInt64
	for i := range height {
		for j := range y {
			y[j] = (y[j] + dy[j] + height) % height
		}
		varY := varianceN(y)
		if varY < bestVarY {
			bestVarY = varY
			bestTimeY = i
		}
	}

	// t = bestVarX (mod width)
	// t = bestVarY (mod height)

	// Original solution: use inverse by mod (shoud be implemented):
	// return 1 + bestTimeX + ((inverse(width, height)*(bestTimeY-bestTimeX+height))%height)*width

	for t := range width * height {
		if t%width == bestTimeX && t%height == bestTimeY {
			return t + 1
		}
	}
	panic("no time found")
}

func parsePV(s string) (fld.Pos, fld.Pos) {
	a := aoc.Ints(s)
	must.Equal(len(a), 4)
	return fld.Pos{Col: a[0], Row: a[1]}, fld.Pos{Col: a[2], Row: a[3]}
}

func stddev(a []int) float64 {
	mean := 0.0
	for _, v := range a {
		mean += float64(v)
	}
	mean /= float64(len(a))
	variance := 0.0
	for _, v := range a {
		diff := float64(v) - mean
		variance += diff * diff
	}
	return math.Sqrt(variance / float64(len(a)))
}

// returns variance(a) * len(a)
func varianceN(a []int) int {
	mean := 0
	for _, v := range a {
		mean += v
	}
	mean /= len(a)
	variance := 0
	for _, v := range a {
		diff := v - mean
		variance += diff * diff
	}
	return variance
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2, SolvePart2Fast}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
