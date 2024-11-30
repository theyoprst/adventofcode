package fld

type Pos struct {
	Row, Col int
}

var (
	Zero  = Pos{0, 0}
	Right = Pos{0, 1}
	Left  = Pos{0, -1}
	Up    = Pos{-1, 0}
	Down  = Pos{1, 0}

	UpRight   = Up.Add(Right)
	UpLeft    = Up.Add(Left)
	DownRight = Down.Add(Right)
	DownLeft  = Down.Add(Left)

	East  = Right
	West  = Left
	North = Up
	South = Down

	NorthEast = UpRight
	NorthWest = UpLeft
	SouthEast = DownRight
	SouthWest = DownLeft

	DirsSimple = []Pos{Left, Right, Up, Down}
	DirsDiag   = []Pos{UpRight, UpLeft, DownRight, DownLeft}
	DirsAll    = append(DirsSimple, DirsDiag...)
)

func NewPos(row, col int) Pos {
	return Pos{
		Row: row,
		Col: col,
	}
}

func (p Pos) Add(dir Pos) Pos {
	p.Row += dir.Row
	p.Col += dir.Col
	return p
}

func (p Pos) Sub(other Pos) Pos {
	return p.Add(other.Reverse())
}

func (p Pos) Mult(mult int) Pos {
	p.Row *= mult
	p.Col *= mult
	return p
}

func (p Pos) Reverse() Pos {
	return p.Mult(-1)
}
