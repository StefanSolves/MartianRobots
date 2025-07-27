package main

import "testing"



// TestTurnLeft verifies the robot correctly rotates 90 degrees to the left

func TestTurnLeft(t *testing.T) {
	//Setup: Create a robot facing North
	robot := &Robot{Position: Position{X: 0, Y: 0, Orientation: 'N'}}

	// Action: Turn the robot left
	robot.turnLeft()

	// Assert: The robot should now be facing West
	if robot.Position.Orientation != 'W' {
		t.Errorf("Expected orientation 'W', got '%c'", robot.Position.Orientation)
	}
}

// TestTurnRight verifies the robot correctly rotates 90 degrees to the right
func TestTurnRight(t *testing.T) {
	
	// Setup: Create a robot facing North
	robot := &Robot{Position: Position{X: 0, Y: 0, Orientation: 'N'}}

	// Action: Turn the robot right
	robot.turnRight()

	// Assert: Check if the robot is now facing East
	if robot.Position.Orientation != 'E' {
		t.Errorf("Expected orientation @ after turning right from North, but got %c", robot.Position.Orientation)
	}	

}

// NewGrid is a helper for tests, we`ll implement it properly in main.go
func NewGrid(maxX, maxY int) *Grid {
	return &Grid{
		MaxX: maxX,
		MaxY: maxY,
		Scents: make(map[string]bool),
	}
}

// TestMoveForwardAndScents uses table-driven tests to cover multiple scenarios.
func TestMoveForwardAndScents(t *testing.T) {
	testCases := []struct {
		name          string
		grid          *Grid
		startRobot    Robot
		expectedRobot Robot
	}{
		{
			name:          "Normal move within grid",
			grid:          NewGrid(5, 5),
			startRobot:    Robot{Position: Position{X: 1, Y: 1, Orientation: 'N'}},
			expectedRobot: Robot{Position: Position{X: 1, Y: 2, Orientation: 'N'}},
		},
		{
			name:          "Robot gets lost at boundary",
			grid:          NewGrid(5, 3),
			startRobot:    Robot{Position: Position{X: 3, Y: 3, Orientation: 'N'}},
			expectedRobot: Robot{Position: Position{X: 3, Y: 3, Orientation: 'N'}, IsLost: true},
		},
		{
			name:          "Robot ignores move due to existing scent",
			grid:          &Grid{MaxX: 5, MaxY: 3, Scents: map[string]bool{"3,3": true}},
			startRobot:    Robot{Position: Position{X: 3, Y: 3, Orientation: 'N'}},
			expectedRobot: Robot{Position: Position{X: 3, Y: 3, Orientation: 'N'}, IsLost: false},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			robotToTest := tc.startRobot
			// This will fail to compile until we create moveForward
			robotToTest.moveForward(tc.grid)

			if robotToTest.Position != tc.expectedRobot.Position || robotToTest.IsLost != tc.expectedRobot.IsLost {
				t.Errorf("Expected robot %+v, but got %+v", tc.expectedRobot, robotToTest)
			}
		})
	}
}