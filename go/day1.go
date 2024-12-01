package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	problem2()
}

func problem1() {
	file, err := os.Open("1input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()

	var first []int
	var second []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		for i, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("konnte Zahl nicht lesen")
				return
			}

			if i%2 == 0 {
				first = append(first, number)
			} else {
				second = append(second, number)
			}
		}
	}

	sort.Ints(first)
	sort.Ints(second)
	sum := 0
	for i := 0; i < len(first); i++ {
		diff := first[i] - second[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}
	fmt.Println(sum)
}

func problem2() {
	file, err := os.Open("1input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()

	var first []int
	var second []int
	firstMap := make(map[int]int)
	secondMap := make(map[int]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		for i, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("konnte Zahl nicht lesen")
				return
			}

			if i%2 == 0 {
				first = append(first, number)
				firstMap[number]++
			} else {
				second = append(second, number)
				secondMap[number]++
			}
		}
	}

	sort.Ints(first)
	sort.Ints(second)
	sum := 0
	for _, number := range first {
		if count, exists := secondMap[number]; exists {
			sum += firstMap[number] * number * count
		}
	}
	fmt.Println(sum)
}
