package days

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	X, Y int
}

func Day10() {
	fmt.Println("Solution for Day 10, Part 1:", solveDay10Part1())
	fmt.Println("Solution for Day 10, Part 2:", solveDay10Part2())
}

func solveDay10Part1() int {
	file, err := os.Open("days/10input.txt")
	if err != nil {
		fmt.Println("error opening file:", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var nums [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		nums = append(nums, []rune(line))
	}

	rows := len(nums)
	if rows == 0 {
		return 0
	}
	cols := len(nums[0])

	var trailheads []Pos
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if nums[y][x] == '0' {
				trailheads = append(trailheads, Pos{X: x, Y: y})
			}
		}
	}

	totalScore := 0
	for _, start := range trailheads {
		score := bfsCountNines(nums, start)
		totalScore += score
	}

	return totalScore
}

func bfsCountNines(nums [][]rune, start Pos) int {
	rows := len(nums)
	cols := len(nums[0])
	visited := make(map[Pos]bool)
	queue := []Pos{start}
	ninePositions := make(map[Pos]bool)
	for len(queue) > 0 {
		nextQueue := []Pos{}
		for _, pos := range queue {
			if visited[pos] {
				continue
			}
			visited[pos] = true

			currentVal := nums[pos.Y][pos.X]

			if currentVal == '9' {
				ninePositions[pos] = true
			} else {
				// Otherwise, attempt to move to neighbors with val+1
				targetVal := currentVal + 1

				// Left
				if pos.X > 0 && nums[pos.Y][pos.X-1] == targetVal {
					nextQueue = append(nextQueue, Pos{X: pos.X - 1, Y: pos.Y})
				}
				// Right
				if pos.X < cols-1 && nums[pos.Y][pos.X+1] == targetVal {
					nextQueue = append(nextQueue, Pos{X: pos.X + 1, Y: pos.Y})
				}
				// Up
				if pos.Y > 0 && nums[pos.Y-1][pos.X] == targetVal {
					nextQueue = append(nextQueue, Pos{X: pos.X, Y: pos.Y - 1})
				}
				// Down
				if pos.Y < rows-1 && nums[pos.Y+1][pos.X] == targetVal {
					nextQueue = append(nextQueue, Pos{X: pos.X, Y: pos.Y + 1})
				}
			}
		}
		queue = nextQueue
	}

	return len(ninePositions)
}

func solveDay10Part2() int {
	// Part Two: We want the rating of each trailhead, which is the number of distinct
	// trails (paths) from that trailhead (`0`) to any `9`.
	// We'll use DFS + memoization to count paths.

	file, err := os.Open("days/10input.txt")
	if err != nil {
		fmt.Println("error opening file:", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var nums [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		nums = append(nums, []rune(line))
	}

	rows := len(nums)
	if rows == 0 {
		return 0
	}
	cols := len(nums[0])

	// Identify all trailheads (0-cells)
	var trailheads []Pos
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if nums[y][x] == '0' {
				trailheads = append(trailheads, Pos{X: x, Y: y})
			}
		}
	}

	// Memo for DFS: memo[y][x] = number of distinct trails from (x,y) to '9'
	memo := make([][]int, rows)
	for i := range memo {
		memo[i] = make([]int, cols)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 means uncomputed
		}
	}

	// Directions
	dirs := []Pos{{X: 1, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1}}

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if memo[y][x] != -1 {
			return memo[y][x]
		}

		val := nums[y][x]
		// If at '9', exactly one trail (the endpoint itself).
		if val == '9' {
			memo[y][x] = 1
			return 1
		}

		// Otherwise, sum ways from all neighbors that are val+1
		ways := 0
		targetVal := val + 1
		for _, d := range dirs {
			nx, ny := x+d.X, y+d.Y
			if nx >= 0 && nx < cols && ny >= 0 && ny < rows {
				if nums[ny][nx] == targetVal {
					ways += dfs(nx, ny)
				}
			}
		}

		memo[y][x] = ways
		return ways
	}

	// Sum the ratings for all trailheads
	sumRatings := 0
	for _, start := range trailheads {
		sumRatings += dfs(start.X, start.Y)
	}

	return sumRatings
}
