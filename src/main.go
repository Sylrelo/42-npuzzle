package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func NewNode(open_set *PriorityQueue, closed_set [][]int, current_node Node, new_board []int, zindex int, direction int) {
	goal := []int{1, 2, 3, 8, 0, 4, 7, 6, 5}

	if Same(closed_set, new_board) {
		//fmt.Println("\033[33m~ Board already explored. Skipping.\033[0m")
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
	//fmt.Println("\033[1;36m+ Queue push\033[0m")
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
	open_set 		:= make(PriorityQueue, 0)
	goal			:= []int{1, 2, 3, 8, 0, 4, 7, 6, 5}

	var base 		[]int
	var node 		Node
	var size 		Size
	var closed_set 	[][]int

	size.nsize 		= 9
	size.ncol 		= 3

	reader 		:= bufio.NewReader(os.Stdin)
	content, _ 	:= reader.ReadString(0)
	content_splitted := strings.Split(content, "\n")


	if content_splitted[0][0] == '#' { 
		size.ncol, _ = strconv.Atoi(content_splitted[1])

		for i := 2; i <= 4; i++ {

			split_values := strings.Split(content_splitted[i], " ")
			for _, n := range split_values {
				atoi, _ := strconv.Atoi(n)
				base = append(base, atoi)
			}
		}
	}

	time_start := time.Now()
	
	node = Node{
		board:        base,
		move:         NONE,
		cost:         0,
		parent_count: 0,
		distance:     0,
		parent:       nil,
		zindex:       FindIndex(base, 0)}

	heap.Push(&open_set, &Item{node: node, priority: 0})

	oh := 0
	for oh < 10000 {

		if open_set.Len() == 0 {
			fmt.Println("\033[1;31mEmpty queue, break.\033[0m")
			break
		}
		//fmt.Println("\033[1;34m- Queue pop.\033[0m")
		node := heap.Pop(&open_set).(*Item).node

		if Compare(node.board, goal) {
			fmt.Println("\033[1;34mSolution Found !\033[0m")
			fmt.Println(node.parent.parent_count, " parents")
			fmt.Println(len(closed_set), " closed sets")
			fmt.Println(len(open_set), " open sets")
			break 
		}
		//PrintBoard(node.board, size)
		closed_set = append(closed_set, node.board)

		//open_set.Push(&Item{value: node, priority: i})
		// node = open_set 			get highest priority
		// node.board == goal 		stop

		Move(&open_set, closed_set, node, UP)
		Move(&open_set, closed_set, node, DOWN)
		Move(&open_set, closed_set, node, LEFT)
		Move(&open_set, closed_set, node, RIGHT)
		oh++
	}
	time_elapsed := time.Since(time_start)

	fmt.Printf("Time taken %s \n", time_elapsed)
	//fmt.Println(closed_set)
	//fmt.Println(open_set)

	//for open_set.Len() > 0 {
		//cnode := heap.Pop(&open_set).(*Item)
		//fmt.Println(cnode.priority, cnode.node.move, cnode.node.parent_count, cnode.node.board)
	//}

}
