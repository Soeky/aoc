package days

import (
	"bufio"
	"fmt"
	"os"
)

func Day4() {
	fmt.Println("Solution for Day 4, Part 1:")
	solveDay4Part1()
	fmt.Println("Solution for Day 4, Part 2:")
	solveDay4Part2()
}

func solveDay4Part1() {
	file, err := os.Open("days/4input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var grid [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	sum := countXMAS1(grid)
	fmt.Println("Sum Was:", sum)
}

func countXMAS1(grid [][]rune) int {
	xmas := "XMAS"
	xmasLen := len(xmas)
	dirs := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}
	sum := 0

	inBounds := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for _, dir := range dirs {
				matched := true
				for k := 0; k < xmasLen; k++ {
					ni, nj := i+k*dir[0], j+k*dir[1]
					if !inBounds(ni, nj) || grid[ni][nj] != rune(xmas[k]) {
						matched = false
						break
					}
				}
				if matched {
					sum++
				}
			}
		}
	}

	return sum
}

func solveDay4Part2() {
	file, err := os.Open("days/4input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var grid [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	sum := countXMAS2(grid)
	fmt.Println("Sum Was:", sum)
}
func countXMAS2(grid [][]rune) int {
	sum := 0

	inBounds := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// mitte check
			if grid[i][j] != 'A' {
				continue
			}
			matchFirstDiag := false
			matchSecondDiag := false
			if inBounds(i-1, j-1) && inBounds(i+1, j+1) {
				if (grid[i-1][j-1] == 'M' && grid[i+1][j+1] == 'S') || (grid[i-1][j-1] == 'S' && grid[i+1][j+1] == 'M') {
					matchFirstDiag = true
				}
			}
			if inBounds(i-1, j+1) && inBounds(i+1, j-1) {
				if (grid[i-1][j+1] == 'M' && grid[i+1][j-1] == 'S') || (grid[i-1][j+1] == 'S' && grid[i+1][j-1] == 'M') {
					matchSecondDiag = true
				}
			}
			if matchFirstDiag && matchSecondDiag {
				sum++
			}
		}
	}

	return sum
}
