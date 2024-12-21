package main

import (
	"math"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

var (
	numericKeypad = buttonToPos(
		"789",
		"456",
		"123",
		" 0A",
	)
	directionalKeypad = buttonToPos(
		" ^A",
		"<v>",
	)
)

func SolvePart1(lines []string) any {
	return solve(lines, 2)
}

func SolvePart2Right(lines []string) any {
	return solve(lines, 25)
}

func buttonToPos(lines ...string) map[byte]fld.Pos {
	keypad := make(map[byte]fld.Pos)
	for r, row := range lines {
		for c, char := range row {
			keypad[byte(char)] = fld.NewPos(r, c)
		}
	}
	return keypad
}

func solve(lines []string, robots int) int {
	// Idea.
	// When any text is typed on a "keyboard" using a directional keypad, for each char we know how "cursor" should be moved:
	// from which position (previous char or 'A' if the beginning) to which position (current char position in the keypad).
	// There are multiple ways to move the "cursor" from the previous char to the current char.
	// But only two are optimal: first move vertically and then horizontally, or first move horizontally and then vertically.
	// If one of them goes through no-button (empty) button, it should be discarded.
	// Others are not optimal because they can be shortened to the optimal ones.
	// So, we can try both ways and choose the one that requires the minimum number of key presses.
	// This selection is independent of other chars in the code, and could be done for each char separately.
	// Fore each char the sequence of commands which outputs it depends only of two parameters:
	// 1) the "jump" between positions we need to make and
	// 2) the keypad index we are currently working with.
	// Second parameter is 25 for the second part of task.
	// First one is just a number of unique short paths on the keypad, which is < 2*4*5 = 20, taking 2 shortest optimal paths between
	// each pair of 5 buttons on numerical (24 to be exact).
	// So we can use dynamic programming with caching recursive calls, and there will be up to 500-1000 states, which
	// is pretty much fast.

	type cacheKey struct {
		code        string
		restKeypads int
	}
	cache := make(map[cacheKey]int)

	// minPresses returns the minimum number of key presses on the last (human-typed) keypad
	// to enter the code on the current keypad (with index keypadIdx).
	var minPresses func(code string, keypadIdx int) int
	minPresses = func(code string, keypadIdx int) int {
		keypad := directionalKeypad
		if keypadIdx == 0 {
			keypad = numericKeypad
		}
		if keypadIdx > robots {
			return len(code)
		}
		if cached, ok := cache[cacheKey{code, keypadIdx}]; ok {
			return cached
		}
		sum := 0
		cur := keypad['A']
		for _, char := range []byte(code) {
			next := keypad[char]
			movesVert := strings.Repeat("v", max(0, next.Row-cur.Row)) + strings.Repeat("^", max(0, cur.Row-next.Row))
			movesHor := strings.Repeat(">", max(0, next.Col-cur.Col)) + strings.Repeat("<", max(0, cur.Col-next.Col))
			curMin := math.MaxInt
			if fld.NewPos(next.Row, cur.Col) != keypad[' '] {
				// Try moving vertically first: check that there is soem button in the corner.
				curMin = min(curMin, minPresses(movesVert+movesHor+"A", keypadIdx+1))
			}
			if fld.NewPos(cur.Row, next.Col) != keypad[' '] {
				// Try moving horizontally first: check that there is soem button in the corner.
				curMin = min(curMin, minPresses(movesHor+movesVert+"A", keypadIdx+1))
			}
			// Note that other paths are not considered because they are not optimal.
			sum += curMin
			cur = next
		}
		cache[cacheKey{code, keypadIdx}] = sum
		return sum
	}

	sum := 0
	for _, code := range lines {
		num := aoc.Ints(code)[0]
		sum += num * minPresses(code, 0)
	}
	return sum
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2Right}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
