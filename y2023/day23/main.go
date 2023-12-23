package main

import (
	"fmt"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1BruteForce(lines []string) any {
	field := fld.NewByteField(lines)
	startPos := fld.NewPos(0, 1)
	finishPos := fld.NewPos(field.Rows()-1, field.Cols()-2)

	seen := containers.NewSet[fld.Pos]()
	var maxPathDFS func(cur fld.Pos, dist int) int
	maxPathDFS = func(cur fld.Pos, dist int) int {
		seen.Add(cur)
		defer seen.Remove(cur)
		if cur == finishPos {
			return dist
		}
		maxPath := -1
		for _, dir := range dirsPart1(field, cur) {
			next := cur.Add(dir)
			if field.Inside(next) && field.Get(next) != '#' && !seen.Has(next) {
				maxPath = max(maxPath, maxPathDFS(next, dist+1))
			}
		}
		return maxPath
	}

	return maxPathDFS(startPos, 0)
}

func SolvePart1(lines []string) any {
	field := fld.NewByteField(lines)

	// Remove forbidden directions, make DFS-based topsort.
	var topsort []fld.Pos
	var dfs func(p fld.Pos)
	seen := containers.NewSet[fld.Pos]()
	dfs = func(p fld.Pos) {
		seen.Add(p)
		for _, dir := range dirsPart1(field, p) {
			np := p.Add(dir)
			if field.Inside(np) && field.Get(np) != '#' && !seen.Has(np) {
				// Hack: forbid moving to `>` from the right and to `v` from the bottom.
				if field.Get(np) == '>' && dir == fld.Left || field.Get(np) == 'v' && dir == fld.Up {
					continue
				}
				dfs(np)
			}
		}
		topsort = append(topsort, p)
	}
	startPos := fld.NewPos(0, 1)
	dfs(startPos)
	slices.Reverse(topsort)

	dp := map[fld.Pos]int{startPos: 0}
	for i := 1; i < len(topsort); i++ { // Skip startPos.
		cur := topsort[i]
		must.True(field.Inside(cur))
		maxLen := -1
		for _, dir := range fld.DirsSimple {
			prev := cur.Add(dir)
			if !field.Inside(prev) {
				continue
			}
			for _, dir := range dirsPart1(field, prev) {
				if prev.Add(dir) == cur { // Only look at cells from which current cell is reachable.
					if val, ok := dp[prev]; ok && val > maxLen {
						maxLen = val
					}
					break
				}
			}
		}
		dp[cur] = maxLen + 1
	}
	finishPos := fld.NewPos(field.Rows()-1, field.Cols()-2)
	return dp[finishPos]
}

func SolvePart2(lines []string) any {
	field := fld.NewByteField(lines)
	startPos := fld.NewPos(0, 1)
	finishPos := fld.NewPos(field.Rows()-1, field.Cols()-2)

	junctions := containers.NewSet[fld.Pos](startPos, finishPos)
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
				junctions.Add(pos)
			}
		}
	}
	type PosDist struct {
		pos  fld.Pos
		dist int
	}
	seen := containers.NewSet[fld.Pos]()
	var adjJunctions func(pos fld.Pos, dist int) []PosDist
	adjJunctions = func(pos fld.Pos, dist int) []PosDist {
		seen.Add(pos)
		var res []PosDist
		for _, dir := range fld.DirsSimple {
			npos := pos.Add(dir)
			if field.Inside(npos) && field.Get(npos) != '#' && !seen.Has(npos) {
				if junctions.Has(npos) {
					// Memoize the connection and stop.
					res = append(res, PosDist{pos: npos, dist: dist + 1})
				} else {
					res = append(res, adjJunctions(npos, dist+1)...)
				}
			}
		}
		return res
	}
	g := map[fld.Pos][]PosDist{}
	for first := range junctions {
		clear(seen)
		g[first] = adjJunctions(first, 0)
		fmt.Println(first, g[first])
	}

	// startPos and finishPos have only one neighbour, change them to speed up brute force.
	saveDist := 0
	removeSinglePath := func(cur fld.Pos) fld.Pos {
		for len(g[cur]) == 1 {
			saveDist += g[cur][0].dist
			next := g[cur][0].pos
			for i, pd := range g[next] {
				if pd.pos == cur {
					g[next] = slices.Delete(g[next], i, i+1)
					break
				}
			}
			cur = next
		}
		return cur
	}

	startPos = removeSinglePath(startPos)
	finishPos = removeSinglePath(finishPos)

	// Now brute-force on junctions graphs.
	clear(seen)
	maxDist := 0
	var dfs func(cur fld.Pos, dist int)
	dfs = func(cur fld.Pos, dist int) {
		seen.Add(cur)
		if cur == finishPos {
			if dist > maxDist {
				maxDist = dist
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

	return saveDist + maxDist
}

func dirsPart1(field fld.ByteField, p fld.Pos) []fld.Pos {
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

var (
	solvers1 = []aoc.Solver{SolvePart1, SolvePart1BruteForce}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
