package aoc

import "slices"

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

func (f ByteField) Transpose() ByteField {
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
