package aoc2023

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amra.satara/learning-go/fileparsers"
)

func Day01Part1() int {
	input := fileparsers.ReadLines("inputs2023\\day1.txt")

	sum := 0
	for _, v := range input {
		x := parseNumber(v)
		sum += x
	}
	return sum
}

func Day01Part2() int {
	input := fileparsers.ReadLines("inputs2023\\day1.txt")

	ch := make(chan int)
	for _, v := range input {
		go func(line string) {
			x := parseNumberAndLetters(line)
			ch <- x
		}(v)
	}

	sum := 0
	for i := 0; i < len(input); i++ {
		sum += <-ch
	}

	return sum
}

func parseNumber(input string) int {
	first := -1
	last := 0

	for _, v := range input {
		if v >= rune('0') && v <= rune('9') {
			x, _ := strconv.Atoi(string(v))
			if first == -1 {
				first = x
			}
			last = x
		}
	}

	if first == -1 {
		first = 0
	}

	n := first*10 + last
	return n
}
func parseNumberAndLetters(input string) int {
	result := 0
	n := findFirstNumber(input) + string(findLastNumber(input))
	fmt.Sscanf(n, "%d", &result)
	return result
}

func findFirstNumber(input string) string {
	min := len(input)
	minValue := ""

	for _, v := range listForReplace {
		index := strings.Index(input, v.From)
		if index != -1 && index < min {
			min = index
			minValue = v.To[0:1]
		}
	}
	return minValue
}

func findLastNumber(input string) string {

	max := -1
	maxValue := ""

	for _, v := range listForReplace {

		index := strings.LastIndex(input, v.From)
		if index != -1 && index > max {
			max = index
			length := len(v.To)
			maxValue = v.To[length-1 : length]
		}
	}
	return maxValue
}

type Dataday1 struct {
	From string
	To   string
}

var listForReplace = []Dataday1{
	{From: "ten", To: "10"},
	{From: "eleven", To: "11"},
	{From: "twelve", To: "12"},
	{From: "thirteen", To: "13"},
	{From: "fourteen", To: "14"},
	{From: "fifteen", To: "15"},
	{From: "sixteen", To: "16"},
	{From: "seventeen", To: "17"},
	{From: "eighteen", To: "18"},
	{From: "nineteen", To: "19"},
	{From: "one", To: "1"},
	{From: "two", To: "2"},
	{From: "three", To: "3"},
	{From: "four", To: "4"},
	{From: "five", To: "5"},
	{From: "six", To: "6"},
	{From: "seven", To: "7"},
	{From: "eight", To: "8"},
	{From: "nine", To: "9"},
	{From: "1", To: "1"},
	{From: "2", To: "2"},
	{From: "3", To: "3"},
	{From: "4", To: "4"},
	{From: "5", To: "5"},
	{From: "6", To: "6"},
	{From: "7", To: "7"},
	{From: "8", To: "8"},
	{From: "9", To: "9"},
}
