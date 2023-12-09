package aoc2023

import (
	"strconv"
	"strings"

	"github.com/amra.satara/learning-go/fileparsers"
)

func Day06Part1() int {
	input := fileparsers.ReadLines("inputs2023\\day6.txt")

	races := make([]RaceDay6, 0)

	line1 := strings.Replace(input[0], "Time:    ", "", 1)
	line2 := strings.Replace(input[1], "Distance:", "", 1)
	line1 = strings.ReplaceAll(line1, "     ", " ")
	line2 = strings.ReplaceAll(line2, "     ", " ")
	line1 = strings.ReplaceAll(line1, "   ", " ")
	line2 = strings.ReplaceAll(line2, "   ", " ")
	line1 = strings.ReplaceAll(line1, "  ", " ")
	line2 = strings.ReplaceAll(line2, "  ", " ")
	parts1 := strings.Split(line1, " ")
	parts2 := strings.Split(line2, " ")

	for i := range parts1 {
		t, e := strconv.Atoi(parts1[i])
		d, e2 := strconv.Atoi(parts2[i])
		if e != nil || e2 != nil {
			continue
		}
		races = append(races, RaceDay6{t, d})
	}

	sum := 1
	for _, rd := range races {
		distances := FindHoldTimes(rd.Time)
		count := CountBigger(distances, rd.Distance)
		sum *= count
	}

	return sum
}

func Day06Part2() int {
	line1 := "45977295"
	line2 := "305106211101695"
	t, _ := strconv.Atoi(line1)
	d, _ := strconv.Atoi(line2)
	rd := RaceDay6{t, d}
	distances := FindHoldTimes(rd.Time)
	count := CountBigger(distances, rd.Distance)

	return count
}

func FindHoldTimes(time int) []int {
	result := make([]int, 0)
	for i := 1; i < time; i++ {
		x := i * (time - i)
		result = append(result, x)

	}

	return result
}
func CountBigger(list []int, tr int) int {
	result := 0
	for _, v := range list {
		if v > tr {
			result++
		}
	}

	return result
}

type RaceDay6 struct {
	Time     int
	Distance int
}
