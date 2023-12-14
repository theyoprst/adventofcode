package aoc

import (
	"bytes"
	"slices"
)

type ByteField [][]byte

func MakeByteField(lines []string) ByteField {
	var field [][]byte
	for _, line := range lines {
		field = append(field, []byte(line))
	}
	return field
}

func (f ByteField) AddBorder(b byte) ByteField {
	cols := len(f[0]) + 2
	res := make([][]byte, 0, len(f)+2)
	res = append(res, MakeSlice(b, cols))
	for _, s := range f {
		line := append(append([]byte{b}, s...), b)
		res = append(res, line)
	}
	res = append(res, MakeSlice(b, cols))
	return ByteField(res)
}

func (f ByteField) Transposed() ByteField {
	rows := len(f)
	cols := len(f[0])
	t := make([][]byte, cols)
	for col := range t {
		t[col] = make([]byte, rows)
	}
	for row, line := range f {
		for col, x := range line {
			t[col][row] = x
		}
	}
	return ByteField(t)
}

func (f ByteField) ReverseColumns() {
	for i := range f {
		slices.Reverse(f[i])
	}
}

func (f ByteField) ReverseRows() {
	slices.Reverse(f)
}

// Rotates the field clockwise.
func (f ByteField) RotateRight() ByteField {
	cols := len(f[0])
	ncols := len(f)
	nf := make([][]byte, cols)
	for nrow := range nf {
		nf[nrow] = make([]byte, ncols)
	}
	for nrow, line := range nf {
		for ncol := range line {
			nf[nrow][ncol] = f[ncols-ncol-1][nrow]
		}
	}
	return nf
}

// Rotates the field counter-clockwise.
func (f ByteField) RotateLeft() ByteField {
	cols := len(f[0])
	ncols := len(f)
	nf := make([][]byte, cols)
	for nrow := range nf {
		nf[nrow] = make([]byte, ncols)
	}
	for nrow, line := range nf {
		for ncol := range line {
			nf[nrow][ncol] = f[ncol][cols-nrow-1]
		}
	}
	return nf
}

func (f ByteField) String() string {
	return string(append(bytes.Join(f, []byte{'\n'}), '\n'))
}

func (f ByteField) Swap(rowA, colA, rowB, colB int) {
	f[rowA][colA], f[rowB][colB] = f[rowB][colB], f[rowA][colA]
}
