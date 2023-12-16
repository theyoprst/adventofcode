package fld

import (
	"bytes"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
)

type Field[T any] [][]T

type ByteField = Field[byte]

func NewByteField(lines []string) Field[byte] {
	var field [][]byte
	for _, line := range lines {
		field = append(field, []byte(line))
	}
	return field
}

func (f Field[T]) Rows() int {
	return len(f)
}

func (f Field[T]) Cols() int {
	return len(f[0])
}

func (f Field[T]) AddBorder(b T) Field[T] {
	cols := f.Cols() + 2
	res := make([][]T, 0, len(f)+2)
	res = append(res, aoc.MakeSlice(b, cols))
	for _, s := range f {
		line := append(append([]T{b}, s...), b)
		res = append(res, line)
	}
	res = append(res, aoc.MakeSlice(b, cols))
	return res
}

func (f Field[T]) Transpose() Field[T] {
	rows := len(f)
	cols := len(f[0])
	t := make([][]T, cols)
	for col := range t {
		t[col] = make([]T, rows)
	}
	for row, line := range f {
		for col, x := range line {
			t[col][row] = x
		}
	}
	return t
}

func (f Field[T]) ReverseColumns() {
	for i := range f {
		slices.Reverse(f[i])
	}
}

func (f Field[T]) ReverseRows() {
	slices.Reverse(f)
}

// Rotates the field clockwise.
func (f Field[T]) RotateRight() Field[T] {
	cols := len(f[0])
	ncols := len(f)
	nf := make([][]T, cols)
	for nrow := range nf {
		nf[nrow] = make([]T, ncols)
	}
	for nrow, line := range nf {
		for ncol := range line {
			nf[nrow][ncol] = f[ncols-ncol-1][nrow]
		}
	}
	return nf
}

// Rotates the field counter-clockwise.
func (f Field[T]) RotateLeft() Field[T] {
	cols := len(f[0])
	ncols := len(f)
	nf := make([][]T, cols)
	for nrow := range nf {
		nf[nrow] = make([]T, ncols)
	}
	for nrow, line := range nf {
		for ncol := range line {
			nf[nrow][ncol] = f[ncol][cols-nrow-1]
		}
	}
	return nf
}

// ToString() print table of characters.
// Can't make it as a method because of Go generics limitations.
func ToString(f Field[byte]) string {
	return string(append(bytes.Join(f, []byte{'\n'}), '\n'))
}

func (f Field[T]) Swap(a, b Pos) {
	f[a.Row][a.Col], f[b.Row][b.Col] = f[b.Row][b.Col], f[a.Row][a.Col]
}

func (f Field[T]) Inside(pos Pos) bool {
	return 0 <= pos.Row && pos.Row < len(f) &&
		0 <= pos.Col && pos.Col < len(f[pos.Row])
}
