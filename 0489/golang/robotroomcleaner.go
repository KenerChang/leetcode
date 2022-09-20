package robotroomcleaner

import "fmt"

type Robot struct{}

func (robot *Robot) Move() bool {
	return false
}

func (robot *Robot) TurnLeft() {}

func (robot *Robot) TurnRight() {}

func (robot *Robot) Clean() {}

var (
	// directions: 0: up, 1: right, 2: down, 3: left
	directions = [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
)

type Cell []int

func goBack(robot *Robot) {
	// go back to the previous cell
	// since we use right-hand rule, the previous cell is in the back
	robot.TurnRight()
	robot.TurnRight()
	robot.Move()
	robot.TurnRight()
	robot.TurnRight()
}

func backtrack(robot *Robot, visited map[string]bool, cell Cell, direction int) {
	// clean the cell
	robot.Clean()

	// mark the cell as visited
	key := fmt.Sprintf("%d-%d", cell[0], cell[1])
	visited[key] = true

	// find the next cell in clockwise order
	for i := 0; i < 4; i++ {
		nextDirection := (direction + i) % 4
		nextCell := []int{
			cell[0] + directions[nextDirection][0],
			cell[1] + directions[nextDirection][1],
		}

		nextCellKey := fmt.Sprintf("%d-%d", nextCell[0], nextCell[1])
		if !visited[nextCellKey] && robot.Move() {
			backtrack(robot, visited, nextCell, nextDirection)
			goBack(robot)
		}

		// turn right
		robot.TurnRight()
	}
}

func cleanRoom(robot *Robot) {
	// To solve this problem, we need to keep 2 concepts in mind:
	// 1. constrained programming
	// 2. backtracking
	// Constrained programming means that we have some constraints in the problem,
	// and usually we don't solve constrained problem in optimal way but in a feasible way
	// Backtracking so that we solve the problem recursively, and we can backtrack

	// Combine these 2 concepts, we can use right-hand rule to solve this problem
	// everytime the robot encounters a obstacle, it turn right
	// and the clean task finishes when these is no more cells in backtracking
	visited := map[string]bool{}
	backtrack(robot, visited, []int{0, 0}, 0)
}
