package fld

type Pos struct {
	Row, Col int
}

var (
	Zero  = Pos{0, 0}  // Zero value for Pos
	Right = Pos{0, 1}  // Right is a Pos direction pointing to the right
	Left  = Pos{0, -1} // Left is a Pos direction pointing to the left
	Up    = Pos{-1, 0} // Up is a Pos direction pointing up
	Down  = Pos{1, 0}  // Down is a Pos direction pointing down

	UpRight   = Up.Add(Right)   // UpRight is a Pos direction pointing up and right
	UpLeft    = Up.Add(Left)    // UpLeft is a Pos direction pointing up and left
	DownRight = Down.Add(Right) // DownRight is a Pos direction pointing down and right
	DownLeft  = Down.Add(Left)  // DownLeft is a Pos direction pointing down and left

	East  = Right // East  = Right
	West  = Left  // West  = Left
	North = Up    // North = Up
	South = Down  // South = Down

	NorthEast = UpRight   // NorthEast = UpRight
	NorthWest = UpLeft    // NorthWest = UpLeft
	SouthEast = DownRight // SouthEast = DownRight
	SouthWest = DownLeft  // SouthWest = DownLeft

	DirsSimple = []Pos{Left, Right, Up, Down}                // DirsSimple is a slice of 4 simple directions: Left, Right, Up, Down.
	DirsDiag   = []Pos{UpRight, UpLeft, DownRight, DownLeft} // DirsDiag is a slice of 4 diagonal directions: UpRight, UpLeft, DownRight, DownLeft.
	DirsAll    = append(DirsSimple, DirsDiag...)             // DirsAll is a slice of 8 simple and diagonal directions: Left, Right, Up, Down, UpRight, UpLeft, DownRight, DownLeft.
)

// NewPos returns a new Pos with given row and column.
func NewPos(row, col int) Pos {
	return Pos{
		Row: row,
		Col: col,
	}
}

// Add returns the sum of two Pos.
func (p Pos) Add(dir Pos) Pos {
	p.Row += dir.Row
	p.Col += dir.Col
	return p
}

// Sub returns the difference of two Pos.
func (p Pos) Sub(other Pos) Pos {
	return p.Add(other.Reverse())
}

// Mult returns the Pos multiplied by a scalar.
func (p Pos) Mult(mult int) Pos {
	p.Row *= mult
	p.Col *= mult
	return p
}

// Reverse returns the Pos with reversed sign.
func (p Pos) Reverse() Pos {
	return p.Mult(-1)
}
