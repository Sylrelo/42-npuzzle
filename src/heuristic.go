package main

import (
	_ "fmt"
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

func ManhattanDistance(board []int, goal []int, size int) int {
	var total float64

	total = 0
	for i := range board {
		if board[i] != 0 {
			pos := FindIndex(goal, board[i])
			total += math.Abs(float64(i % size - pos % size)) + math.Abs(float64(i / size - pos / size))
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

func EuclideanDistance(board []int, goal []int, size int) int {
	var total float64

	for i := range board {
		if board[i] != 0 {
			pos := FindIndex(goal, board[i])
			cx := (i % size)
			gx := (pos % size)
			cy := (i / size)
			gy := (pos / size)
			total += math.Sqrt(float64(((cx - gx) * (cx - gx)) + ((cy - gy) * (cy - gy))))
		}
	}
	return int(total)
}

/*
**
*/

func LinearConflict(board []int, goal []int, size int) int {
	var linear 		LinearConflictHelper

	conflicts 		:= 0
	glen 			:= len(goal)
	linear.board 	= make([]vec2i, 0)
	linear.goal 	= make([]vec2i, 0)

	for i := 0; i < glen; i++ {
			gi 	:= FindIndex(goal, i)
			gix := gi % size
			giy := gi / size
			bi 	:= FindIndex(board, i)
			bix := bi % size
			biy := bi / size
			linear.board = append(linear.board, vec2i{bix, biy})
			linear.goal = append(linear.goal, vec2i{gix, giy})
	}

	for i := 1; i < glen; i++ {
		gi := linear.goal[i]
		bi := linear.board[i]
		for j := i + 1; j < glen; j++ {
			gj := linear.goal[j]
			bj := linear.board[j]
			if bi.x == bj.x && gi.x == gj.x && ((bi.y > bj.y && gi.y < gj.y) || (bi.y < bj.y && gi.y > gj.y)) {
				conflicts++
			}
			if bi.y == bj.y && gi.y == gj.y && ((bi.x > bj.x && gi.x < gj.x) || (bi.x < bj.x && gi.x > gj.x)) {
				conflicts++
			}
		}
	}
	return conflicts
}
