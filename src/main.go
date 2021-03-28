package main

import (
	_ "container/heap"
	_ "container/list"
	"fmt"
)

func NewNode(open_set *[]Node, closed_set [][]int, current_node Node, new_board []int, direction int) {
	goal := []int{1, 2, 3, 8, 0, 4, 7, 6, 5}

	priority := current_node.parent_count + ManhattanDistance(current_node.board, goal)
	new_node := Node{
		board:        new_board,
		move:         direction,
		cost:         priority,
		parent_count: current_node.parent_count + 1,
		distance:     priority,
		parent:       &current_node}

	*open_set = append(*open_set, new_node)
}

func Move(open_set *[]Node, closed_set [][]int, current_node Node, direction int) {
	zindex := FindIndex(current_node.board, 0)
	new_board := current_node.board

	switch direction {
	case UP:
		if zindex-NCOL >= 0 {
			new_board[zindex] = new_board[zindex-NCOL]
			new_board[zindex-NCOL] = 0
			NewNode(open_set, closed_set, current_node, new_board, direction)
		}
	case DOWN:
		if zindex+NCOL < NSIZE {
			new_board[zindex] = new_board[zindex+NCOL]
			new_board[zindex+NCOL] = 0
			NewNode(open_set, closed_set, current_node, new_board, direction)
		}
	case LEFT:
		if zindex%NCOL >= 1 {
			new_board[zindex] = new_board[zindex-1]
			new_board[zindex-1] = 0
			NewNode(open_set, closed_set, current_node, new_board, direction)
		}
	case RIGHT:
		if zindex%NCOL <= 1 {
			new_board[zindex] = new_board[zindex+1]
			new_board[zindex+1] = 0
			NewNode(open_set, closed_set, current_node, new_board, direction)
		}
	default:
		fmt.Print("--")
	}
	//return false, board
}

func main() {
	base := []int{1, 8, 2, 0, 4, 3, 7, 6, 5}

	var node Node
	var size Size
	var open_set []Node
	var closed_set [][]int

	size.nsize = 9
	size.ncol = 3

	node = Node{
		board:        base,
		move:         NONE,
		cost:         0,
		parent_count: 0,
		distance:     0,
		parent:       nil}

	_ = open_set

	for i := 0; i < 2; i++ {
		closed_set = append(closed_set, node.board)

		Move(&open_set, closed_set, node, UP)
		Move(&open_set, closed_set, node, DOWN)
		Move(&open_set, closed_set, node, LEFT)
		Move(&open_set, closed_set, node, RIGHT)

		//_ = state
		//_ = new_board
	}

	fmt.Println(closed_set)
	fmt.Println(open_set)

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
