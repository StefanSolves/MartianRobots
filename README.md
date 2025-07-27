# MartianRobots


Martian Robots Challenge - Go Implementation
This repository contains a command-line application written in Go that solves the Martian Robots programming problem. The solution was developed using a Test-Driven Development (TDD) methodology to ensure correctness and reliability.

How to Run

Prerequisites

Go (version 1.18 or later) installed on your system. ()

An input file (e.g., input.txt) formatted according to the problem description.

5 3
1 1 E
RFRFRFRF
3 2 N
FRRFLLFFRRFLL
0 3 W
LLFFFLFLFL


1. Run the Unit Tests
To verify that the core logic is working correctly, run the built-in test suite:

go test 
Output -> Please see screenshot attached at the bottom

2. Execute the Program
Run the main application from your terminal, providing the path to your input file as a command-line argument:

go run main.go input.txt
Output -> Please see screenshot attached at the bottom

The final positions of the robots will be printed to the console.

Design and Approach
Language: Go

Chosen for its simplicity, strong typing, performance, and excellent standard library, which allows for a robust solution with no external dependencies.

Test-Driven Development (TDD)

The core logic (turning, moving, boundary checks, scent handling) was developed by first writing failing unit tests and then implementing the code to make them pass. This ensures each component is verified before being integrated into the main application.

Data Structures

Grid: A struct holding the world's boundaries and a map[string]bool to efficiently store and look up "scent" locations.

Robot: A struct containing its Position and an IsLost flag to manage its state.

Position: A simple struct to cleanly group a robot's coordinates and orientation.

Input Handling

The program reads from a file specified as a command-line argument. This was chosen as a robust and user-friendly method for handling the multi-line input required by the problem.

<details>
<summary><strong>Click to view TDD Process Screenshots</strong></summary>

The "Red-Green" Cycle
The project was built by following a strict Test-Driven Development cycle.

1. Write a Failing Test (The "Red" Step)

![Screenshot of failing Go tests](screenshots/Failing-test.png)

First, a test is written for a feature that doesn't exist yet. Running go test at this stage results in a compile error or a test failure.

2. Write Code to Pass the Test (The "Green" Step)
![Screenshot of passing Go tests](screenshots/Test-Pass.png)

Next, the simplest possible code is written to make the failing test pass successfully.

This cycle ensures that every component of the application is verified and reliable before moving on.

</details>