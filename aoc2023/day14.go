package aoc2023

import (
	"strings"

	"github.com/amra.satara/learning-go/fileparsers"
)

var data [][]rune
var original [][]rune
var sizeX, sizeY int

func Day14Part1() int {
	input := fileparsers.ReadLines("inputs2023\\day14.txt")

	parseDay14(input)

	goUp()

	sum := 0

	for i := 0; i < sizeY; i++ {
		for j := 0; j < sizeX; j++ {
			v := data[i][j]
			if v == 'O' {
				sum += (sizeY - i)
			}
		}
	}

	// for _, v := range data {
	// 	for _, v2 := range v {
	// 		fmt.Print(string(v2))
	// 	}
	// 	fmt.Println("")
	// }
	return sum
}

func goUp() {
	for i := 0; i < sizeY; i++ {
		for j := 0; j < sizeX; j++ {
			v := data[i][j]

			if v != 'O' {
				continue
			}

			k := i - 1
			for ; k >= 0; k-- {
				if data[k][j] != '.' {
					break
				}
			}
			k++

			if k < i {
				data[k][j], data[i][j] = 'O', '.'
			}
		}
	}
}

func goDown() {
	for i := sizeY - 1; i >= 0; i-- {
		for j := 0; j < sizeX; j++ {
			v := data[i][j]

			if v != 'O' {
				continue
			}

			k := i + 1
			for ; k < sizeY; k++ {
				if data[k][j] != '.' {
					break
				}
			}
			k--

			if k > i {
				data[k][j], data[i][j] = 'O', '.'
			}
		}
	}
}

func goWest() {
	for i := 0; i < sizeY; i++ {
		for j := sizeX - 1; j >= 0; j-- {
			v := data[i][j]

			if v != 'O' {
				continue
			}

			k := j + 1
			for ; k < sizeX; k++ {

				x := data[i][k]
				if x != '.' {
					break
				}
			}
			k--

			if k > j {
				data[i][k], data[i][j] = 'O', '.'
			}
		}
	}
}
func goEast() {
	for i := 0; i < sizeY; i++ {
		for j := 0; j < sizeX; j++ {
			v := data[i][j]

			if v != 'O' {
				continue
			}

			k := j - 1
			for ; k >= 0; k-- {
				if data[i][k] != '.' {
					break
				}
			}
			k++

			if k < j {
				data[i][k], data[i][j] = 'O', '.'
			}
		}
	}
}

func parseDay14(input []string) {
	sizeX = len(input[0])
	sizeY = len(input)
	data = make([][]rune, 0)

	for _, v := range input {
		data = append(data, []rune(v))
	}
}

func Day14Part2() int {
	input := fileparsers.ReadLines("inputs2023\\day14.txt")

	parseDay14(input)

	list := make([]string, 0)

	var i = 0
	circle := 0
	for ; i < 1000000000; i++ {
		goUp()
		goEast()
		goDown()
		goWest()

		y := MakeString()
		for index, v := range list {
			if v == y {
				circle = i - index
				break
			}
		}
		if circle > 0 {
			break
		}
		list = append(list, y)
	}

	times := (1000000000 - i - 1) % circle
	for j := 0; j < times; j++ {
		goUp()
		goEast()
		goDown()
		goWest()
	}
	sum := calculateNorth()

	return sum
}

func calculateNorth() int {
	sum := 0

	for i := 0; i < sizeY; i++ {
		for j := 0; j < sizeX; j++ {
			v := data[i][j]
			if v == 'O' {
				sum += (sizeY - i)
			}
		}
	}
	return sum
}

func MakeString() string {
	sb := strings.Builder{}

	for _, v := range data {
		for _, v2 := range v {
			sb.WriteRune(v2)
		}
	}
	return sb.String()
}
