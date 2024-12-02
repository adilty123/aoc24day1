package main

import (
	"aoc2024/utils"
	"strings"
)

func day2() {

	lines := utils.ReadLines("inputs/02.txt")

	var safeCount int

	for i := 0; i < len(lines); i++ {
		if IsSafe(lines[i]) {
			safeCount += 1
		}
	}
	println(safeCount)
}

func IsSafe(line string) bool {

	values := strings.Split(line, " ")

	n := len(values)

	var result bool = true
	var val int

	for i := 0; i < n-2; i++ {
		first := utils.ConvertToInt(values[i])
		second := utils.ConvertToInt(values[i+1])

		IsDecrease := IsDecrease(first, second)

		val := second - first

		if second <= first {
			result = false
			break
		}

		if val > 3 {
			result = false
			break
		}
	}

	return result
}

func IsDecrease(first int, second int) bool {

}
