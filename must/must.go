package must

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
)

// RemovePrefix removes the prefix `p` from `s` and returns the result. If `s` does not start with `p`, it panics.
func RemovePrefix(s, p string) string {
	if !strings.HasPrefix(s, p) {
		panic(fmt.Sprintf("string %q has no prefix %q", s, p))
	}
	return s[len(p):]
}

// Split2 splits s by sep and returns the two parts.
// If s does not contain exactly one occurrence of sep, it panics.
// If you expect not-fixed occurrences of sep between integer numbers, use aoc.Ints() instead.
func Split2(s string, sep string) (_, _ string) {
	split := strings.Split(s, sep)
	if len(split) != 2 {
		panic(fmt.Sprintf("Split %q by %q: got %d parts, want %d", s, sep, len(split), 2))
	}
	return split[0], split[1]
}

// Split3 splits s by sep and returns the three parts.
// If s does not contain exactly two occurrence of sep, it panics.
// If you expect not-fixed occurrences of sep between integer numbers, use aoc.Ints() instead.
func Split3(s string, sep string) (_, _, _ string) {
	split := strings.Split(s, sep)
	if len(split) != 3 {
		panic(fmt.Sprintf("Split %q by %q: got %d parts, want %d", s, sep, len(split), 3))
	}
	return split[0], split[1], split[2]
}

// NoError panics if err != nil.
func NoError(err error) {
	if err != nil {
		panic(err)
	}
}

// Atoi is a wrapper around strconv.Atoi that panics on error.
func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	NoError(err)
	return n
}

// Equal asserts that got == target (it panics otherwise).
func Equal[T comparable](got, target T) {
	if got != target {
		panic(fmt.Sprintf("Got %v, want %v", got, target))
	}
}

// True asserts that b is true (it panics otherwise).
func True(b bool) {
	if !b {
		panic("Contition is false, want true")
	}
}

// False asserts that b is false (it panics otherwise).
func False(b bool) {
	if b {
		panic("Contition is true, want false")
	}
}

// NotEqual asserts that got != target (it panics otherwise).
func NotEqual[T comparable](got, target T) {
	if got == target {
		panic(fmt.Sprintf("Got %v, want not equal to %v", got, target))
	}
}

// Greater asserts that got > target (it panics otherwise).
func Greater[T cmp.Ordered](got, target T) {
	if got <= target {
		panic(fmt.Sprintf("Got %v, want greater than %v", got, target))
	}
}

// GreaterOrEqual asserts that got >= target (it panics otherwise).
func GreaterOrEqual[T cmp.Ordered](got, target T) {
	if got < target {
		panic(fmt.Sprintf("Got %v, want greater or equal than %v", got, target))
	}
}

// Less asserts that got < target (it panics otherwise).
func Less[T cmp.Ordered](got, target T) {
	if got >= target {
		panic(fmt.Sprintf("Got %v, want less than %v", got, target))
	}
}

// LessOrEqual asserts that got <= target (it panics otherwise).
func LessOrEqual[T cmp.Ordered](got, target T) {
	if got > target {
		panic(fmt.Sprintf("Got %v, want less or equal than %v", got, target))
	}
}

// Lookup returns the value for key `k` in map `m`. If `k` is not in `m`, it panics.
func Lookup[K comparable, V any](m map[K]V, k K) V {
	v, ok := m[k]
	if !ok {
		panic(fmt.Sprintf("No key '%v' in map %#v", k, m))
	}
	return v
}
