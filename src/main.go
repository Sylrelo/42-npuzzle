package main

import (
	"fmt"
	_"container/list"
)

const (
	UP 		= iota
	DOWN 	= iota
	LEFT 	= iota
	RIGHT 	= iota
)

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

// type Node struct {
// 	relative   *Node
// 	nbrMove    int
// 	cout       int
// 	status     int
// 	StateBoard [][]int
// }

// type Solver struct {
// 	nbrRow    int
// 	openList  *list.List
// 	closeList *list.List
// 	Heuristic func([][]int, int) int
// }

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

func move(board [9]int, direction int) [9]int {
	zindex := FindIndex(board, 0)
	
	PrintBoard(board)
	_ = zindex

	switch direction {
		case UP:
			if (zindex - 3 >= 0) {
				board[zindex] = board[zindex - 3]
				board[zindex - 3] = 0
				fmt.Print("UP")
			}
		case DOWN:
			if (zindex + 3 < 9) {
				board[zindex] = board[zindex + 3]
				board[zindex + 3] = 0
				fmt.Print("DOWN")
			}
		case LEFT:
			if (zindex % 3 >= 1) {
				board[zindex] = board[zindex - 1]
				board[zindex - 1] = 0
				fmt.Print("LEFT")
			}
		case RIGHT:
			if (zindex % 3 <= 1) {
				board[zindex] = board[zindex + 1]
				board[zindex + 1] = 0
				fmt.Print("RIGHT")
			}
		default:
			fmt.Print("--")
	}
	fmt.Println("")
	PrintBoard(board)

	fmt.Println("")
	fmt.Println("--")
	return board
}

func main() {
	//var base [3][3] int

	base := [9]int {1, 8, 2, 0, 4, 3, 7, 6, 5}
	//base := [3][3] int {{1,2,3}, {4,5,6}, {7,8,9}}



	base = [9]int {1, 8, 2, 0, 4, 3, 7, 6, 5}
	move(base, RIGHT)
	base = [9]int {1, 8, 0, 2, 4, 3, 7, 6, 5}
	move(base, RIGHT)
	base = [9]int {1, 0, 8, 2, 4, 3, 7, 6, 5}
	move(base, RIGHT)
	base = [9]int {0, 1, 8, 2, 4, 3, 7, 6, 5}
	move(base, RIGHT)
	base = [9]int {4, 1, 8, 2, 0, 3, 7, 6, 5}
	move(base, RIGHT)
	base = [9]int {4, 1, 8, 2, 3, 0, 7, 6, 5}
	move(base, RIGHT)
	base = [9]int {4, 1, 8, 2, 3, 7, 0, 6, 5}
	move(base, RIGHT)
	base = [9]int {4, 1, 8, 2, 3, 7, 6, 0, 5}
	move(base, RIGHT)
	base = [9]int {4, 1, 8, 2, 3, 7, 6, 5, 0}
	move(base, RIGHT)

	_ = base

	for i := 0; i < 1; i++ {
		//tmp := base

		//move(tmp, UP)
		// move(tmp, DOWN)
		// move(tmp, LEFT)
		// move(tmp, RIGHT)
		// fmt.Println(tmp)
		//fmt.Println(caca)
	}
}