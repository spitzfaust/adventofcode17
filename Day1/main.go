package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// StringToNumbers converts a string to a slice of ints.
func StringToNumbers(s string) []int {
	var numbers []int
	s = strings.TrimSpace(s)
	for _, r := range s {
		c := string(r)
		number, err := strconv.Atoi(c)
		check(err)
		numbers = append(numbers, number)
	}
	return numbers
}

// FindNumbersForSum finds numbers for sum.
func FindNumbersForSum(numbers []int) []int {
	size := len(numbers)
	var numbersForSum []int
	for i, j := 0, 1; j < size; i, j = i+1, j+1 {
		if numbers[i] == numbers[j] {
			numbersForSum = append(numbersForSum, numbers[i])
		}
	}
	if numbers[0] == numbers[size-1] {
		numbersForSum = append(numbersForSum, numbers[0])
	}
	return numbersForSum
}

// FindNumbersForHalfwaySum finds the numbers for the halfway sum.
func FindNumbersForHalfwaySum(numbers []int) []int {
	size := len(numbers)
	middle := size / 2
	var numbersForSum []int
	for i := 0; i < size; i++ {
		j := (middle + i) % size
		if numbers[i] == numbers[j] {
			numbersForSum = append(numbersForSum, numbers[i])
		}
	}
	return numbersForSum
}

// Sum returns the sum of an int slice.
func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// ReadFileToString reads the contents of a file into a string.
func ReadFileToString(filepath string) string {
	dat, err := ioutil.ReadFile(filepath)
	check(err)
	return string(dat)
}

func main() {
	s := ReadFileToString("./input.txt")
	fmt.Println(s)
	numbers := StringToNumbers(s)

	numbersForSum := FindNumbersForSum(numbers)
	sum := Sum(numbersForSum)
	fmt.Println(sum)

	numbersForHalfwaySum := FindNumbersForHalfwaySum(numbers)
	halfwaySum := Sum(numbersForHalfwaySum)
	fmt.Println(halfwaySum)
}
