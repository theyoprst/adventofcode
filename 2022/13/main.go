package main

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(lines []string) any {
	blocks := aoc.Split(lines, "")
	ans := 0
	for i, block := range blocks {
		first := parsePacket(block[0])
		second := parsePacket(block[1])
		if packetsCompare(first, second) < 0 {
			ans += i + 1
		}
	}
	return ans
}

func SolvePart1JSON(lines []string) any {
	blocks := aoc.Split(lines, "")
	ans := 0
	for i, block := range blocks {
		first := parsePacketJSON(block[0])
		second := parsePacketJSON(block[1])
		if jsonPacketsCompare(first, second) < 0 {
			ans += i + 1
		}
	}
	return ans
}

func SolvePart2(lines []string) any {
	blocks := aoc.Split(lines, "")
	div1 := parsePacket("[[2]]")
	div2 := parsePacket("[[6]]")
	packets := []Packet{div1, div2}
	for _, block := range blocks {
		packets = append(packets, parsePacket(block[0]), parsePacket(block[1]))
	}
	slices.SortFunc(packets, packetsCompare)
	idx1 := slices.IndexFunc(packets, func(p Packet) bool {
		return packetsCompare(p, div1) == 0
	})
	idx2 := slices.IndexFunc(packets, func(p Packet) bool {
		return packetsCompare(p, div2) == 0
	})
	return (idx1 + 1) * (idx2 + 1)
}

func SolvePart2JSON(lines []string) any {
	blocks := aoc.Split(lines, "")
	div1 := parsePacketJSON("[[2]]")
	div2 := parsePacketJSON("[[6]]")
	packets := []any{div1, div2}
	for _, block := range blocks {
		packets = append(packets, parsePacketJSON(block[0]), parsePacketJSON(block[1]))
	}
	slices.SortFunc(packets, jsonPacketsCompare)
	idx1 := slices.IndexFunc(packets, func(p any) bool {
		return jsonPacketsCompare(p, div1) == 0
	})
	idx2 := slices.IndexFunc(packets, func(p any) bool {
		return jsonPacketsCompare(p, div2) == 0
	})
	return (idx1 + 1) * (idx2 + 1)
}

type PacketType int

const (
	PacketTypeNumber PacketType = 0
	PacketTypeList   PacketType = 1
)

type Packet struct {
	Type PacketType
	Num  int
	List []Packet
}

func parsePacket(str string) Packet {
	i := 0

	get := func() byte {
		return str[i]
	}

	next := func() {
		i++
	}

	consume := func(ch byte) {
		must.Equal(ch, get())
		next()
	}

	isDigit := func(ch byte) bool {
		return '0' <= ch && ch <= '9'
	}

	var parseAnyPacket, parseListPacket, parseNumberPacket func() Packet

	parseAnyPacket = func() Packet {
		if isDigit(get()) {
			return parseNumberPacket()
		}
		return parseListPacket()
	}

	parseNumberPacket = func() Packet {
		num := 0
		for isDigit(get()) {
			num = num*10 + int(get()-'0')
			next()
		}
		return Packet{Type: PacketTypeNumber, Num: num}
	}

	parseListPacket = func() Packet {
		consume('[')
		var list []Packet
		for get() != ']' {
			list = append(list, parseAnyPacket())
			if get() != ']' {
				consume(',')
			}
		}
		consume(']')
		return Packet{Type: PacketTypeList, List: list}
	}

	return parseListPacket()
}

func packetsCompare(a, b Packet) int {
	if a.Type == PacketTypeNumber && b.Type == PacketTypeNumber {
		return a.Num - b.Num
	}
	if a.Type == PacketTypeList && b.Type == PacketTypeList {
		for i := 0; i < min(len(a.List), len(b.List)); i++ {
			if cmp := packetsCompare(a.List[i], b.List[i]); cmp != 0 {
				return cmp
			}
		}
		return len(a.List) - len(b.List)
	}
	if a.Type == PacketTypeNumber {
		a = Packet{Type: PacketTypeList, List: []Packet{a}}
	} else {
		b = Packet{Type: PacketTypeList, List: []Packet{b}}
	}
	return packetsCompare(a, b)
}

func getJsonType(value any) PacketType {
	switch value.(type) {
	case float64:
		return PacketTypeNumber
	case []any:
		return PacketTypeList
	}
	panic(fmt.Sprintf("invalid type %T", value))
}

func jsonPacketsCompare(a, b any) int {
	aType := getJsonType(a)
	bType := getJsonType(b)
	if aType == PacketTypeNumber && bType == PacketTypeNumber {
		return int(a.(float64)) - int(b.(float64))
	}
	if aType == PacketTypeList && bType == PacketTypeList {
		aList := a.([]any)
		bList := b.([]any)
		for i := range min(len(aList), len(bList)) {
			if cmp := jsonPacketsCompare(aList[i], bList[i]); cmp != 0 {
				return cmp
			}
		}
		return len(aList) - len(bList)
	}
	if aType == PacketTypeNumber {
		a = []any{a}
	} else {
		b = []any{b}
	}
	return jsonPacketsCompare(a, b)
}

func parsePacketJSON(str string) any {
	var value any
	must.NoError(json.Unmarshal([]byte(str), &value))
	return value
}

var (
	solvers1 = []aoc.Solver{SolvePart1, SolvePart1JSON}
	solvers2 = []aoc.Solver{SolvePart2, SolvePart2JSON}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
