package fld

import (
	"bytes"
	"fmt"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

type Field[T comparable] [][]T

type ByteField = Field[byte]

func NewByteField(lines []string) Field[byte] {
	field := make([][]byte, len(lines))
	for i, line := range lines {
		field[i] = []byte(line)
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

func (f Field[T]) ReverseColumnsInPlace() {
	for i := range f {
		slices.Reverse(f[i])
	}
}

func (f Field[T]) ReverseRowsInPlace() {
	slices.Reverse(f)
}

func (f Field[T]) Clone() Field[T] {
	cloned := make(Field[T], f.Rows())
	for row := range cloned {
		cloned[row] = slices.Clone(f[row])
	}
	return cloned
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

func (f Field[T]) Get(p Pos) T {
	return f[p.Row][p.Col]
}

func (f Field[T]) Set(p Pos, val T) {
	f[p.Row][p.Col] = val
}

func (f Field[T]) FindFirst(val T) Pos {
	var p Pos
	for row := range f {
		for col := range f[row] {
			if f[row][col] == val {
				p.Row = row
				p.Col = col
				return p
			}
		}
	}
	must.NoError(fmt.Errorf("cannot find %v", val))
	return p
}
