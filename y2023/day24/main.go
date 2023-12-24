package main

import (
	"math"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

type Hailstorm struct {
	px, py, pz, vx, vy, vz int
}

func (h Hailstorm) LineXY() Line2D {
	a := h.vy
	b := -h.vx
	c := -b*h.py - a*h.px
	return Line2D{a, b, c}
}

type Line2D struct {
	// a*x + b*y + c = 0
	a, b, c int
}

func Intersect2D(l1, l2 Line2D) (x, y float64) {
	a1 := float64(l1.a)
	b1 := float64(l1.b)
	c1 := float64(l1.c)
	a2 := float64(l2.a)
	b2 := float64(l2.b)
	c2 := float64(l2.c)
	x = (b1*c2 - b2*c1) / (a1*b2 - a2*b1)
	y = -(a2*x + c2) / b2
	return x, y
}

const (
	MinCoord = 200000000000000
	MaxCoord = 400000000000000
	// MinCoord = 7
	// MaxCoord = 27
)

func SolvePart1(lines []string) any {
	var hailstorms []Hailstorm
	for _, line := range lines {
		nn := aoc.Ints(line)
		hailstorms = append(hailstorms, Hailstorm{nn[0], nn[1], nn[2], nn[3], nn[4], nn[5]})
	}
	var ans int
	for i := 0; i < len(hailstorms); i++ {
		h1 := hailstorms[i]
		l1 := h1.LineXY()
		must.NotEqual(h1.vx, 0)
		must.NotEqual(h1.vy, 0)
		for j := i + 1; j < len(hailstorms); j++ {
			h2 := hailstorms[j]
			l2 := h2.LineXY()
			x, y := Intersect2D(l1, l2)
			if math.IsInf(x, 0) || math.IsInf(y, 0) {
				continue
			}
			t1 := (x - float64(h1.px)) / float64(h1.vx)
			t2 := (x - float64(h2.px)) / float64(h2.vx)
			if MinCoord <= x && x <= MaxCoord && MinCoord <= y && y <= MaxCoord && t1 >= 0 && t2 >= 0 {
				ans++
			}
		}
	}
	return ans
}

type Vector3D struct {
	x, y, z float64
}

func (v Vector3D) Scalar(u Vector3D) float64 {
	return v.x*u.x + v.y*u.y + v.z*u.z
}

func (v Vector3D) Cross(u Vector3D) Vector3D {
	return Vector3D{
		x: v.y*u.z - v.z*u.y,
		y: -v.x*u.z + v.z*u.x,
		z: v.x*u.y - v.y*u.x,
	}
}

func (v Vector3D) Add(u Vector3D) Vector3D {
	return Vector3D{
		x: v.x + u.x,
		y: v.y + u.y,
		z: v.z + u.z,
	}
}

func (v Vector3D) Sub(u Vector3D) Vector3D {
	return Vector3D{
		x: v.x - u.x,
		y: v.y - u.y,
		z: v.z - u.z,
	}
}

func SolvePart2(lines []string) any {
	var p []Vector3D
	var v []Vector3D
	for _, line := range lines {
		nn := aoc.Ints(line)
		p = append(p, Vector3D{
			x: float64(nn[0]),
			y: float64(nn[1]),
			z: float64(nn[2]),
		})
		v = append(v, Vector3D{
			x: float64(nn[3]),
			y: float64(nn[4]),
			z: float64(nn[5]),
		})
	}
	p1, p2, p3 := p[0], p[1], p[2]
	v1, v2, v3 := v[0], v[1], v[2]
	// For any point p[i] there is some t[i]: p[i] + t[i] * v[i] = p0 + t[i] * v0.
	// Or p0 - p[i] = t[i] * (v[i] - v0).
	// Which that vectors (p0 - p[i]) and (v0 - v[i]) a collinear.
	// Which means that their cross product is zero. Now we have 3 equastions:
	// 1-3): (p0 - pi) x (v0 - vi) = 0
	// <=>
	// 1-3): p0 x v0 - pi x v0 - p0 x vi + pi x vi = 0
	// Get rid of p0 x v0 resulting in 2 equations:
	// 1): p1 x v0 + p0 x v1 - p1 x v1 = p2 x v0 - p0 x v2 - p2 x v2
	// 2): p1 x v0 + p0 x v1 - p1 x v1 = p3 x v0 - p0 x v3 - p3 x v3
	// <=>
	// 1): (p1 - p2) x v0 + p0 x (v1 - v2) = p1 x v1 - p2 x v2
	// 2): (p1 - p3) x v0 + p0 x (v1 - v3) = p1 x v1 - p3 x v3

	// Equation in vector form:
	// v x a2 + p x b2 = c2
	// v x a3 + p x b3 = c3
	// where (for p0 miltiplier used antisimmetry property)
	a2 := p2.Sub(p1)
	a3 := p3.Sub(p1)
	b2 := v1.Sub(v2)
	b3 := v1.Sub(v3)
	c2 := p1.Cross(v1).Sub(p2.Cross(v2))
	c3 := p1.Cross(v1).Sub(p3.Cross(v3))

	// Order in the system: v.x, v.y, v.z, p.x, p.y, p.z
	system := [][]float64{
		// x: v.y * ai.z - v.z * ai.y + p.y * bi.z - p.z * bi.y = ci.x
		{0, a2.z, -a2.y, 0, b2.z, -b2.y, c2.x},
		{0, a3.z, -a3.y, 0, b3.z, -b3.y, c3.x},
		// y: v.z * ai.x - v.x * ai.z + p.z * bi.x - p.x * bi.z = ci.y
		{-a2.z, 0, a2.x, -b2.z, 0, b2.x, c2.y},
		{-a3.z, 0, a3.x, -b3.z, 0, b3.x, c3.y},
		// z: v.x * ai.y - v.y * ai.x + p.x * bi.y - p.y * bi.x = ci.z
		{a2.y, -a2.x, 0, b2.y, -b2.x, 0, c2.z},
		{a3.y, -a3.x, 0, b3.y, -b3.x, 0, c3.z},
	}
	for v := 0; v < len(system)-1; v++ {
		for u := v; u < len(system); u++ {
			if system[u][v] != 0 {
				system[v], system[u] = system[u], system[v]
				break
			}
		}
		for row := v + 1; row < len(system); row++ {
			k := system[row][v] / system[v][v]
			for col := v; col < len(system[row]); col++ {
				system[row][col] -= k * system[v][col]
			}
		}
	}

	lastCol := len(system)
	for row := len(system) - 1; row >= 0; row-- {
		system[row][lastCol] /= system[row][row]
		system[row][row] = 1.0
		val := system[row][lastCol]
		for up := row - 1; up >= 0; up-- {
			system[up][lastCol] -= val * system[up][row]
			system[up][row] = 0
		}
	}

	var sol []float64
	for _, line := range system {
		sol = append(sol, math.Round(line[lastCol]))
	}
	return int(sol[3] + sol[4] + sol[5])
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
