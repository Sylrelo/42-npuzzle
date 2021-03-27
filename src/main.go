package main

import (
	"fmt"
	"container/list"
	_"container/heap"
)

const (
	UP 		= iota
	DOWN 	= iota
	LEFT 	= iota
	RIGHT 	= iota
)
const NSIZE = 9
const NCOL 	= 3

func FindIndex(haystack [9]int, needle int) int {
	for i, n := range haystack {
		if (needle == n) {
			return i
		}
	}
	return -1
}

func Same(a [9]int, b [9]int) bool {
	for i, n := range a {
		if (b[i] != n) {
			return false
		}
	}
	return true
}

type Node struct {
	parent		*Node
	board		[9]int
	move		int
	cost		int
	distance	int
}

type Solver struct {
	open		*list.List
	closed		*list.List
	ncol		int
}

func PrintBoard(board [9]int) {
	for i := 0; i < 3; i++ {
		fmt.Print(board[0 + i], " ")
	}
	fmt.Println("")
	for i := 0; i < 3; i++ {
		fmt.Print(board[3 + i], " ")
	}
	fmt.Println("")
	for i := 0; i < 3; i++ {
		fmt.Print(board[6 + i], " ")
	}
	fmt.Println("")
}

func Move(board [9]int, direction int) (bool, [9]int) {
	zindex := FindIndex(board, 0)

	switch direction {
		case UP:
			if (zindex - NCOL >= 0) {
				board[zindex] = board[zindex - NCOL]
				board[zindex - NCOL] = 0
				return true, board
			}
		case DOWN:
			if (zindex + NCOL < NSIZE) {
				board[zindex] = board[zindex + NCOL]
				board[zindex + NCOL] = 0
				return true, board
			}
		case LEFT:
			if (zindex % NCOL >= 1) {
				board[zindex] = board[zindex - 1]
				board[zindex - 1] = 0
				return true, board
			}
		case RIGHT:
			if (zindex % NCOL <= 1) {
				board[zindex] = board[zindex + 1]
				board[zindex + 1] = 0
				return true, board
			}
		default:
			fmt.Print("--")
	}
	return false, board
}

func main() {
	base := [9]int {1, 8, 2, 0, 4, 3, 7, 6, 5}
	var solver Solver
	//base := [3][3] int {{1,2,3}, {4,5,6}, {7,8,9}}

	solver.open 	= list.New()
	solver.closed 	= list.New()

	var node Node

	node = Node{
		board: base,
		move: 5,
		cost: 0,
		distance: 0}

	_ = node

	solver.open.PushBack(base)

	fmt.Println(solver.open)
	for e := solver.open.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }

	for i := 0; i < 2; i++ {
		var state, new_board = Move(base, RIGHT)

		_ = state
		_ = new_board
	}


	// base = [9]int {1, 8, 2, 0, 4, 3, 7, 6, 5}
	// Move(base, RIGHT)
	// base = [9]int {1, 8, 0, 2, 4, 3, 7, 6, 5}
	// Move(base, RIGHT)
	// base = [9]int {1, 0, 8, 2, 4, 3, 7, 6, 5}
	// Move(base, RIGHT)
	// base = [9]int {0, 1, 8, 2, 4, 3, 7, 6, 5}
	// Move(base, RIGHT)
	// base = [9]int {4, 1, 8, 2, 0, 3, 7, 6, 5}
	// Move(base, RIGHT)
	// base = [9]int {4, 1, 8, 2, 3, 0, 7, 6, 5}
	// Move(base, RIGHT)
	// base = [9]int {4, 1, 8, 2, 3, 7, 0, 6, 5}
	// Move(base, RIGHT)
	// base = [9]int {4, 1, 8, 2, 3, 7, 6, 0, 5}
	// Move(base, RIGHT)
	// base = [9]int {4, 1, 8, 2, 3, 7, 6, 5, 0}
	// Move(base, RIGHT)

	_ = base

}