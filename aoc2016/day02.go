package aoc2016

import (
	"fmt"
	"strings"

	"github.com/amra.satara/learning-go/fileparsers"
)

func Day02Part1() string {
	input := fileparsers.ReadLines("inputs2016\\day2.txt")

	builder := strings.Builder{}
	for _, line := range input {
		var n = PhoneNumber{Number: 5}
		for _, r := range strings.Split(line, "") {
			n.MoveForNormalPhone(r)
		}

		builder.WriteString(fmt.Sprintf("%d", n.Number))
	}
	return builder.String()
}

func Day02Part2() string {
	input := fileparsers.ReadLines("inputs2016\\day2.txt")

	var n = PhoneNumber{Number: 5, Letter: "5"}
	builder := strings.Builder{}
	for _, line := range input {
		for _, r := range strings.Split(line, "") {
			n.MoveForSpecificPhone(r)
		}

		builder.WriteString(n.Letter)
	}
	return builder.String()
}

type PhoneNumber struct {
	Number int
	Letter string
}

func (number *PhoneNumber) MoveForNormalPhone(instr string) {
	switch instr {
	case "U":
		if number.Number > 3 {
			number.Number -= 3
		}
	case "D":
		if number.Number < 7 {
			number.Number += 3
		}
	case "L":
		if number.Number != 1 && number.Number != 4 && number.Number != 7 {
			number.Number--
		}
	case "R":
		if number.Number != 3 && number.Number != 6 && number.Number != 9 {
			number.Number++
		}
	}
}
func (number *PhoneNumber) MoveForSpecificPhone(instr string) {

	lefts := make(map[string]string)
	rights := make(map[string]string)
	ups := make(map[string]string)
	downs := make(map[string]string)

	populate(lefts, "432")
	populate(lefts, "98765")
	populate(lefts, "CBA")
	populate(rights, "234")
	populate(rights, "56789")
	populate(rights, "ABC")
	populate(ups, "A62")
	populate(ups, "DB731")
	populate(ups, "C84")
	populate(downs, "26A")
	populate(downs, "137BD")
	populate(downs, "48C")

	switch instr {
	case "U":
		value, ok := ups[number.Letter]
		if ok {
			number.Letter = value
		}
	case "D":
		value, ok := downs[number.Letter]
		if ok {
			number.Letter = value
		}
	case "L":
		value, ok := lefts[number.Letter]
		if ok {
			number.Letter = value
		}
	case "R":
		value, ok := rights[number.Letter]
		if ok {
			number.Letter = value
		}
	}
}

func populate(m map[string]string, input string) {
	for i := 1; i < len(input); i++ {

		m[string(input[i-1])] = string(input[i])
	}
}
