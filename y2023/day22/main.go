package main

import (
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

type Point3 struct {
	x, y, z int
}

type Point2 struct {
	x, y int
}

type Brick struct {
	first, last Point3
}

type Brick2 struct {
	first, last Point3
	below       map[int]bool
	above       map[int]bool
}

func SolvePart1(lines []string) any {
	bricks := parseAndFallBricks(lines)

	important := map[int]bool{}
	for _, brick := range bricks {
		if len(brick.below) == 1 {
			for brickI := range brick.below {
				important[brickI] = true
			}
		}
	}
	return len(bricks) - len(important)
}

func SolvePart2(lines []string) any {
	var bricks []Brick
	type TopItem struct {
		z     int
		brick int
	}
	bottom := Brick{
		first: Point3{0, 0, 0},
		last:  Point3{0, 0, 0},
	}
	top := map[Point2]TopItem{}

	for _, line := range lines {
		nn := aoc.Ints(line)
		brick := Brick{
			first: Point3{x: nn[0], y: nn[1], z: nn[2]},
			last:  Point3{x: nn[3], y: nn[4], z: nn[5]},
		}
		bricks = append(bricks, brick)
		must.LessOrEqual(brick.first.x, brick.last.x)
		bottom.first.x = min(bottom.first.x, brick.first.x)
		bottom.last.x = max(bottom.last.x, brick.last.x)
		bottom.first.y = min(bottom.first.y, brick.first.y)
		bottom.last.y = max(bottom.last.y, brick.last.y)
	}
	bricks = append(bricks, bottom)
	slices.SortFunc(bricks, func(a, b Brick) int {
		return a.first.z - b.first.z
	})
	g := map[int]map[int]bool{0: {}}
	below := map[int]map[int]bool{}
	above := map[int]map[int]bool{0: {}}
	for i := 1; i < len(bricks); i++ {
		brick := bricks[i]
		maxZ := 0
		for x := brick.first.x; x <= brick.last.x; x++ {
			for y := brick.first.y; y <= brick.last.y; y++ {
				maxZ = max(maxZ, top[Point2{x, y}].z)
			}
		}
		must.Less(maxZ, brick.first.z)
		diff := brick.first.z - (maxZ + 1)
		brick.first.z -= diff
		brick.last.z -= diff
		bricks[i] = brick
		g[i] = map[int]bool{}
		below[i] = map[int]bool{}
		above[i] = map[int]bool{}
		for x := brick.first.x; x <= brick.last.x; x++ {
			for y := brick.first.y; y <= brick.last.y; y++ {
				before := top[Point2{x, y}]
				now := TopItem{
					z:     brick.last.z,
					brick: i,
				}
				top[Point2{x, y}] = now
				if brick.first.z == before.z+1 {
					g[before.brick][i] = true
					g[i][before.brick] = true
					below[i][before.brick] = true
					above[before.brick][i] = true
				}
			}
		}
	}
	timer := 0
	seen := map[int]bool{}
	tin := map[int]int{}
	fup := map[int]int{}
	cutpoints := map[int]bool{}
	var dfs func(i int, prev int)
	dfs = func(cur int, prev int) {
		timer++
		seen[cur] = true
		tin[cur] = timer
		fup[cur] = timer
		children := 0
		for to := range g[cur] {
			if to == prev {
				continue
			}
			if seen[to] {
				fup[cur] = min(fup[cur], tin[to])
			} else {
				dfs(to, cur)
				fup[cur] = min(fup[cur], fup[to])
				if fup[to] >= tin[cur] && prev != -1 {
					cutpoints[cur] = true
				}
				children++
			}
		}
		if prev == -1 && children > 1 {
			cutpoints[cur] = true
		}
	}
	dfs(0, -1)
	// return len(bricks) - len(cutpoints) - 1

	var removed int
	var dfs2 func(cur int)
	var seen2 map[int]bool
	dfs2 = func(cur int) {
		if seen2[cur] || cur == removed {
			return
		}
		seen2[cur] = true
		for next := range above[cur] {
			dfs2(next)
		}
	}
	ans := 0
	for removed = 1; removed < len(bricks); removed++ {
		seen2 = map[int]bool{}
		dfs2(0)
		if len(seen2) < len(bricks)-1 {
			ans += len(bricks) - 1 - len(seen2)
		}
	}
	return ans
}

func parseAndFallBricks(lines []string) []Brick2 {
	bricks := []Brick2{
		{above: map[int]bool{}}, // Append virtual brick for the group first.
	}
	type TopItem struct {
		z     int
		brick int
	}
	top := map[Point2]TopItem{}

	for _, line := range lines {
		nn := aoc.Ints(line)
		brick := Brick2{
			first: Point3{x: nn[0], y: nn[1], z: nn[2]},
			last:  Point3{x: nn[3], y: nn[4], z: nn[5]},
			below: map[int]bool{},
			above: map[int]bool{},
		}
		bricks = append(bricks, brick)
		must.LessOrEqual(brick.first.x, brick.last.x)
	}
	slices.SortFunc(bricks, func(a, b Brick2) int {
		return a.first.z - b.first.z
	})

	// Fall
	for i := 1; i < len(bricks); i++ { // Skip i=0 which is ground.
		brick := bricks[i]
		maxZ := 0
		for x := brick.first.x; x <= brick.last.x; x++ {
			for y := brick.first.y; y <= brick.last.y; y++ {
				maxZ = max(maxZ, top[Point2{x, y}].z)
			}
		}
		must.Less(maxZ, brick.first.z)
		diff := brick.first.z - (maxZ + 1)
		brick.first.z -= diff
		brick.last.z -= diff
		for x := brick.first.x; x <= brick.last.x; x++ {
			for y := brick.first.y; y <= brick.last.y; y++ {
				before := top[Point2{x, y}]
				now := TopItem{
					z:     brick.last.z,
					brick: i,
				}
				top[Point2{x, y}] = now
				if brick.first.z == before.z+1 {
					brick.below[before.brick] = true
					bricks[before.brick].above[i] = true
				}
			}
		}
		bricks[i] = brick
	}
	return bricks
}

var (
	solvers1 []aoc.Solver = []aoc.Solver{SolvePart1}
	solvers2 []aoc.Solver = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
