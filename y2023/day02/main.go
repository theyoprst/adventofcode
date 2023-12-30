package main

import (
	"log"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

type Handful struct {
	r, g, b int
}

func union(hh []Handful) Handful {
	var res Handful
	for _, h := range hh {
		res.r = max(res.r, h.r)
		res.g = max(res.g, h.g)
		res.b = max(res.b, h.b)
	}
	return res
}

func SolvePart1(lines []string) any {
	var ans int
	for gameI, line := range lines {
		handfuls := parseGame(line)
		u := union(handfuls)
		if u.r <= 12 && u.g <= 13 && u.b <= 14 {
			ans += gameI + 1
		}
	}
	return ans
}

func SolvePart2(lines []string) any {
	var ans int
	for _, line := range lines {
		handfuls := parseGame(line)
		u := union(handfuls)
		ans += u.r * u.b * u.g
	}
	return ans
}

func parseGame(game string) []Handful {
	_, game = must.Split2(game, ":")
	var handfuls []Handful
	for _, try := range strings.Split(game, ";") {
		var handful Handful
		for _, colorN := range strings.Split(try, ",") {
			colorN = strings.TrimSpace(colorN)
			nStr, color := must.Split2(colorN, " ")
			n := must.Atoi(nStr)
			switch color {
			case "red":
				must.Equal(handful.r, 0)
				handful.r = n
			case "green":
				must.Equal(handful.g, 0)
				handful.g = n
			case "blue":
				must.Equal(handful.b, 0)
				handful.b = n
			default:
				panic("unreachable")
			}
		}
		handfuls = append(handfuls, handful)
	}
	return handfuls
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	log.SetFlags(0)
	aoc.Main(solvers1, solvers2)
}
