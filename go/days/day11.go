package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day11() {
	fmt.Println("Solution for Day 11, Part 1:", solveDay11Part1())
	fmt.Println("Solution for Day 11, Part 2:", solveDay11Part2())
}

func solveDay11Part1() int {
	file, err := os.Open("days/11input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	var res []int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		for _, part := range parts {
			number, _ := strconv.Atoi(part)
			res = append(res, getAllStonesForNumber(number, 25)...)
			sum += getAllStonesForNumberFast(number, 25)
		}
	}
	fmt.Println("Size of res =", len(res))
	fmt.Println("Size of sum", sum)
	return sum
}

func getAllStonesForNumber(num int, ran int) []int {
	var res []int
	count := 0
	res = append(res, num)
	cache := make(map[int][]int)

	for count < ran {
		var tmp []int
		for _, number := range res {
			// Falls das Ergebnis bereits in der Map ist
			if val, exists := cache[number]; exists {
				tmp = append(tmp, val...)
				continue
			}

			// Berechnung, falls noch nicht gecached
			var newStones []int
			if number == 0 {
				newStones = []int{1}
			} else {
				digits := int(math.Log10(float64(number)) + 1) // Anzahl der Ziffern berechnen
				if digits%2 == 0 {
					half := int(math.Pow(10, float64(digits/2))) // 10^(digits/2)
					left := number / half                        // Linke Hälfte
					right := number % half                       // Rechte Hälfte
					newStones = []int{left, right}
				} else {
					newStones = []int{number * 2024}
				}
			}

			// Speichere das Ergebnis in der Map
			cache[number] = newStones

			// Füge die berechneten Steine in tmp ein
			tmp = append(tmp, newStones...)
		}
		res = tmp
		count++
	}
	return res
}

func solveDay11Part2() int {
	file, err := os.Open("days/11input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	var res []int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		for _, part := range parts {
			number, _ := strconv.Atoi(part)
			sum += getAllStonesForNumberFast(number, 75)
		}
	}
	fmt.Println("Size of res =", len(res))
	return sum
}

func getAllStonesForNumberFast(num int, loop int) int {
	sum := 0
	myMap := make(map[int]int)
	myMap[num] = 1
	count := 0
	for count < loop {
		tmp := make(map[int]int)
		for key, value := range myMap {
			if key == 0 {
				tmp[1] += value
				continue
			}
			digits := int(math.Log10(float64(key)) + 1)
			if digits%2 == 0 {
				half := int(math.Pow(10, float64(digits/2)))
				left := key / half
				right := key % half
				tmp[left] += value
				tmp[right] += value
				continue
			}
			tmp[key*2024] += value
		}
		myMap = tmp
		count++
	}
	for _, value := range myMap {
		sum += value
	}

	return sum

}
