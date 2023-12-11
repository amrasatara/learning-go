package aoc2023

import (
	"github.com/amra.satara/learning-go/fileparsers"
)

func Day11Part1() int {
	input := fileparsers.ReadLines("inputs2023\\day11.txt")

	parseNodes(input)

	sum := 0
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			node1 := nodes[i]
			node2 := nodes[j]

			sum += Distance(&node1, &node2, 2)
		}
	}

	return sum
}

var nodes []Node
var emptyX []int
var emptyY []int

func Day11Part2() int {
	input := fileparsers.ReadLines("inputs2023\\day11.txt")

	parseNodes(input)

	sum := 0
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			node1 := nodes[i]
			node2 := nodes[j]

			sum += Distance(&node1, &node2, 1000000)
		}

	}

	return sum
}

func parseNodes(input []string) {
	nodes = make([]Node, 0)
	for i, v := range input {

		for j, l := range v {
			if l == '#' {
				nodes = append(nodes, Node{j, i})
			}
		}
	}
	for i := 0; i < len(input); i++ {
		empty := true
		for _, n := range nodes {
			if n.Y == i {
				empty = false
				break
			}
		}
		if empty {
			emptyY = append(emptyY, i)
		}
	}
	for j := 0; j < len(input[0]); j++ {
		empty := true
		for _, n := range nodes {
			if n.X == j {
				empty = false
				break
			}
		}
		if empty {
			emptyX = append(emptyX, j)
		}
	}
}
func Distance(node1 *Node, node2 *Node, times int) int {
	xdist := 0

	for i := min(node1.X, node2.X); i < max(node1.X, node2.X); i++ {

		found := false
		for _, v := range emptyX {
			if v == i {
				found = true
				break
			}
		}
		if found {
			xdist += times
		} else {
			xdist++
		}
	}
	for i := min(node1.Y, node2.Y); i < max(node1.Y, node2.Y); i++ {

		found := false
		for _, v := range emptyY {
			if v == i {
				found = true
				break
			}
		}
		if found {
			xdist += times
		} else {
			xdist++
		}
	}
	return xdist
}

type Node struct {
	X int
	Y int
}
