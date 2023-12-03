package aoc2023

import (
	"strconv"
	"strings"

	"github.com/amra.satara/learning-go/fileparsers"
)

func Day03Part1() int {
	input := fileparsers.ReadLines("inputs2023\\day3.txt")
	numbers := parseNumberFromMatrix(input)
	result := filterNumbers(input, numbers)

	return Sum(numbers) - Sum(result)
}

func Day03Part2() int {
	input := fileparsers.ReadLines("inputs2023\\day3.txt")
	numbers := parseNumberFromMatrix(input)
	sum := filterNumbersPart2(input, numbers)

	return sum
}

func filterNumbers(input []string, numbers []NumberDay3) []NumberDay3 {
	result := make([]NumberDay3, 0)

	sizeY := len(input)
	sizeX := len(input[0])

	for _, number := range numbers {
		var line1 string
		var part string
		minx := max(0, number.MinX-1)
		maxx := min(sizeX, number.MaxX+2)
		if number.Y > 0 {
			line1 = input[number.Y-1]
			part = line1[minx:maxx]
			if len(strings.ReplaceAll(part, ".", "")) > 0 {
				continue
			}
		}
		if number.Y < sizeY-1 {

			line1 = input[number.Y+1]
			part = line1[minx:maxx]
			if len(strings.ReplaceAll(part, ".", "")) > 0 {
				continue
			}
		}
		before := input[number.Y][max(0, number.MinX-1):number.MinX]
		after := input[number.Y][min(sizeX, number.MaxX+1):min(sizeX, number.MaxX+2)]

		if len(before) > 0 && before != "." {
			continue
		}
		if len(after) > 0 && after != "." {
			continue
		}

		result = append(result, number)

	}
	return result
}

func parseNumberFromMatrix(input []string) []NumberDay3 {
	numbers := make([]NumberDay3, 0)

	for i, line := range input {
		b := strings.Builder{}
		minx := 0
		j := 0
		for _, v := range line {
			if v >= rune('0') && v <= rune('9') {
				b.WriteRune(v)
			} else if b.Len() > 0 {
				n, _ := strconv.Atoi(b.String())
				b = strings.Builder{}
				numbers = append(numbers, NumberDay3{
					Y:     i,
					Value: n,
					MinX:  minx,
					MaxX:  j - 1,
				})
				minx = j + 1

			} else {
				minx++
			}
			j++
		}

		if b.Len() > 0 {
			n, _ := strconv.Atoi(b.String())
			numbers = append(numbers, NumberDay3{
				Y:     i,
				Value: n,
				MinX:  minx,
				MaxX:  j - 1,
			})
		}
	}
	return numbers
}

func filterNumbersPart2(input []string, numbers []NumberDay3) int {
	sum := 0
	for y, line := range input {
		for x, v := range line {
			if v == '*' {
				neighbours := findNeighbours(numbers, x, y)

				if len(neighbours) > 1 {
					sum += Multiple(neighbours)
				}
			}
		}
	}
	return sum
}

func findNeighbours(numbers []NumberDay3, x, y int) []NumberDay3 {

	result := make([]NumberDay3, 0)

	for _, number := range numbers {
		if (number.MinX-1) <= (x) && (number.MaxX+1) >= (x) {
			if (number.Y-1) <= y && (number.Y+1) >= y {
				result = append(result, number)
			}
		}
	}

	return result

}

type NumberDay3 struct {
	Value int
	MinX  int
	MaxX  int
	Y     int
}

func Sum(list []NumberDay3) int {
	sum := 0
	for _, v := range list {
		sum += v.Value
	}
	return sum
}

func Multiple(list []NumberDay3) int {
	p := 1
	for _, v := range list {
		p *= v.Value
	}
	return p
}
