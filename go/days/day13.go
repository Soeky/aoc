package days

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	Ax, Ay int
	Bx, By int
	PrizeX int
	PrizeY int
}

func Day13() {
	fmt.Println("Solution for Day 13, Part 1:", solveDay13Part1())
	fmt.Println("Solution for Day 13, Part 2:", solveDay13Part2())
}

func solveDay13Part1() int {
	file, err := os.Open("days/13input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//	var machines []Machine
	sum := 0
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Button") || strings.Contains(line, "Prize") {
			lines = append(lines, line)
		}
	}
	sum = evaluateMachines(lines)

	return sum
}

func solveLinearSystem(Ax, Ay, Bx, By, Px, Py int) (int, int, error) {
	// Coefficients matrix
	det := Ax*By - Ay*Bx
	if det == 0 {
		return 0, 0, errors.New("no solution (determinant is zero)")
	}

	// Solve for a and b
	a := (Px*By - Py*Bx) / det
	b := (Ax*Py - Ay*Px) / det

	// Check if the solution is valid (non-negative integers)
	if a < 0 || b < 0 || Px != a*Ax+b*Bx || Py != a*Ay+b*By {
		return 0, 0, errors.New("no valid solution")
	}

	return a, b, nil
}

func evaluateMachines(lines []string) int {
	totalTokens := 0
	var machines []Machine
	re1 := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
	re2 := regexp.MustCompile(`X=(\d+), Y=(\d+)`)

	for i := 0; i < len(lines); i += 3 {
		matches := re1.FindStringSubmatch(lines[i])
		Ax, _ := strconv.Atoi(matches[1])
		Ay, _ := strconv.Atoi(matches[2])
		matches = re1.FindStringSubmatch(lines[i+1])
		Bx, _ := strconv.Atoi(matches[1])
		By, _ := strconv.Atoi(matches[2])
		matches = re2.FindStringSubmatch(lines[i+2])
		PrizeX, _ := strconv.Atoi(matches[1])
		PrizeY, _ := strconv.Atoi(matches[2])
		machine := Machine{Ax: Ax, Ay: Ay, Bx: Bx, By: By, PrizeX: PrizeX, PrizeY: PrizeY}
		machines = append(machines, machine)
	}

	for _, machine := range machines {
		a, b, err := solveLinearSystem(machine.Ax, machine.Ay, machine.Bx, machine.By, machine.PrizeX, machine.PrizeY)
		if err == nil {
			// Calculate the cost for this machine
			cost := a*3 + b*1
			totalTokens += cost
		} else {
		}
	}

	return totalTokens
}
func solveDay13Part2() int {
	file, err := os.Open("days/13input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//	var machines []Machine
	sum := 0
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Button") || strings.Contains(line, "Prize") {
			lines = append(lines, line)
		}
	}
	sum = evaluateMachines2(lines)

	return sum
}
func evaluateMachines2(lines []string) int {
	totalTokens := 0
	var machines []Machine
	re1 := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
	re2 := regexp.MustCompile(`X=(\d+), Y=(\d+)`)

	for i := 0; i < len(lines); i += 3 {
		matches := re1.FindStringSubmatch(lines[i])
		Ax, _ := strconv.Atoi(matches[1])
		Ay, _ := strconv.Atoi(matches[2])
		matches = re1.FindStringSubmatch(lines[i+1])
		Bx, _ := strconv.Atoi(matches[1])
		By, _ := strconv.Atoi(matches[2])
		matches = re2.FindStringSubmatch(lines[i+2])
		PrizeX, _ := strconv.Atoi(matches[1])
		PrizeX += 10000000000000
		PrizeY, _ := strconv.Atoi(matches[2])
		PrizeY += 10000000000000
		machine := Machine{Ax: Ax, Ay: Ay, Bx: Bx, By: By, PrizeX: PrizeX, PrizeY: PrizeY}
		machines = append(machines, machine)
	}

	for _, machine := range machines {
		a, b, err := solveLinearSystem(machine.Ax, machine.Ay, machine.Bx, machine.By, machine.PrizeX, machine.PrizeY)
		if err == nil {
			// Calculate the cost for this machine
			cost := a*3 + b*1
			totalTokens += cost
		} else {
		}
	}

	return totalTokens
}
