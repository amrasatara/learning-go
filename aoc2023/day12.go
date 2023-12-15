package aoc2023

import (
	"strconv"
	"strings"

	"github.com/amra.satara/learning-go/fileparsers"
	"github.com/amra.satara/learning-go/helpers"
)

func Day12Part1() int {
	input := fileparsers.ReadLines("inputs2023\\day12.txt")

	sum := 0
	data := make([]SpringsDay12, len(input))
	for i, v := range input {
		data[i] = parseToSpring(v)
		findOptions(&data[i])
		sum += data[i].Options
		//fmt.Println(data[i].Options, v)
	}

	//fmt.Println(data)

	return sum
}

func findOptions(data *SpringsDay12) {
	findNext(data, data.Raw, data.Counts)
}

func parseToSpring(v string) SpringsDay12 {
	parts := strings.Split(v, " ")

	c := strings.Split(parts[1], ",")
	counts := make([]int, len(c))
	for i, v2 := range c {
		counts[i], _ = strconv.Atoi(v2)
	}

	return SpringsDay12{
		Raw:     parts[0],
		Counts:  counts,
		Options: 0,
	}
}

func Day12Part2() int {
	input := fileparsers.ReadLines("inputs2023\\day12.txt")

	ch := make(chan int)
	for _, v := range input {
		go func(line string) {
			d := parseToSpring2(line)
			findOptions(&d)
			ch <- d.Options
		}(v)

	}

	return helpers.SumChanelValues(ch, len(input))
}

func parseToSpring2(v string) SpringsDay12 {
	parts := strings.Split(v, " ")

	c := strings.Split(parts[1], ",")
	counts := make([]int, 0)

	raw := ""
	for j := 0; j < 5; j++ {
		for _, v2 := range c {

			f, _ := strconv.Atoi(v2)
			counts = append(counts, f)
		}
		raw += parts[0]
		if j != 4 {
			raw += "?"
		}

	}

	return SpringsDay12{
		Raw:     raw,
		Counts:  counts,
		Options: 0,
	}
}

func findNext(spring *SpringsDay12, input string, numbers []int) bool {

	if len(numbers) == 0 {
		spring.Options++
		return true
	}
	sum := 0
	for _, v := range numbers {
		sum += v + 1
	}

	if len(input) < sum-1 {
		return false
	}

	number := numbers[0]
	for i, v := range input {
		if i+number > len(input) {
			return false
		}

		toTake := input[i : i+number]

		if strings.Contains(toTake, ".") {
			if v == '#' {
				break
			}
			continue
		}

		after := ""
		//not last one

		if (i + number + 1) <= len(input) {
			after = input[(i + number) : i+number+1]
			if after == "#" {
				if v == '#' {
					break
				}
				continue
			}
			if strings.Contains(input[i+number:], "#") && len(numbers) == 1 {
				if v == '#' {
					break
				}
				continue
			}
		} else if len(numbers) > 1 {
			break
		}

		//n := taken + input[0:i] + strings.ReplaceAll(toTake, "?", "#") + after
		var rest = after

		if len(after) > 0 {
			rest = input[i+number+1:]
		}
		_ = findNext(spring, rest, numbers[1:])

		if v == '#' {
			break
		}
	}
	return false
}

type SpringsDay12 struct {
	Raw     string
	Counts  []int
	Options int
}
