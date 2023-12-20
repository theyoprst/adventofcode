package main

import (
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

type Impulse struct {
	value Imp
	from  string
}

type Imp int

const (
	Low Imp = iota
	High
)

type FlipFlop int

const (
	On FlipFlop = iota
	Off
)

func SolvePart1(lines []string) any {
	flip := map[string]FlipFlop{}
	conj := map[string]map[string]Imp{}
	dispatch := map[string][]string{}
	for _, line := range lines {
		src, dstsStr := must.Split2(line, " -> ")
		dsts := strings.Split(dstsStr, ", ")
		if src == "broadcaster" {
		} else if src[0] == '%' {
			// flip-flop
			src = src[1:]
			flip[src] = Off
		} else if src[0] == '&' {
			// conjunction
			src = src[1:]
			conj[src] = map[string]Imp{}
		} else {
			panic("unreachable")
		}
		dispatch[src] = dsts
	}
	for src, dsts := range dispatch {
		for _, dst := range dsts {
			if conj[dst] != nil {
				conj[dst][src] = Low
			}
		}
	}
	lows := 0
	highs := 0
	for step := 0; step < 1000; step++ {
		impulses := map[string]Impulse{
			"broadcaster": {value: Low, from: ""},
		}
		for len(impulses) > 0 {
			// fmt.Println(impulses)
			nImpulses := map[string]Impulse{}
			for cur, impulse := range impulses {
				if impulse.value == Low {
					lows++
				} else {
					highs++
				}
				nImp := impulse.value
				if _, ok := flip[cur]; ok {
					if impulse.value == High {
						continue
					}
					flip[cur] = 1 - flip[cur]
					nImp = High
					if flip[cur] == Off {
						nImp = Low
					}
				} else if conj[cur] != nil {
					conj[cur][impulse.from] = impulse.value
					nImp = Low
					for _, imp := range conj[cur] {
						if imp != High {
							nImp = High
							break
						}
					}
				}
				for _, dst := range dispatch[cur] {
					if _, ok := nImpulses[dst]; ok {
						panic(dst)
					}
					nImpulses[dst] = Impulse{value: nImp, from: cur}
					// TODO: check if exists?
				}
			}
			impulses = nImpulses
		}
	}
	return highs * lows
}

func SolvePart2(lines []string) any {
	_ = lines
	return 0
}

var (
	solvers1 []aoc.Solver = []aoc.Solver{SolvePart1}
	solvers2 []aoc.Solver = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
