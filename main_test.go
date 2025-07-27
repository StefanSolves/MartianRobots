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