package aoc2023

import (
	"strconv"
	"strings"

	"github.com/amra.satara/learning-go/fileparsers"
)

func Day04Part1() int {
	input := fileparsers.ReadLines("inputs2023\\day4.txt")

	sum := 0
	for _, line := range input {
		game := parse(line)
		sum += game.Score
	}
	return sum
}

func Day04Part2() int {
	input := fileparsers.ReadLines("inputs2023\\day4.txt")

	games := make([]GameDay4, 0)
	sum := 0
	for _, line := range input {
		game := parse(line)
		games = append(games, game)
	}

	for i := 0; i < len(games); i++ {
		c := games[i]
		for k := 0; k < c.Copies+1; k++ {
			for j := 0; j < c.MatchesCount; j++ {
				games[i+j+1].Copies++
			}
		}
	}

	for i := 0; i < len(games); i++ {
		sum++
		sum += games[i].Copies
	}

	return sum
}

func parse(line string) (game GameDay4) {
	game = GameDay4{}
	line, _ = strings.CutPrefix(line, "Card")

	parts := strings.Split(line, ":")

	game.Id, _ = strconv.Atoi(strings.ReplaceAll(parts[0], " ", ""))

	parts = strings.Split(parts[1], "|")

	game.WinningNumbers = makeArray(parts[0])
	game.MyNumbers = makeArray(parts[1])

	game = FIndMatches(game)

	return game
}
func makeArray(line string) []int {
	result := make([]int, 0)
	for i := 0; i < len(line)-2; i += 3 {
		s := strings.ReplaceAll(line[i:i+3], " ", "")
		x, _ := strconv.Atoi(s)
		result = append(result, x)
	}
	return result
}

type GameDay4 struct {
	Id             int
	WinningNumbers []int
	MyNumbers      []int
	// part1
	Score int
	//part2
	MatchesCount int
	Copies       int
}

func FIndMatches(game GameDay4) GameDay4 {

	game.Score = 0
	for _, m := range game.MyNumbers {
		for _, w := range game.WinningNumbers {
			if m == w {
				game.MatchesCount++
				if game.Score == 0 {
					game.Score = 1
				} else {
					game.Score *= 2
				}
			}
		}
	}

	return game
}
