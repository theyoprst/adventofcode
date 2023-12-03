package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/theyoprst/adventofcode/must"
)

type Handful struct {
	r, g, b int
}

func parseGame(game string) []Handful {
	var handfuls []Handful
	for _, try := range strings.Split(game, ";") {
		var handful Handful
		for _, colorN := range strings.Split(try, ",") {
			colorN = strings.TrimSpace(colorN)
			nStr, color := must.Split2(colorN, " ")
			n := must.Atoi(nStr)
			must.Greater(n, 0)
			must.Less(n, 100000)
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

func isPossible(hh []Handful) bool {
	for _, h := range hh {
		if h.r > 12 || h.g > 13 || h.b > 14 {
			return false
		}
	}
	return true
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var sum1, sum2 int
	for scanner.Scan() {
		line := must.RemovePrefix(scanner.Text(), "Game ")
		gameI, game := must.Split2(line, ":")
		handfuls := parseGame(game)
		if isPossible(handfuls) {
			sum1 += must.Atoi(gameI)
		}
		u := union(handfuls)
		sum2 += u.r * u.b * u.g
	}
	fmt.Println("Ans1:", sum1)
	fmt.Println("Ans2:", sum2)
}
