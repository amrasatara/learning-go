package aoc2023

import (
	"fmt"
	"strings"

	"github.com/amra.satara/learning-go/fileparsers"
	gameday2 "github.com/amra.satara/learning-go/gameDay2"
)

func Day02Part1() int {

	input := fileparsers.ReadLines("inputs2023\\day2.txt")

	games := make([]gameday2.GameDay2, len(input))
	sum := 0

	for i, line := range input {
		game := parseGame(line)
		if game.MaxBlue() <= 14 && game.MaxGreen() <= 13 && game.MaxRed() <= 12 {

			sum += game.Id
		}
		games[i] = game
	}

	return sum
}

func Day02Part2() int {
	input := fileparsers.ReadLines("inputs2023\\day2.txt")

	games := make([]gameday2.GameDay2, len(input))
	sum := 0

	for i, line := range input {
		game := parseGame(line)
		sum += (game.MaxBlue() * game.MaxGreen() * game.MaxRed())

		games[i] = game
	}

	return sum
}

//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green

func parseGame(line string) gameday2.GameDay2 {
	game := gameday2.GameDay2{}
	parts1 := strings.Split(line, ":")

	fmt.Sscanf(parts1[0], "Game %d", &game.Id)

	roundparts := strings.Split(parts1[1], ";")

	for _, p := range roundparts {
		game.Rounds = append(game.Rounds, parseRound(p))

	}

	return game

}

func parseRound(line string) gameday2.GameRoundDay2 {
	// 1 red, 2 green, 6 blue
	parts := strings.Split(line, ",")

	round := gameday2.GameRoundDay2{}
	for _, part := range parts {
		var d1 int
		var s1 string
		fmt.Sscanf(part, "%d %s", &d1, &s1)
		if s1 == "green" {
			round.Green = d1
		}
		if s1 == "red" {
			round.Red = d1
		}
		if s1 == "blue" {
			round.Blue = d1
		}
	}

	return round
}
