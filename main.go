package main

import ("fmt"
)




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


// NewGrid creates and initializes a new Grid.
func NewGrid(maxX, maxY int) *Grid {
	return &Grid{
		MaxX:   maxX,
		MaxY:   maxY,
		Scents: make(map[string]bool),
	}
}

// moveForward moves the robot one step forward in its current direction.
// It also handles boundary checks and scents.
func (r *Robot) moveForward(grid *Grid) {
	// A robot that is already lost cannot move.
	if r.IsLost {
		return
	}

	// Calculate the potential next position.
	nextX, nextY := r.Position.X, r.Position.Y
	switch r.Position.Orientation {
	case 'N':
		nextY++
	case 'E':
		nextX++
	case 'S':
		nextY--
	case 'W':
		nextX--
	}

	// Check if the next move is off the grid.
	if nextX < 0 || nextX > grid.MaxX || nextY < 0 || nextY > grid.MaxY {
		// The robot is about to fall. Check for a scent at its current position.
		scentPosition := fmt.Sprintf("%d,%d", r.Position.X, r.Position.Y)
		
		// If a scent exists at the current spot, the instruction is ignored.
		if grid.Scents[scentPosition] {
			return // Ignore the move.
		}

		// No scent found, so the robot gets lost and leaves one.
		r.IsLost = true
		grid.Scents[scentPosition] = true // Leave a scent.
		return // The robot is lost and does not move from its last position.
	}

	// If the move is valid and on the grid, update the robot's position.
	r.Position.X = nextX
	r.Position.Y = nextY
}

// processInstructions executes a sequence of instructions for the robot.
func (r *Robot) processInstructions(instructions string, grid *Grid) {
	// Loop through each character in the instruction string.
	for _, instruction := range instructions {
		// If the robot is lost, stop processing further instructions.
		if r.IsLost {
			break
		}
		
		// Execute the command based on the instruction character.
		switch instruction {
		case 'L':
			r.turnLeft()
		case 'R':
			r.turnRight()
		case 'F':
			r.moveForward(grid)
		// Note: Any character other than L, R, or F is simply ignored.
		}
	}
}