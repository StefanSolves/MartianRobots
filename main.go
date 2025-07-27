package main






// Position represents the coordinates and orientation on a the grid
type Position struct {
	X, Y int
	Orientation rune
}

// Robot represents a single martian robot
type Robot struct {
	Position Position
	IsLost bool
}


// Grid represents the rectangular surface of Mars
type Grid struct {
	MaxX, MaxY int
	Scents map[string]bool // Scents are positions where robots have been lost as a result of falling off the grid - Scents have been defined in the requirements of the task

}

