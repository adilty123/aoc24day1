package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var leftList []int
var rightList []int

func main() {
	day3()
}

func day1() {
	Read("input.txt")

	Sort(leftList)
	Sort(rightList)

	FindDiff()
	findSimilarity()

}

func findSimilarity() {
	var result int
	leftLen := len(leftList)
	rightLen := len(rightList)

	for i := 0; i < leftLen; i++ {
		var occurances int
		left := leftList[i]

		for k := 0; k < rightLen; k++ {
			right := rightList[k]
			if left == right {
				occurances += 1
			}
		}

		result += left * occurances
	}

	println(result)
}

func FindDiff() {
	var result int
	n := len(leftList)

	for i := 0; i < n; i++ {
		left := leftList[i]
		right := rightList[i]

		if left < right {
			result += right - left
		} else {
			result += left - right
		}
	}

	println(result)
}

func Sort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			// Swap if the current element is smaller than the next element
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func Read(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			s := string(line)
			replaced_string := strings.Replace(s, " ", ",", 1)
			new := strings.ReplaceAll(replaced_string, " ", "")
			values := strings.Split(new, ",")

			left, err := strconv.Atoi(values[0])
			if err != nil {
				// ... handle error
				panic(err)
			}

			right, err := strconv.Atoi(values[1])
			if err != nil {
				// ... handle error
				panic(err)
			}

			leftList = append(leftList, left)
			rightList = append(rightList, right)
		}
		if err != nil {
			break
		}
	}
}
