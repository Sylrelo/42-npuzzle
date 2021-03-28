package main

import (
	"container/heap"
	_ "container/list"
	"fmt"
)

func NewNode(open_set *PriorityQueue, closed_set [][]int, current_node Node, new_board []int, direction int) {
	goal := []int{1, 2, 3, 8, 0, 4, 7, 6, 5}

	//PrintBoard(current_node.board, Size{9, 3})
	//PrintBoard(new_board, Size{9, 3})

	if Same(closed_set, new_board) {
		fmt.Println("\033[33m~ Board already explored. Skipping.\033[0m")
		return
	}

	priority := current_node.parent_count + ManhattanDistance(new_board, goal)

	new_node := Node{
		board:        new_board,
		move:         direction,
		cost:         priority,
		parent_count: current_node.parent_count + 1,
		distance:     priority,
		parent:       &current_node}

	heap.Push(open_set, &Item{node: new_node, priority: priority})
	fmt.Println("\033[1;36m+ Queue push\033[0m")
}

func Move(open_set *PriorityQueue, closed_set [][]int, current_node Node, direction int) {
	zindex := FindIndex(current_node.board, 0)
	new_board := make([]int, len(current_node.board))
	copy(new_board, current_node.board)

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
	open_set := make(PriorityQueue, 0)
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

	heap.Push(&open_set, &Item{node: node, priority: 0})

	//heap.Init(&open_set)

	for i := 0; i < 8; i++ {

		if open_set.Len() == 0 {
			fmt.Println("\033[1;31mEmpty queue, break.\033[0m")
			break
		}
		fmt.Println("\033[1;34m- Queue pop.\033[0m")
		node := heap.Pop(&open_set).(*Item).node

		PrintBoard(node.board, size)
		closed_set = append(closed_set, node.board)

		//open_set.Push(&Item{value: node, priority: i})
		// node = open_set 			get highest priority
		// node.board == goal 		stop

		Move(&open_set, closed_set, node, UP)
		Move(&open_set, closed_set, node, DOWN)
		Move(&open_set, closed_set, node, LEFT)
		Move(&open_set, closed_set, node, RIGHT)
	}

	fmt.Println(closed_set)
	fmt.Println(open_set)

	for open_set.Len() > 0 {
		cnode := heap.Pop(&open_set).(*Item)
		fmt.Println(cnode.priority, cnode.node.move, cnode.node.parent_count, cnode.node.board)
	}

}
