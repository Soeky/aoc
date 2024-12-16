package days

import (
	"bufio"
	"fmt"
	"os"
)

func Day12() {
	fmt.Println("Solution for Day 12, Part 1:", solveDay12Part1())
	fmt.Println("Solution for Day 12, Part 2:", solveDay12Part2())
}

type Point struct {
	x, y int
}

var directions = []Point{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func solveDay12Part1() int {
	file, err := os.Open("days/12input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var myMap [][]rune
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		myMap = append(myMap, []rune(line))
	}
	sum = calculateCost(myMap)
	return sum
}

func calculateCost(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var bfs func(int, int, rune) (int, int)
	bfs = func(y, x int, plant rune) (int, int) {
		queue := []Point{{x, y}}
		visited[y][x] = true
		area, perimeter := 0, 0

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			area++
			for _, dir := range directions {
				nx, ny := curr.x+dir.x, curr.y+dir.y
				if nx < 0 || ny < 0 || ny >= rows || nx >= cols || grid[ny][nx] != plant {
					perimeter++
				} else if !visited[ny][nx] {
					visited[ny][nx] = true
					queue = append(queue, Point{nx, ny})
				}
			}

		}

		return area, perimeter
	}

	total := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] {
				area, perimeter := bfs(i, j, grid[i][j])
				total += area * perimeter
			}
		}
	}

	return total
}

func solveDay12Part2() int {
	file, err := os.Open("days/12input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	n, m := len(lines), len(lines[0])
	graph := make(map[Point]rune)
	visited := make(map[Point]bool)

	for i, row := range lines {
		for j, c := range row {
			graph[Point{i, j}] = c
		}
	}
	for i := -1; i <= n; i++ {
		graph[Point{i, -1}] = '#'
		graph[Point{i, m}] = '#'
	}
	for j := -1; j <= m; j++ {
		graph[Point{-1, j}] = '#'
		graph[Point{n, j}] = '#'
	}

	var dfs func(Point, rune, Point) (int, int)
	dfs = func(node Point, color rune, dir Point) (int, int) {
		if graph[node] != color {
			if graph[Point{node.x + dir.y, node.y - dir.x}] == color || graph[Point{node.x - dir.x + dir.y, node.y - dir.y - dir.x}] != color {
				return 0, 1
			}
			return 0, 0
		}
		if visited[node] {
			return 0, 0
		}
		visited[node] = true
		area, sides := 1, 0
		for _, d := range directions {
			a, s := dfs(Point{node.x + d.x, node.y + d.y}, color, d)
			area += a
			sides += s
		}
		return area, sides
	}

	sum := 0
	for node, color := range graph {
		if color != '#' && !visited[node] {
			area, sides := dfs(node, color, Point{0, 1})
			sum += area * sides
		}
	}

	return sum
}
