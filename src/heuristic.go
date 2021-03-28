package main

import (
	"fmt"
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

/*
**
*/
func LinearConflict(board []int, goal []int) int {
	var conflicts int
	glen := len(goal)

	conflicts = 0

	for i := 1; i < glen; i++ {
		pb := FindIndex(board, i)
		pg := FindIndex(goal, i)

		gx := pg % NCOL
		gy := pg / NCOL
		cx := pb % NCOL
		cy := pb / NCOL

		fmt.Println("[", i, "]", cx, cy, " - ", gx, gy)
		if cx == gx {
			if (cy > gy) {
				fmt.Println("CONFLIT 1")
				conflicts++
			}
			if (cy < gy) {
				fmt.Println("CONFLIT 2")	
				conflicts++
			}
		}
		if cy == gy  {
			if (cx > gx) {
				//fmt.Println("CONFLIT 3")
				conflicts++
			}
			if (cx < gx) {
				//fmt.Println("CONFLIT 4")
				conflicts++
			}
		}
	}

	return conflicts
}
