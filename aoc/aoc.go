package aoc

import (
	"bufio"
	"cmp"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/theyoprst/adventofcode/must"
)

func ReadInputLines() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	must.NoError(scanner.Err())
	return lines
}

func AddBorder2D(a []string, r rune) []string {
	b := string(r)
	cols := len(a[0]) + 2
	res := make([]string, 0, len(a)+2)
	res = append(res, strings.Repeat(b, cols))
	for _, s := range a {
		res = append(res, b+s+b)
	}
	res = append(res, strings.Repeat(b, cols))
	return res
}

func IsDigit[T byte | rune](ch T) bool {
	return '0' <= ch && ch <= '9'
}

func Split[T comparable](a []T, by T) [][]T {
	var g []T
	var gg [][]T
	for _, x := range append(a, by) {
		if x == by {
			gg = append(gg, g)
			g = []T{}
		} else {
			g = append(g, x)
		}
	}
	return gg
}

var allIntsRe = regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

func Ints(s string) []int {
	var ints []int
	for _, word := range allIntsRe.FindAllString(s, -1) {
		ints = append(ints, must.Atoi(word))
	}
	return ints
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Solves equation a*x^2 + b*x + c = 0. Returns x1, x2: x1 <= x2.
func SolveQuadratic(a, b, c int) (_, _ *float64) {
	af := float64(a)
	bf := float64(b)
	cf := float64(c)
	d := math.Sqrt(bf*bf - 4*af*cf)
	if math.IsNaN(d) {
		return nil, nil
	}
	x1 := (-bf - d) / 2 / af
	x2 := (-bf + d) / 2 / af
	x1, x2 = min(x1, x2), max(x1, x2)
	return &x1, &x2
}

func MapSortedValues[K comparable, V cmp.Ordered](m map[K]V) []V {
	var vv []V
	for _, v := range m {
		vv = append(vv, v)
	}
	slices.Sort(vv)
	return vv
}
