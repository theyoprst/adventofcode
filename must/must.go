package must

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
)

func RemovePrefix(s, p string) string {
	if !strings.HasPrefix(s, p) {
		panic(fmt.Sprintf("string %q has no prefix %q", s, p))
	}
	return s[len(p):]
}

func Split2(s string, sep string) (_, _ string) {
	split := strings.Split(s, sep)
	if len(split) != 2 {
		panic(fmt.Sprintf("Split %q by %q: got %d parts, want %d", s, sep, len(split), 2))
	}
	return split[0], split[1]
}

func Split3(s string, sep string) (_, _, _ string) {
	split := strings.Split(s, sep)
	if len(split) != 3 {
		panic(fmt.Sprintf("Split %q by %q: got %d parts, want %d", s, sep, len(split), 3))
	}
	return split[0], split[1], split[2]
}

func NoError(err error) {
	if err != nil {
		panic(err)
	}
}

func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	NoError(err)
	return n
}

func Equal[T comparable](got, target T) {
	if got != target {
		panic(fmt.Sprintf("Got %v, want %v", got, target))
	}
}

func True(b bool) {
	if !b {
		panic("Contition is false, want true")
	}
}

func NotEqual[T comparable](got, target T) {
	if got == target {
		panic(fmt.Sprintf("Got %v, want not equal to %v", got, target))
	}
}

func Greater[T cmp.Ordered](got, target T) {
	if got <= target {
		panic(fmt.Sprintf("Got %v, want greater than %v", got, target))
	}
}

func GreaterOrEqual[T cmp.Ordered](got, target T) {
	if got < target {
		panic(fmt.Sprintf("Got %v, want greater or equal than %v", got, target))
	}
}

func Less[T cmp.Ordered](got, target T) {
	if got >= target {
		panic(fmt.Sprintf("Got %v, want less than %v", got, target))
	}
}

func LessOrEqual[T cmp.Ordered](got, target T) {
	if got > target {
		panic(fmt.Sprintf("Got %v, want less or equal than %v", got, target))
	}
}

func Lookup[K comparable, V any](m map[K]V, k K) V {
	v, ok := m[k]
	if !ok {
		panic(fmt.Sprintf("No key '%v' in map %#v", k, m))
	}
	return v
}
