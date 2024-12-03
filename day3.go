package main

import (
	"aoc2024/utils"
	"regexp"
	"strings"
)

func day3() {

	lines := utils.ReadLines("inputs/03p2.txt")

	var total int
	var result int
	var do bool = true

	for i := 0; i < len(lines); i++ {
		result, do = GetMulDoDont(lines[i], do)
		total += result
	}
	println(result)
}

func GetMul(line string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)

	matches := re.FindAllString(line, -1)

	var result int = 0

	if len(matches) > 0 {
		result = CalculateMul(matches)
	}

	return result
}

func GetMulDoDont(line string, do bool) (int, bool) {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|\bdo\(\)|\bdon't\(\)`)
	matches := re.FindAllString(line, -1)

	var result int = 0

	if len(matches) > 0 {
		result, do = CalculateMulDoDont(matches, do)
	}

	return result, do
}

func CalculateMulDoDont(matches []string, do bool) (int, bool) {
	var result int = 0

	for i := 0; i < len(matches); i++ {
		current := matches[i]

		if current == "do()" {
			do = true
		} else if current == "don't()" {
			do = false
		} else {
			if do {
				values := GetArray(matches[i])

				first := utils.ConvertToInt(values[0])
				second := utils.ConvertToInt(values[1])

				result += first * second
			}
		}
	}
	return result, do
}

func CalculateMul(matches []string) int {

	var result int

	for i := 0; i < len(matches); i++ {
		values := GetArray(matches[i])

		first := utils.ConvertToInt(values[0])
		second := utils.ConvertToInt(values[1])

		result += first * second
	}

	return result
}

func GetArray(value string) []string {

	result := GetXY(value)
	values := strings.Split(result, ",")

	return values
}

func GetXY(input string) string {
	if strings.HasPrefix(input, "mul(") && strings.HasSuffix(input, ")") {
		return input[4 : len(input)-1]
	}
	return ""
}
