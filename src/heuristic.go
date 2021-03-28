package main

import "math"

func HammingDistance(board []int, goal []int) int {
	var total int = 0

	for i := range board {
		if board[i] != 0 && board[i] != goal[i] {
			total++
		}
	}
	return total
}

func ManhattanDistance(board []int, goal []int) int {
	var total float64

	total = 0
	for i := range board {
		if board[i] != 0 {
			pos := FindIndex(goal, board[i])
			total += math.Abs(float64(i % NSIZE - pos % NSIZE)) + math.Abs(float64(i / NSIZE - pos / NSIZE))
		}
	}
	return int(total)
}

func EuclideanDistance() {
	//
}

func LinearConflict() {
	//
}
