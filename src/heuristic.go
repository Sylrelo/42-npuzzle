package main

import "math"

func HammingDistance(board [9]int, goal [9]int) int {
	var total int

	total = 0
	for i := range board {
		if board[i] != 0 && board[i] != goal[i] {
			total++
		}
	}
	return total
}

func ManhattanDistance(board [9]int, goal [9]int) float64  {
	var total float64

	total = 0
	for i := range board {
		if (board[i] != 0) {
			pos := FindIndex(goal, board[i])
			total += math.Abs(float64(i % NSIZE - pos % NSIZE)) + math.Abs(float64(i / NSIZE - pos / NSIZE))
		}
	}
	return total
}

func EuclideanDistance() {
	//
}

func LinearConflict() {
	//
}