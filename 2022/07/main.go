package main

import (
	"context"
	"math"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

const (
	cdPrefix = "$ cd "
	lsCmd    = "$ ls"
)

func SolvePart1(_ context.Context, lines []string) any {
	dirs := parseDirs(lines)
	ans := 0
	for _, size := range dirs {
		if size <= 100000 {
			ans += size
		}
	}
	return ans
}

func SolvePart2(_ context.Context, lines []string) any {
	dirs := parseDirs(lines)
	free := 70000000 - dirs[""]
	toFree := 30000000 - free
	must.Greater(toFree, 0)
	minSize := math.MaxInt
	for _, size := range dirs {
		if size >= toFree && size < minSize {
			minSize = size
		}
	}
	return minSize
}

func parseDirs(lines []string) map[string]int {
	var path []string
	dirs := map[string]int{}
	for len(lines) > 0 {
		line := lines[0]
		lines = lines[1:]
		if strings.HasPrefix(line, cdPrefix) {
			target := line[len(cdPrefix):]
			switch target {
			case "/":
				path = path[:0]
			case "..":
				path = path[:len(path)-1]
			default:
				path = append(path, target)
			}
		} else {
			must.Equal(line, lsCmd)
			sum := 0
			for len(lines) > 0 && lines[0][0] != '$' {
				line := lines[0]
				lines = lines[1:]
				if !strings.HasPrefix(line, "dir ") {
					sizeStr, _ := must.Split2(line, " ")
					sum += must.Atoi(sizeStr)
				}
			}
			for i := 0; i <= len(path); i++ { // TODO: go 1.22
				name := strings.Join(path[0:i], "/")
				dirs[name] += sum
			}
		}
	}
	return dirs
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
