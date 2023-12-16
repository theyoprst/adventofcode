package fld

type Pos struct {
	Row, Col int
}

var (
	Right = Pos{0, 1}
	Left  = Pos{0, -1}
	Up    = Pos{-1, 0}
	Down  = Pos{1, 0}

	East  = Right
	West  = Left
	North = Up
	South = Down
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

func (p Pos) Mult(mult int) Pos {
	p.Row *= mult
	p.Col *= mult
	return p
}
