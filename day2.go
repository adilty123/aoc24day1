package main

import (
	"aoc2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func day2() {

	lines := utils.ReadLines("inputs/02.txt")

	var safeCount int = 0

	for i := 0; i < len(lines); i++ {
		if IsSafe(lines[i]) {
			safeCount += 1
		}
	}
	println(safeCount)
}

func IsSafe(line string) bool {

	values := ConvertToArray(line)
	n := len(values)

	var result bool = true
	f := values[0]
	s := values[1]

	IsDecrease := IsDecrease(f, s)
	//fmt.Println(IsDecrease)

	for i := 0; i < n-1; i++ {
		first := values[i]
		second := values[i+1]

		if !IsDecrease {
			if IncreaseBreaker(first, second) {
				result = false
				break
			}
		} else {
			if DecreaseBreaker(first, second) {
				result = false
				break
			}
		}
	}
	if !result {
		if ProblemDampner(values) {
			result = true
		}
	}

	return result
}

func ProblemDampner(values []int) bool {

	var result bool = false
	n := len(values)

	for i := 0; i < n; i++ {
		newArr := remove(values, i)
		if IsSafeDampner(newArr) {
			result = true
			break
		}
	}

	return result
}

func IsSafeDampner(values []int) bool {
	n := len(values)

	var result bool = true
	f := values[0]
	s := values[1]

	IsDecrease := IsDecrease(f, s)

	for i := 0; i < n-1; i++ {
		first := values[i]
		second := values[i+1]

		if !IsDecrease {
			if IncreaseBreaker(first, second) {
				result = false
				break
			}
		} else {
			if DecreaseBreaker(first, second) {
				result = false
				break
			}
		}
	}

	return result
}

func remove(slice []int, s int) []int {

	newSlice := make([]int, len(slice))
	copy(newSlice, slice)

	return append(newSlice[:s], newSlice[s+1:]...)
}

func ConvertToArray(line string) []int {
	parts := strings.Fields(line)

	// Step 2: Convert the slice of strings into a slice of integers
	var numbers []int
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("Error converting '%s' to an integer: %v\n", part, err)
			continue // Skip invalid values
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func IncreaseBreaker(first int, second int) bool {

	var result bool = false
	val := second - first

	if second < first {
		result = true
	} else if val > 3 {
		result = true
	} else if second == first {
		result = true
	}

	return result
}

func DecreaseBreaker(first int, second int) bool {
	var result bool = false
	val := first - second

	if second > first {
		result = true
	} else if val > 3 {
		result = true
	} else if second == first {
		result = true
	}
	return result
}

func IsDecrease(first int, second int) bool {
	var result bool = false

	if first > second {
		result = true
	}

	return result
}
