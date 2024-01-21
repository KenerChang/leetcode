package jumpgameii

import (
	"math"
)

func jump(nums []int) int {
	// If we use a tree to represet all the possible pathes, then
	// the path from root to leaf represents a jump path. To find the
	// shortest path, we could use BFS to find the shortest path.
	queue := [][]int{}
	memory := make([]int, len(nums)) // keep reachable height

	target := len(nums) - 1
	for i := 1; i <= target; i++ {
		memory[i] = math.MaxInt
	}

	// init queue
	for step := 1; step <= nums[0]; step++ {
		if step <= target {
			queue = append(queue, []int{step, 1})
			memory[step] = 1
		}
	}

	// do bfs
	for len(queue) > 0 {
		node := queue[0]
		nodeIdx := node[0]
		maxSteps := nums[nodeIdx]

		if nodeIdx == target {
			return node[1]
		}

		// pop node
		queue = queue[1:]

		// put reachable nodes to the queue
		for step := 1; step <= maxSteps; step++ {
			nextNode := nodeIdx + step
			nextHeight := node[1] + 1

			if nextNode <= target && nextHeight < memory[nextNode] {
				queue = append(queue, []int{nodeIdx + step, node[1] + 1})
				memory[nextNode] = nextHeight
			}
		}

	}

	return 0
}
