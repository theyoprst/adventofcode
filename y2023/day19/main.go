package main

import (
	"maps"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

// TODO: use regexps.

type Rule struct {
	key       string
	threshold int
	opCode    byte
	next      string
}

func (r Rule) Op(value int) bool {
	switch r.opCode {
	case '<':
		return value < r.threshold
	case '>':
		return value > r.threshold
	default:
		return true
	}
}

func parseWorkflow(s string) (name string, workflow []Rule) {
	var rulesStr string
	name, rulesStr = must.Split2(s, "{")
	rulesStr = rulesStr[:len(rulesStr)-1]
	for _, rule := range strings.Split(rulesStr, ",") {
		workflow = append(workflow, parseRule(rule))
	}
	return name, workflow
}

func parseRatings(s string) map[string]int {
	ratings := map[string]int{}
	s = s[1 : len(s)-1]
	for _, part := range strings.Split(s, ",") {
		key, value := must.Split2(part, "=")
		ratings[key] = must.Atoi(value)
	}
	return ratings
}

func parseRule(s string) Rule {
	colonI := strings.Index(s, ":")
	if colonI == -1 {
		return Rule{next: s}
	}
	ruleS := s[:colonI]
	next := s[colonI+1:]
	lessI := strings.Index(ruleS, "<")
	if lessI != -1 {
		return Rule{
			key:       ruleS[:lessI],
			threshold: must.Atoi(ruleS[lessI+1:]),
			opCode:    '<',
			next:      next,
		}
	}
	greaterI := strings.Index(ruleS, ">")
	if greaterI != -1 {
		return Rule{
			key:       ruleS[:greaterI],
			threshold: must.Atoi(ruleS[greaterI+1:]),
			opCode:    '>',
			next:      next,
		}
	}
	panic("Unreachable")
}

func SolvePart1(lines []string) any {
	groups := aoc.Split(lines, "")
	workflows := map[string][]Rule{}
	for _, workflowStr := range groups[0] {
		key, workflow := parseWorkflow(workflowStr)
		workflows[key] = workflow
	}

	runWorkflows := func(name string, ratings map[string]int) string {
	mainLoop:
		for name != "A" && name != "R" {
			workflow := workflows[name]
			for _, rule := range workflow {
				r := ratings[rule.key]
				if rule.Op(r) {
					name = rule.next
					continue mainLoop
				}
			}
			panic(name)
		}
		return name
	}

	ans := 0
	for _, ratingStr := range groups[1] {
		ratings := parseRatings(ratingStr)
		verdict := runWorkflows("in", ratings)
		if verdict == "A" {
			for _, r := range ratings {
				ans += r
			}
		}
	}
	return ans
}

func SolvePart2(lines []string) any {
	groups := aoc.Split(lines, "")
	workflows := map[string][]Rule{}
	for _, workflowStr := range groups[0] {
		key, workflow := parseWorkflow(workflowStr)
		workflows[key] = workflow
	}
	type Interval struct {
		first, after int
	}
	rect := map[string]Interval{}
	for _, ch := range "xmas" {
		rect[string(ch)] = Interval{
			first: 1,
			after: 4001,
		}
	}
	var countAccepted func(workflowName string, rect map[string]Interval) int
	countAccepted = func(workflowName string, rect map[string]Interval) int {
		rect = maps.Clone(rect)
		if workflowName == "R" {
			return 0
		}
		if workflowName == "A" {
			p := 1
			for _, interval := range rect {
				p *= interval.after - interval.first
			}
			return max(0, p)
		}
		count := 0
		for _, rule := range workflows[workflowName] {
			if rule.opCode == 0 {
				count += countAccepted(rule.next, rect)
				continue
			}

			src := rect[rule.key]
			if rule.opCode == '<' {
				// [first:T], [T:after]
				rect[rule.key] = Interval{src.first, rule.threshold}
				count += countAccepted(rule.next, rect)
				rect[rule.key] = Interval{rule.threshold, src.after}
			} else if rule.opCode == '>' {
				// [T+1:afterV], [firstV:T+1],
				rect[rule.key] = Interval{rule.threshold + 1, src.after}
				count += countAccepted(rule.next, rect)
				rect[rule.key] = Interval{src.first, rule.threshold + 1}
			} else {
				panic("Unreachable")
			}
		}
		return count
	}
	return countAccepted("in", rect)
}

var (
	solvers1 []aoc.Solver = []aoc.Solver{SolvePart1}
	solvers2 []aoc.Solver = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
