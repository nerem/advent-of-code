package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		left := -1
		leftIndex := -1
		right := -1
		rightIndex := -1
		calibrationValue := 0
		line := scanner.Text()
		size := len(line)

		for i := 0; i < len(line); i++ {
			if left == -1 && unicode.IsNumber(rune(line[i])) {
				left = int(line[i] - '0')
				leftIndex = i
			}

			if right == -1 && unicode.IsNumber(rune(line[size-i-1])) {
				right = int(line[size-i-1] - '0')
				rightIndex = size - i - 1
			}

			if left > -1 && right > -1 {
				break
			}
		}

		considerWrittenNumbers(line, leftIndex, rightIndex, &left, &right)

		calibrationValue = left*10 + right

		sum += calibrationValue
	}

	fmt.Println(sum)
}

type writtenNumber struct {
	number string
	value  int
}

func considerWrittenNumbers(line string, leftIndex, rightIndex int, left *int, right *int) {
	writtenNumbers := []writtenNumber{
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
	}

	for _, n := range writtenNumbers {
		if i := strings.Index(line, n.number); i < leftIndex && strings.Contains(line, n.number) {
			*left = n.value
			leftIndex = i
		}
		if i := strings.LastIndex(line, n.number); i > rightIndex && strings.Contains(line, n.number) {
			*right = n.value
			rightIndex = i
		}
	}

}
