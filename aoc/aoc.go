package aoc

import (
	"bufio"
	"cmp"
	"io"
	"math"
	"os"
	"regexp"
	"slices"

	"golang.org/x/exp/constraints"

	"github.com/theyoprst/adventofcode/must"
)

func ReadInputLines() []string {
	return ReadLines(os.Stdin)
}

func ReadLines(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	must.NoError(scanner.Err())
	return lines
}

func MakeSlice[T any](elem T, n int) []T {
	s := make([]T, n)
	for i := range s {
		s[i] = elem
	}
	return s
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

var allIntsRe = regexp.MustCompile(`[-+]?\d+`)

func Ints(s string) []int {
	words := allIntsRe.FindAllString(s, -1)
	ints := make([]int, len(words))
	for i, word := range words {
		ints[i] = must.Atoi(word)
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
// Returns NaN if not real solutions.
func SolveQuadratic(a, b, c int) (x1, x2 float64) {
	af := float64(a)
	bf := float64(b)
	cf := float64(c)
	d := math.Sqrt(bf*bf - 4*af*cf)
	x1 = (-bf - d) / 2 / af
	x2 = (-bf + d) / 2 / af
	return min(x1, x2), max(x1, x2)
}

func MapSortedValues[K comparable, V cmp.Ordered](m map[K]V) []V {
	vv := make([]V, 0, len(m))
	for _, v := range m {
		vv = append(vv, v)
	}
	slices.Sort(vv)
	return vv
}

// Greates Common Divisor.
func GCD(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		res = gcd(res, a[i])
	}
	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Least Common Multiple.
func LCM(a ...int) int {
	res := 1
	for _, x := range a {
		res = lcm(res, x)
	}
	return res
}

func lcm(a, b int) int {
	a /= GCD(a, b)
	must.Less(b, math.MaxInt/a)
	return a * b
}

func Reversed[S ~[]E, E any](a S) []E {
	r := slices.Clone(a)
	slices.Reverse(r)
	return r
}

func Abs[T constraints.Signed | constraints.Float](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

// PartialSum returns partial sum list based on a: []int{a[0], a[0]+a[1], ...}}.
func PartialSum[T constraints.Integer | constraints.Float](a []T) []T {
	var sum T
	partial := make([]T, len(a))
	for i, x := range a {
		sum += x
		partial[i] = sum
	}
	return partial
}

func CountBinaryOnes[T constraints.Integer](n T) int {
	must.GreaterOrEqual(n, 0)
	ones := 0
	for n > 0 {
		ones++
		n &= n - 1
	}
	return ones
}

func MapContains[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}
