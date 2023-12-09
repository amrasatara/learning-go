package aoc2023

import (
	"fmt"
	"strings"

	"github.com/amra.satara/learning-go/fileparsers"
)

var directions []rune
var instructions []*InstrDay8

func Day08Part1() int {
	input := fileparsers.ReadLines("inputs2023\\day8.txt")

	MakeDirections(input[0])
	MakeInstructions(input[2:])

	return FindSteps(GetInstr("AAA"), "ZZZ")
}

func FindSteps(next *InstrDay8, endsWith string) int {
	for i := 0; ; i++ {
		d := NextDirection(i)
		if d == 'R' {
			next = next.Right
		} else {
			next = next.Left
		}
		if len(next.Name) > len(strings.TrimSuffix(next.Name, endsWith)) {
			return i + 1
		}
	}
}

func Day08Part2() int {
	input := fileparsers.ReadLines("inputs2023\\day8.txt")

	MakeDirections(input[0])
	MakeInstructions(input[2:])

	list := make([]int, 0)

	for _, instruction := range instructions {

		if instruction.Name[2] == 'A' {
			list = append(list, FindSteps(instruction, "Z"))
		}

	}

	result := list[0]
	for i := 1; i < len(list); i++ {
		result = lcm(result, list[i])

	}
	return result
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func MakeInstructions(list []string) {

	instructions = make([]*InstrDay8, 0)

	for _, v := range list {
		var s1, s2, s3 string

		fmt.Sscanf(v, "%3s = (%3s, %3s)", &s1, &s2, &s3)
		i1 := GetInstr(s1)
		i2 := GetInstr(s2)
		i3 := GetInstr(s3)

		i1.Left = i2
		i1.Right = i3
	}
}

func GetInstr(s1 string) *InstrDay8 {
	for _, r := range instructions {
		if r.Name == s1 {
			return r
		}
	}
	r2 := InstrDay8{
		Name: s1,
	}
	instructions = append(instructions, &r2)

	return &r2
}

type InstrDay8 struct {
	Name  string
	Left  *InstrDay8
	Right *InstrDay8
}

func MakeDirections(input string) {
	directions = []rune(input)
}

func NextDirection(step int) rune {
	max := len(directions)

	return directions[step%max]
}
