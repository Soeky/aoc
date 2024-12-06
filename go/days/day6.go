package days

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	X, Y int
}

type CoordinateSet map[Coordinate]struct{}

func NewCoordinateSet() CoordinateSet {
	return make(CoordinateSet)
}

func (cs CoordinateSet) Add(coord Coordinate) {
	cs[coord] = struct{}{}
}

func (cs CoordinateSet) Contains(coord Coordinate) bool {
	_, exists := cs[coord]
	return exists
}

func (cs CoordinateSet) Size() int {
	return len(cs)
}

type State struct {
	X, Y      int
	Direction int
}

type StateSet map[State]struct{}

func NewStateSet() StateSet {
	return make(StateSet)
}

func (cs StateSet) Add(coord State) {
	cs[coord] = struct{}{}
}

func (cs StateSet) Contains(coord State) bool {
	_, exists := cs[coord]
	return exists
}

func (cs StateSet) Size() int {
	return len(cs)
}

func Day6() {
	fmt.Println("Solution for Day 6, Part 1:")
	solveDay6Part1()
	fmt.Println("Solution for Day 6, Part 2:")
	solveDay6Part2()
}

func solveDay6Part1() {
	file, err := os.Open("days/6input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var compass [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		compass = append(compass, []rune(line))
	}
	posX := 0
	posY := 0
	for i := 0; i < len(compass); i++ {
		for j, char := range compass[i] {
			if char == '^' {
				posY = i
				posX = j
			}
		}
	}
	fmt.Println("Position X:", posX, "Position Y:", posY)
	visited := NewCoordinateSet()
	direction := 0 // 0 = up, 1 = right, 2 = down, 3 = left
	dx := []int{0, 1, 0, -1}
	dy := []int{-1, 0, 1, 0}

	for posX >= 0 && posX < len(compass[0]) && posY >= 0 && posY < len(compass) {
		visited.Add(Coordinate{X: posX, Y: posY})

		nextX := posX + dx[direction]
		nextY := posY + dy[direction]

		if nextX < 0 || nextX >= len(compass[0]) || nextY < 0 || nextY >= len(compass) {
			break // Wächter verlässt die Karte
		}

		// Prüfen, ob vor dem Wächter ein Hindernis (#) ist
		if compass[nextY][nextX] != '#' {
			// Kein Hindernis, Wächter bewegt sich
			posX, posY = nextX, nextY
		} else {
			// Hindernis: Richtung ändern (90 Grad nach rechts)
			direction = (direction + 1) % 4
		}
	}
	fmt.Println("Ergebnis Waechter:", visited.Size())
}

func solveDay6Part2() {
	file, err := os.Open("days/6input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var compass [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		compass = append(compass, []rune(line))
	}

	var startX, startY int
	found := false
	for i := 0; i < len(compass) && !found; i++ {
		for j, char := range compass[i] {
			if char == '^' {
				startY, startX = i, j
				found = true
				break
			}
		}
	}

	fmt.Println("Position X:", startX, "Position Y:", startY)

	dx := []int{0, 1, 0, -1}
	dy := []int{-1, 0, 1, 0}
	loops := 0

	for i := 0; i < len(compass); i++ {
		for j := 0; j < len(compass[0]); j++ {
			copyA := compass[i][j]
			compass[i][j] = '#'

			if simulateGuard(compass, startX, startY, dx, dy) {
				loops++
			}

			compass[i][j] = copyA
		}
	}

	fmt.Println("Ergebnis Loops:", loops)
}

func simulateGuard(compass [][]rune, startX, startY int, dx, dy []int) bool {
	posX, posY := startX, startY
	direction := 0
	visited := NewStateSet()

	for posX >= 0 && posX < len(compass[0]) && posY >= 0 && posY < len(compass) {
		if visited.Contains(State{X: posX, Y: posY, Direction: direction}) {
			return true
		}
		visited.Add(State{X: posX, Y: posY, Direction: direction})

		nextX := posX + dx[direction]
		nextY := posY + dy[direction]

		if nextX < 0 || nextX >= len(compass[0]) || nextY < 0 || nextY >= len(compass) {
			break
		}

		if compass[nextY][nextX] != '#' {
			posX, posY = nextX, nextY
		} else {
			direction = (direction + 1) % 4
		}
	}
	return false
}
