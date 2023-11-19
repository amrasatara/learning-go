package aoc2016

import (
	"fmt"

	"github.com/amra.satara/learning-go/fileparsers"
)

func Day03Part1() int {
	input := fileparsers.ReadLines("inputs2016\\day3.txt")
	counter := 0
	for _, line := range input {
		t := ParseLine(line)
		if t.IsValid() {
			counter++
		}
	}

	return counter
}

func Day03Part2() int {
	input := fileparsers.ReadLines("inputs2016\\day3.txt")

	list := make([]Triangle, 0)
	for i := 2; i < len(input); i += 3 {
		var t1, t2, t3 Triangle
		fmt.Sscanf(input[i-2], "%5d%5d%5d", &t1.A, &t2.A, &t3.A)
		fmt.Sscanf(input[i-1], "%5d%5d%5d", &t1.B, &t2.B, &t3.B)
		fmt.Sscanf(input[i], "%5d%5d%5d", &t1.C, &t2.C, &t3.C)
		list = append(list, t1)
		list = append(list, t2)
		list = append(list, t3)
	}

	return CountValid(list)
}

func ParseLine(line string) Triangle {
	var s1, s2, s3 int
	fmt.Sscanf(line, "%5d%5d%5d", &s1, &s2, &s3)
	return Triangle{
		A: s1,
		B: s2,
		C: s3,
	}
}

type Triangle struct {
	A int
	B int
	C int
}

func (t *Triangle) IsValid() bool {
	return t.A+t.B > t.C &&
		t.B+t.C > t.A &&
		t.C+t.A > t.B
}

func CountValid(list []Triangle) int {

	counter := 0
	for _, t := range list {
		if t.IsValid() {
			counter++
		}

	}
	return counter
}
