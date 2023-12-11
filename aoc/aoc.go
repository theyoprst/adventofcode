package aoc

import (
	"bufio"
	"cmp"
	"io"
	"math"
	"os"
	"regexp"
	"slices"

	"github.com/theyoprst/adventofcode/must"
	"golang.org/x/exp/constraints"
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

func ToBytesField(lines []string) [][]byte {
	var field [][]byte
	for _, line := range lines {
		field = append(field, []byte(line))
	}
	return field
}

func MakeSlice[T any](elem T, n int) []T {
	s := make([]T, n)
	for i := range s {
		s[i] = elem
	}
	return s
}

func AddBorder2D[T any](a [][]T, b T) [][]T {
	cols := len(a[0]) + 2
	res := make([][]T, 0, len(a)+2)
	res = append(res, MakeSlice(b, cols))
	for _, s := range a {
		line := append(append([]T{b}, s...), b)
		res = append(res, line)
	}
	res = append(res, MakeSlice(b, cols))
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

func MapSortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	var kk []K
	for k := range m {
		kk = append(kk, k)
	}
	slices.Sort(kk)
	return kk
}

// Greates Common Divisor.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Least Common Multiple.
func LCM(a, b int) int {
	a /= GCD(a, b)
	must.Less(b, math.MaxInt/a)
	return a * b
}

func Reversed[S ~[]E, E any](a S) []E {
	r := slices.Clone(a)
	slices.Reverse(r)
	return r
}

func ToSet[S ~[]E, E comparable](a S) map[E]bool {
	set := map[E]bool{}
	for _, x := range a {
		set[x] = true
	}
	return set
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
