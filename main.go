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

// turnLeft rotates the robot 90 degrees to the left
func (r *Robot) turnLeft() {
	switch r.Position.Orientation {
	case 'N' :
		r.Position.Orientation = 'W'
	case 'W' :
		r.Position.Orientation = 'S'
	case 'S' :
		r.Position.Orientation = 'E'
	case 'E' :
		r.Position.Orientation = 'N'
	}
}

// turnRight rotates the robot 90 degrees to the right
func (r *Robot) turnRight() {
	switch r.Position.Orientation {
	case 'N':
		r.Position.Orientation = 'E'
	case 'E':
		r.Position.Orientation = 'S'
	case 'S':
		r.Position.Orientation = 'W'
	case 'W':
		r.Position.Orientation = 'N'
	}
}