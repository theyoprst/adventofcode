package main

import (
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

type Point3D struct {
	x, y, z int
}

type Point2D struct {
	x, y int
}

type Brick struct {
	first Point3D
	last  Point3D
	below map[int]bool
	above map[int]bool
}

func SolvePart1(lines []string) any {
	bricks := parseAndFallBricks(lines)

	supporting := map[int]bool{}
	for _, brick := range bricks {
		if len(brick.below) == 1 {
			for belowI := range brick.below {
				supporting[belowI] = true
			}
		}
	}
	return len(bricks) - len(supporting)
}

func SolvePart2(lines []string) any {
	bricks := parseAndFallBricks(lines)

	var removedI int
	var dfs func(curI int)
	var seen map[int]bool
	dfs = func(curI int) {
		if seen[curI] || curI == removedI {
			return
		}
		seen[curI] = true
		for nextI := range bricks[curI].above {
			dfs(nextI)
		}
	}
	ans := 0
	for removedI = 1; removedI < len(bricks); removedI++ {
		seen = map[int]bool{}
		dfs(0)
		// All the bricks which are reachable from the ground plus removed one are not falling.
		// All the rest are falling.
		ans += len(bricks) - 1 - len(seen)
	}
	return ans
}

func parseAndFallBricks(lines []string) []Brick {
	bricks := []Brick{
		{above: map[int]bool{}}, // Append virtual brick for the group first.
	}
	type TopItem struct {
		z      int
		brickI int
	}
	topView := map[Point2D]TopItem{}

	for _, line := range lines {
		nn := aoc.Ints(line)
		brick := Brick{
			first: Point3D{x: nn[0], y: nn[1], z: nn[2]},
			last:  Point3D{x: nn[3], y: nn[4], z: nn[5]},
			below: map[int]bool{},
			above: map[int]bool{},
		}
		bricks = append(bricks, brick)
		must.LessOrEqual(brick.first.x, brick.last.x)
	}
	slices.SortFunc(bricks, func(a, b Brick) int {
		return a.first.z - b.first.z
	})

	// Fall
	for brickI := 1; brickI < len(bricks); brickI++ { // Skip brickI=0 which is ground.
		brick := bricks[brickI]
		maxZ := 0
		for x := brick.first.x; x <= brick.last.x; x++ {
			for y := brick.first.y; y <= brick.last.y; y++ {
				maxZ = max(maxZ, topView[Point2D{x, y}].z)
			}
		}
		must.Less(maxZ, brick.first.z)
		diff := brick.first.z - (maxZ + 1)
		brick.first.z -= diff
		brick.last.z -= diff

		// Update top view and find bricks below the current one.
		now := TopItem{
			z:      brick.last.z,
			brickI: brickI,
		}
		for x := brick.first.x; x <= brick.last.x; x++ {
			for y := brick.first.y; y <= brick.last.y; y++ {
				before := topView[Point2D{x, y}]
				topView[Point2D{x, y}] = now
				if brick.first.z == before.z+1 {
					brick.below[before.brickI] = true
					bricks[before.brickI].above[brickI] = true
				}
			}
		}
		bricks[brickI] = brick
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
