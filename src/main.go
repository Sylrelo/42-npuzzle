package main

import (
	"container/heap"
	_ "container/list"
	"fmt"
)

func NewNode(open_set *PriorityQueue, closed_set [][]int, current_node Node, new_board []int, zindex int, direction int) {
	goal := []int{1, 2, 3, 8, 0, 4, 7, 6, 5}

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
		parent:       &current_node,
		zindex:       zindex}

	heap.Push(open_set, &Item{node: new_node, priority: priority})
	fmt.Println("\033[1;36m+ Queue push\033[0m")
}

func Move(open_set *PriorityQueue, closed_set [][]int, current_node Node, direction int) {
	new_board := make([]int, len(current_node.board))
	copy(new_board, current_node.board)

	switch direction {
	case UP:
		if current_node.zindex - NCOL >= 0 {
			new_board[current_node.zindex] = new_board[current_node.zindex - NCOL]
			new_board[current_node.zindex - NCOL] = 0
			NewNode(open_set, closed_set, current_node, new_board, current_node.zindex - NCOL, direction)
		}
	case DOWN:
		if current_node.zindex + NCOL < NSIZE {
			new_board[current_node.zindex] = new_board[current_node.zindex + NCOL]
			new_board[current_node.zindex + NCOL] = 0
			NewNode(open_set, closed_set, current_node, new_board, current_node.zindex + NCOL, direction)
		}
	case LEFT:
		if current_node.zindex % NCOL >= 1 {
			new_board[current_node.zindex] = new_board[current_node.zindex - 1]
			new_board[current_node.zindex - 1] = 0
			NewNode(open_set, closed_set, current_node, new_board, current_node.zindex - 1, direction)
		}
	case RIGHT:
		if current_node.zindex % NCOL <= 1 {
			new_board[current_node.zindex] = new_board[current_node.zindex + 1]
			new_board[current_node.zindex + 1] = 0
			NewNode(open_set, closed_set, current_node, new_board, current_node.zindex + 1, direction)
		}
	default:
		fmt.Print("--")
	}
}

func main() {
	base 			:= []int{1, 8, 2, 0, 4, 3, 7, 6, 5}
	open_set 		:= make(PriorityQueue, 0)

	var node 		Node
	var size 		Size
	var closed_set 	[][]int

	size.nsize 		= 9
	size.ncol 		= 3


	snail := GenerateSnail(3, 3)

	PrintBoard(snail, size)

	return 
	node = Node{
		board:        base,
		move:         NONE,
		cost:         0,
		parent_count: 0,
		distance:     0,
		parent:       nil,
		zindex:       FindIndex(base, 0)}

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
