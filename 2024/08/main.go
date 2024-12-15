package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

func SolvePart1(lines []string) any {
	field := fld.NewByteField(lines)
	antennas := make(map[byte][]fld.Pos)
	for pos := range field.IterPositions() {
		ch := field.Get(pos)
		if ch != '.' {
			antennas[ch] = append(antennas[ch], pos)
		}

	}
	antinodes := containers.NewSet[fld.Pos]()
	for _, positions := range antennas {
		for i, posA := range positions {
			for j, posB := range positions {
				if i == j {
					continue
				}
				antiPos := posA.Mult(2).Sub(posB)
				if field.Inside(antiPos) {
					antinodes.Add(antiPos)
				}
			}
		}
	}
	return len(antinodes)
}

func SolvePart2(lines []string) any {
	field := fld.NewByteField(lines)
	antennas := make(map[byte][]fld.Pos)
	for row := range field.Rows() {
		for col := range field.Cols() {
			pos := fld.NewPos(row, col)
			ch := field.Get(pos)
			if ch != '.' {
				antennas[ch] = append(antennas[ch], pos)
			}
		}
	}
	antinodes := containers.NewSet[fld.Pos]()
	for _, positions := range antennas {
		for i, posA := range positions {
			for j, posB := range positions {
				if i == j {
					continue
				}
				diff := posA.Sub(posB)
				pos := posA
				for field.Inside(pos) {
					antinodes.Add(pos)
					pos = pos.Add(diff)
				}
			}
		}
	}
	return len(antinodes)
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
