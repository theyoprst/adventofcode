package main

import (
	"maps"
	"regexp"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

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

var ruleRe = regexp.MustCompile(`^(([xmas])([<>])(\d+):)?(\w+)$`)

func parseRule(s string) Rule {
	sub := ruleRe.FindStringSubmatch(s)
	var opCode byte
	var threshold int
	if sub[1] != "" {
		opCode = sub[3][0]
		threshold = must.Atoi(sub[4])
	}
	return Rule{
		key:       sub[2],
		threshold: threshold,
		opCode:    opCode,
		next:      sub[5],
	}
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
