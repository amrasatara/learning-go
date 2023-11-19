package aoc2023

import (
	"fmt"

	"github.com/amra.satara/learning-go/fileparsers"
	"github.com/amra.satara/learning-go/positions"
)

func Day01Part1() int {
	input := fileparsers.ReadLines("inputs2023\\day1.txt")
	partsI := fileparsers.ParseToStringObj(input[0], ", ")
	parts := parseInput(partsI)

	position := positions.Position{
		X:    0,
		Y:    0,
		Side: positions.Side{Name: "N"},
	}
	for _, l := range parts {
		position.Move(l.Direction, l.Steps, nil)
	}
	return position.CalculateFrom0() //332
}

func Day01Part2() int {
	input := fileparsers.ReadLines("inputs2023\\day1.txt")
	partsI := fileparsers.ParseToStringObj(input[0], ", ")
	parts := parseInput(partsI)

	position := positions.Position{
		X:    0,
		Y:    0,
		Side: positions.Side{Name: "N"},
	}
	visited := make(map[string]bool)
	for _, l := range parts {
		position.Move(l.Direction, l.Steps, visited)
	}
	return position.CalculateFrom0() //332
}

func parseInput(lines []string) []InstDay1 {
	result := make([]InstDay1, len(lines))

	for index, line := range lines {
		result[index] = parseItem(line)
	}

	return result
}

func parseItem(line string) InstDay1 {
	var s1 string
	var s2 int
	fmt.Sscanf(line, "%1s%d", &s1, &s2)

	return InstDay1{
		Direction: s1,
		Steps:     s2,
	}
}

type InstDay1 struct {
	Direction string
	Steps     int
}
