package aoc2023

import (
	"strconv"
	"strings"

	"github.com/amra.satara/learning-go/fileparsers"
)

func Day09Part1() int {
	input := fileparsers.ReadLines("inputs2023\\day9.txt")

	data := make([]DataDay9, len(input))

	for i, v := range input {
		data[i] = parseData(v)
	}

	sum := 0
	for _, dd := range data {
		dd.Lists = findAllHistory(&dd)
		findNextValue(&dd)
		sum += dd.NextValue
	}

	return sum
}

func Day09Part2() int {
	input := fileparsers.ReadLines("inputs2023\\day9.txt")

	data := make([]DataDay9, len(input))

	for i, v := range input {
		data[i] = parseData(v)
	}

	sum := 0
	for _, dd := range data {
		dd.Lists = findAllHistory(&dd)
		findPreviousValue(&dd)
		sum += dd.PreviousValue
	}

	return sum
}

func findPreviousValue(dd *DataDay9) {
	dd.PreviousValue = dd.Lists[0][0]
	for i, v := range dd.Lists[1:] {
		if i%2 == 0 {
			dd.PreviousValue -= v[0]
		} else {
			dd.PreviousValue += v[0]
		}
	}

}

func findNextValue(dd *DataDay9) {
	for i := len(dd.Lists) - 1; i >= 0; i-- {
		dd.NextValue += dd.Lists[i][len(dd.Lists[i])-1]
	}
}

func findAllHistory(dd *DataDay9) [][]int {
	pToList := &dd.InitialList
	dd.Lists = append(dd.Lists, dd.InitialList)
	for {

		list := make([]int, 0)

		allAreZeros := true
		for i := 0; i < len(*pToList)-1; i++ {

			x := ((*pToList)[i+1] - (*pToList)[i])

			if x != 0 {
				allAreZeros = false
			}
			list = append(list, x)
		}
		dd.Lists = append(dd.Lists, list)
		pToList = &list
		if allAreZeros {
			break
		}

	}
	return dd.Lists
}

func parseData(input string) DataDay9 {
	result := DataDay9{}

	parts := strings.Split(input, " ")
	result.InitialList = make([]int, len(parts))
	for i, v := range parts {
		result.InitialList[i], _ = strconv.Atoi(v)
	}

	return result
}

type DataDay9 struct {
	InitialList   []int
	Lists         [][]int
	NextValue     int
	PreviousValue int
}
