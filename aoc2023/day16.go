package aoc2023

import (
	"strconv"

	"github.com/amra.satara/learning-go/fileparsers"
)

var dataday16 [][]rune
var sizeX16, sizeY16 int
var paths map[string]int
var uniquepaths map[string]int

func Day16Part1() int {
	input := fileparsers.ReadLines("inputs2023\\day16.txt")

	parseDay16(input)
	paths = make(map[string]int, 0)
	uniquepaths = make(map[string]int, 0)

	GoNext('W', 0, 0)

	return len(paths)
}

func GoNext(from rune, x, y int) {
	if x < 0 || y < 0 || x >= sizeX16 || y >= sizeY16 {
		return
	}

	c := dataday16[y][x]
	//fmt.Println("c je ", string(c), " na ", x, " ", y)
	key := string(from) + strconv.Itoa(x) + "-" + strconv.Itoa(y)
	key2 := strconv.Itoa(x) + "-" + strconv.Itoa(y)

	v, ok1 := paths[key2]

	if ok1 {
		paths[key2] = v + 1
	} else {
		paths[key2] = 1
	}
	_, ok := uniquepaths[key]
	if ok {
		return
	} else {
		uniquepaths[key] = 1
	}

	if c == '.' || (c == '-' && (from == 'E' || from == 'W')) || (c == '|' && (from == 'N' || from == 'S')) {
		switch from {
		case 'E':
			{
				GoNext(from, x-1, y)
			}
		case 'W':
			{
				GoNext(from, x+1, y)
			}
		case 'N':
			{
				GoNext(from, x, y+1)
			}
		case 'S':
			{
				GoNext(from, x, y-1)
			}
		}
	} else if c == '\\' {
		switch from {
		case 'E':
			{
				GoNext('S', x, y-1)
			}
		case 'W':
			{
				GoNext('N', x, y+1)
			}
		case 'S':
			{
				GoNext('E', x-1, y)
			}
		case 'N':
			{
				GoNext('W', x+1, y)
			}
		}

	} else if c == '/' {
		switch from {
		case 'W':
			{
				GoNext('S', x, y-1)
			}
		case 'E':
			{
				GoNext('N', x, y+1)
			}
		case 'N':
			{
				GoNext('E', x-1, y)
			}
		case 'S':
			{
				GoNext('W', x+1, y)
			}
		}

	} else if c == '-' {
		GoNext('E', x-1, y)
		GoNext('W', x+1, y)
	} else if c == '|' {
		GoNext('N', x, y+1)
		GoNext('S', x, y-1)
	}
}

func Day16Part2() int {
	input := fileparsers.ReadLines("inputs2023\\day16.txt")

	parseDay16(input)

	maxc := 0
	for i := 0; i < sizeX16; i++ {

		paths = make(map[string]int, 0)
		uniquepaths = make(map[string]int, 0)

		GoNext('N', i, 0)

		pathCounts := len(paths)
		maxc = max(maxc, pathCounts)
	}

	for i := 0; i < sizeX16; i++ {

		paths = make(map[string]int, 0)
		uniquepaths = make(map[string]int, 0)

		GoNext('S', i, sizeY16-1)

		pathCounts := len(paths)
		maxc = max(maxc, pathCounts)
	}

	for i := 0; i < sizeY16; i++ {

		paths = make(map[string]int, 0)
		uniquepaths = make(map[string]int, 0)

		GoNext('W', 0, i)

		pathCounts := len(paths)
		maxc = max(maxc, pathCounts)
	}

	for i := 0; i < sizeY16; i++ {

		paths = make(map[string]int, 0)
		uniquepaths = make(map[string]int, 0)

		GoNext('E', sizeX16-1, i)

		pathCounts := len(paths)
		maxc = max(maxc, pathCounts)
	}

	return maxc
}

type NodeDay16 struct {
	X     int
	Y     int
	Order int
	From  *NodeDay16
}

func parseDay16(input []string) {
	sizeX16 = len(input[0])
	sizeY16 = len(input)
	dataday16 = make([][]rune, 0)

	for _, v := range input {
		dataday16 = append(dataday16, []rune(v))
	}
}
