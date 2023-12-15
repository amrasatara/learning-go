package aoc2023

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amra.satara/learning-go/fileparsers"
)

func Day15Part1() int {
	input := fileparsers.ReadLines("inputs2023\\day15.txt")

	parts := strings.Split(input[0], ",")

	sum := 0

	for _, v := range parts {

		x := HashOperator(v)
		sum += x

	}

	return sum
}

func Day15Part2() int {
	input := fileparsers.ReadLines("inputs2023\\day15.txt")

	parts := strings.Split(input[0], ",")
	boxes := make(map[int]BoxDay15)

	for _, v := range parts {
		label := ""
		isDash := false
		value := 0
		if strings.Contains(v, "-") {
			label = strings.TrimSuffix(v, "-")
			isDash = true

		} else {
			isDash = false
			p := strings.Split(v, "=")
			label = p[0]
			value, _ = strconv.Atoi(p[1])
		}

		order := HashOperator(label)
		fmt.Println("box: ", order)
		box, ok := boxes[order]
		if !ok {

			boxes[order] = BoxDay15{
				Order:  order,
				Lenses: make(map[string]*LenseDay15),
			}
			box = boxes[order]
		}

		lense, okl := box.Lenses[label]
		if isDash {
			if okl {
				delete(box.Lenses, label)

				for _, ld := range box.Lenses {
					if ld.Order > lense.Order {
						ld.Order--
					}
				}
			}

		} else {
			if okl {
				lense.Focus = value
			} else {
				box.Lenses[label] = &LenseDay15{
					Label: label,
					Focus: value,
					Order: len(box.Lenses),
				}
			}
		}
	}

	sum := 0
	for _, box := range boxes {
		for _, lense := range box.Lenses {
			x := (box.Order + 1) * (lense.Order + 1) * lense.Focus
			sum += x
		}

	}
	return sum
}

type BoxDay15 struct {
	Order  int
	Lenses map[string]*LenseDay15
}

type LenseDay15 struct {
	Label string
	Focus int
	Order int
}

func HashOperator(v string) int {
	x := 0
	for _, r := range v {
		x += int(r)
		x *= 17
		x %= 256
	}
	return x
}
