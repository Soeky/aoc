package days

import (
	"aoc/utils"
	"bufio"
	"fmt"
	"os"
)

func Day8() {
	fmt.Println("Solution for Day 8, Part 1:")
	solveDay8Part1()
	fmt.Println("Solution for Day 8, Part 2:")
	solveDay8Part2()
}

func solveDay8Part1() {
	file, err := os.Open("days/8input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Map von Zeichen zu Koordinaten-Slice
	characterCoords := make(map[rune][]utils.Coordinate)
	resSet := utils.NewCoordinateSet()
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
				coordinate := utils.Coordinate{X: x, Y: y}
				characterCoords[char] = append(characterCoords[char], coordinate)
			}
		}
		y++
	}

	for _, coords := range characterCoords {
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				first := coords[i]
				second := coords[j]
				m := utils.Coordinate{X: first.X - second.X, Y: first.Y - second.Y}
				firstNode := utils.Coordinate{X: first.X + m.X, Y: first.Y + m.Y}
				secondNode := utils.Coordinate{X: second.X - m.X, Y: second.Y - m.Y}
				if firstNode.X >= 0 && firstNode.X < y && firstNode.Y >= 0 && firstNode.Y < y {
					resSet.Add(firstNode)
				}
				if secondNode.X >= 0 && secondNode.X < y && secondNode.Y >= 0 && secondNode.Y < y {
					resSet.Add(secondNode)
				}

			}
		}
	}
	fmt.Println("Ergebnis 1:", resSet.Size())
}

func solveDay8Part2() {
	file, err := os.Open("days/8input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Map von Zeichen zu Koordinaten-Slice
	characterCoords := make(map[rune][]utils.Coordinate)
	resSet := utils.NewCoordinateSet()
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
				coordinate := utils.Coordinate{X: x, Y: y}
				characterCoords[char] = append(characterCoords[char], coordinate)
			}
		}
		y++
	}

	for _, coords := range characterCoords {
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				first := coords[i]
				second := coords[j]
				m := utils.Coordinate{X: first.X - second.X, Y: first.Y - second.Y}
				firstNode := utils.Coordinate{X: first.X + m.X, Y: first.Y + m.Y}
				secondNode := utils.Coordinate{X: second.X - m.X, Y: second.Y - m.Y}
				resSet.Add(first)
				resSet.Add(second)
				for firstNode.X >= 0 && firstNode.X < y && firstNode.Y >= 0 && firstNode.Y < y {
					resSet.Add(firstNode)
					firstNode = utils.Coordinate{X: firstNode.X + m.X, Y: firstNode.Y + m.Y}
				}
				for secondNode.X >= 0 && secondNode.X < y && secondNode.Y >= 0 && secondNode.Y < y {
					resSet.Add(secondNode)
					secondNode = utils.Coordinate{X: secondNode.X - m.X, Y: secondNode.Y - m.Y}
				}

			}
		}
	}
	fmt.Println("Ergebnis 1:", resSet.Size())
}
