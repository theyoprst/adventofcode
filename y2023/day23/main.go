package main

import (
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(lines []string) any {
	field := fld.NewByteField(lines)
	var topsort []fld.Pos
	var dfs func(p fld.Pos)
	seen := containers.NewSet[fld.Pos]()
	getDirs := func(p fld.Pos) []fld.Pos {
		var dirs []fld.Pos
		switch field.Get(p) {
		case '.':
			dirs = fld.DirsSimple
		case '>':
			dirs = []fld.Pos{fld.Right}
		case 'v':
			dirs = []fld.Pos{fld.Down}
		case '#':
			dirs = nil
		default:
			panic(string(field.Get(p)))
		}
		return dirs
	}
	dfs = func(p fld.Pos) {
		seen.Add(p)
		for _, dir := range getDirs(p) {
			np := p.Add(dir)
			if field.Inside(np) && field.Get(np) != '#' && !seen.Has(np) {
				if !slices.Equal(getDirs(np), []fld.Pos{dir.Reverse()}) {
					dfs(np)
				}
			}
		}
		topsort = append(topsort, p)
	}
	startPos := fld.NewPos(0, 1)
	dfs(startPos)
	slices.Reverse(topsort)
	// fmt.Println(topsort)

	dp := map[fld.Pos]int{startPos: 0}
	prevs := map[fld.Pos]fld.Pos{}
	for i := 1; i < len(topsort); i++ { // skip start point
		cur := topsort[i]
		must.True(field.Inside(cur))
		maxLen := -1
		var from fld.Pos
		for _, dir := range fld.DirsSimple {
			prev := cur.Add(dir)
			if !field.Inside(prev) {
				continue
			}
			for _, dir := range getDirs(prev) {
				if prev.Add(dir) == cur {
					if val, ok := dp[prev]; ok && val > maxLen {
						maxLen = val
						from = prev
					}
					break
				}
			}
		}
		dp[cur] = maxLen + 1
		prevs[cur] = from
		// fmt.Println(cur, dp[cur])
	}
	finishPos := fld.NewPos(field.Rows()-1, field.Cols()-2)
	pos := finishPos
	for pos != startPos {
		// fmt.Println(pos)
		field.Set(pos, 'O')
		pos = prevs[pos]
	}
	field.Set(startPos, 'S')
	// fmt.Println(fld.ToString(field))
	return dp[finishPos]
}

func SolvePart2(lines []string) any {
	field := fld.NewByteField(lines)
	startPos := fld.NewPos(0, 1)
	finishPos := fld.NewPos(field.Rows()-1, field.Cols()-2)

	crossroads := containers.NewSet[fld.Pos](startPos, finishPos)
	for row, line := range field {
		for col, ch := range line {
			pos := fld.NewPos(row, col)
			if ch == '#' {
				continue
			}
			count := 0
			for _, dir := range fld.DirsSimple {
				npos := pos.Add(dir)
				if field.Inside(npos) && field.Get(npos) != '#' {
					count++
				}
			}
			if count > 2 {
				crossroads.Add(pos)
			}
		}
	}
	type PosDist struct {
		pos  fld.Pos
		dist int
	}
	seen := containers.NewSet[fld.Pos]()
	var adjCrossroads func(pos fld.Pos, dist int) []PosDist
	adjCrossroads = func(pos fld.Pos, dist int) []PosDist {
		seen.Add(pos)
		var res []PosDist
		for _, dir := range fld.DirsSimple {
			npos := pos.Add(dir)
			if field.Inside(npos) && field.Get(npos) != '#' && !seen.Has(npos) {
				if crossroads.Has(npos) {
					// Memoize the connection and stop.
					res = append(res, PosDist{pos: npos, dist: dist + 1})
				} else {
					res = append(res, adjCrossroads(npos, dist+1)...)
				}
			}
		}
		return res
	}
	g := map[fld.Pos][]PosDist{}
	for first := range crossroads {
		clear(seen)
		g[first] = adjCrossroads(first, 0)
		// fmt.Println("Crossroad:", first, " Adjacents:", g[first])
	}
	// fmt.Printf("%v\n", g)

	clear(seen)
	maxDist := 0
	var dfs func(cur fld.Pos, dist int)
	dfs = func(cur fld.Pos, dist int) {
		seen.Add(cur)
		if cur == finishPos {
			if dist > maxDist {
				maxDist = dist
				// fmt.Print(dist, " ")
			}
		} else {
			for _, next := range g[cur] {
				if !seen.Has(next.pos) {
					dfs(next.pos, dist+next.dist)
				}
			}
		}
		seen.Remove(cur)
	}
	dfs(startPos, 0)
	// fmt.Println()

	return maxDist
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
