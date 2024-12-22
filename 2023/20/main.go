package main

import (
	"context"
	"slices"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

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

func (ff FlipFlop) Invert() FlipFlop {
	return 1 - ff
}

const Broadcaster = "broadcaster"

type Impulse struct {
	from  string
	to    string
	value Imp
}

func SolvePart1(_ context.Context, lines []string) any {
	flip := map[string]FlipFlop{}
	conj := map[string]map[string]Imp{}
	dispatch := map[string][]string{}
	for _, line := range lines {
		src, dstsStr := must.Split2(line, " -> ")
		dsts := strings.Split(dstsStr, ", ")
		switch src[0] {
		case '%':
			// flip-flop
			src = src[1:]
			flip[src] = Off
		case '&':
			// conjunction
			src = src[1:]
			conj[src] = map[string]Imp{}
		default:
			must.Equal(src, Broadcaster)
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
		impulses := []Impulse{{from: "", to: Broadcaster, value: Low}}
		for len(impulses) > 0 {
			impulse := impulses[0]
			impulses = impulses[1:]
			cur := impulse.to
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
				flip[cur] = flip[cur].Invert()
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
				impulses = append(impulses, Impulse{
					value: nImp, from: cur, to: dst,
				})
			}
		}
	}
	return highs * lows
}

func SolvePart2(_ context.Context, lines []string) any {
	flip := map[string]FlipFlop{}
	conj := map[string]map[string]Imp{}
	dispatch := map[string][]string{}
	for _, line := range lines {
		src, dstsStr := must.Split2(line, " -> ")
		dsts := strings.Split(dstsStr, ", ")
		switch src[0] {
		case '%':
			// flip-flop
			src = src[1:]
			flip[src] = Off
		case '&':
			// conjunction
			src = src[1:]
			conj[src] = map[string]Imp{}
		default:
			must.Equal(src, Broadcaster)
		}
		dispatch[src] = dsts
	}
	// Want Low signal on "rx", let's see who can send it.
	for src, dsts := range dispatch {
		for _, dst := range dsts {
			if conj[dst] != nil {
				conj[dst][src] = Low
			}
		}
	}
	var rxSources []string
	for src, dsts := range dispatch {
		if slices.Contains(dsts, "rx") {
			rxSources = append(rxSources, src)
		}
	}
	must.Equal(len(rxSources), 1)
	rxSrc := rxSources[0]             // In our test there is only one source of "rx".
	must.Greater(len(conj[rxSrc]), 1) // In our test source of "rx" is conjunction of several other nodes.
	cycles := map[string]int{}

	// Now we need to find step where all nodes from "waitFor" emit High signal.
	// So that their conjunction will emit Low signal.

	for step := 1; step < 10000; step++ {
		impulses := []Impulse{{from: "", to: "broadcaster", value: Low}}
		for len(impulses) > 0 {
			// fmt.Println(impulses)
			nImpulses := []Impulse{}
			for _, impulse := range impulses {
				cur := impulse.to
				// if impulse.value == Low && impulse.to == "rx" {
				// 	return step
				// }
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
				// fmt.Println(cur, dispatch[cur])
				for _, dst := range dispatch[cur] {
					nImpulses = append(nImpulses, Impulse{
						value: nImp, from: cur, to: dst,
					})
				}
			}
			for name, count := range conj[rxSrc] {
				if count > 0 {
					if cycles[name] == 0 {
						cycles[name] = step
					} else {
						must.Equal(step%cycles[name], 0)
					}
				}
			}
			impulses = nImpulses
		}
	}
	must.Equal(len(cycles), len(conj[rxSrc]))
	lcm := 1
	for _, x := range cycles {
		lcm = aoc.LCM(lcm, x)
	}
	return lcm
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
