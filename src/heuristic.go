package main

import (
	"math"
)

/*
** Number of missplaced tiles
 */
func HammingDistance(board []int, goal []int) int {
	var total int = 0

	for i := range board {
		if board[i] != 0 && board[i] != goal[i] {
			total++
		}
	}
	return total
}

/*
** Distance between the current tile and the goal tile
** [x] [x] [x] [x] [g]
** [x] [ ] [ ] [ ] [ ]
** [x] [ ] [ ] [ ] [ ]
** [x] [ ] [ ] [ ] [ ]
** [s] [ ] [ ] [ ] [ ]
*/

func ManhattanDistance(board []int, goal []int) int {
	var total float64

	total = 0
	for i := range board {
		if board[i] != 0 {
			pos := FindIndex(goal, board[i])
			total += math.Abs(float64(i % NCOL - pos % NCOL)) + math.Abs(float64(i / NCOL - pos / NCOL))
		}
	}
	return int(total)
}

/*
** Distance between the current tile and the goal tile
** [ ] [ ] [ ] [ ] [g]
** [ ] [ ] [ ] [x] [ ]
** [ ] [ ] [x] [ ] [ ]
** [ ] [x] [ ] [ ] [ ]
** [s] [ ] [ ] [ ] [ ]
*/

func EuclideanDistance(board []int, goal []int) int {
	var total float64

	for i := range board {
		if board[i] != 0 {
			pos := FindIndex(goal, board[i])
			cx := (i % NCOL)
			gx := (pos % NCOL)
			cy := (i / NCOL)
			gy := (pos / NCOL)
			total += math.Sqrt(float64(((cx - gx) * (cx - gx)) + ((cy - gy) * (cy - gy))))
		}
	}
	return int(total)
}

func LinearConflict() {
	//
}
