package main

import (
	"container/heap"
	"fmt"
	"time"
)

func GenerateNewNode(
	common 			*Common, 
	open_set 		*PriorityQueue, 
	open_hash 		map[string]*Item,
	closed_set 		[]Node, 
	closed 			map[string]Node,  
	current_node 	Node, 
	new_board 		[]int, 
	zindex 			int, 
	direction 		int) {

	// informed		f = g + h	;		à l'air de fonctionner
	// uninformed 	f = h		; 		à l'air de fonctionner
	// greedy 		f = g		; 		est-il fonctionnel ?

	priority := current_node.parent_count// + LinearConflict(new_board, common.goal, common.size)

	strb := fmt.Sprint(new_board)

	new_node := Node{
		board:        new_board,
		move:         direction,
		cost:         priority,
		parent_count: current_node.parent_count + 1,
		distance:     priority,
		parent:       &current_node,
		zindex:       zindex}

	if pq, exists := open_hash[strb]; exists {
		if priority < pq.priority {
			new_item := open_set.update(pq, pq.node, priority)
			open_hash[fmt.Sprint(pq.node.board)] = new_item
		} 
		return 
	} else if closed_node, exists := closed[strb]; exists {
			if priority < closed_node.cost {
				item := &Item{node: new_node, priority: priority}
				heap.Push(open_set, item)
				open_hash[strb] = item
				delete(closed, strb)
			}
			return
		}
	
	item := &Item{node: new_node, priority: priority}
	heap.Push(open_set, item)
	open_hash[strb] = item
}



func GetNextMoves(
	common 			*Common, 
	open_set 		*PriorityQueue, 
	open_hash 		map[string]*Item, 
	closed_set 		[]Node, 
	closed 			map[string]Node, 
	current_node 	Node,
	direction 		int) {
	new_board := make([]int, common.size * common.size)
	copy(new_board, current_node.board)

	switch direction {
		case UP:
			if current_node.zindex - common.size >= 0 {
				new_board[current_node.zindex] = new_board[current_node.zindex - common.size]
				new_board[current_node.zindex - common.size] = 0
				GenerateNewNode(common, open_set, open_hash, closed_set, closed, current_node, new_board, current_node.zindex - common.size, direction)
			}
		case DOWN:
			if current_node.zindex + common.size < common.size * common.size  {
				new_board[current_node.zindex] = new_board[current_node.zindex + common.size]
				new_board[current_node.zindex + common.size] = 0
				GenerateNewNode(common, open_set, open_hash, closed_set, closed, current_node, new_board, current_node.zindex + common.size, direction)
			}
		case LEFT:
			if current_node.zindex % common.size >= 1 {
				new_board[current_node.zindex] = new_board[current_node.zindex - 1]
				new_board[current_node.zindex - 1] = 0
				GenerateNewNode(common, open_set, open_hash, closed_set, closed, current_node, new_board, current_node.zindex - 1, direction)
			}
		case RIGHT:
			if current_node.zindex % common.size < common.size - 1 {
				new_board[current_node.zindex] = new_board[current_node.zindex + 1]
				new_board[current_node.zindex + 1] = 0
				GenerateNewNode(common, open_set, open_hash, closed_set, closed, current_node, new_board, current_node.zindex + 1, direction)
			}
		default:
			fmt.Print("--")
	}
}


func new_astar(common *Common, board []int) {

	var node Node

	open_set 	:= make(PriorityQueue, 0)
	open_hash 	:= make(map[string]*Item)
	closed 		:= make(map[string]Node)
	closed_set	:= make([]Node, 0)


	node = Node{
		board:        board,
		move:         NONE,
		cost:         0,
		parent_count: 0,
		distance:     0,
		parent:       nil,
		zindex:       FindIndex(board, 0)}

	item := &Item{node: node, priority: 0}
	heap.Push(&open_set, item)
	open_hash[fmt.Sprint(node.board)] = item

	ticker := time.NewTicker(1 * time.Second)
	time_start 				:= time.Now()

	var old_closed_count int
	var old_open_count int
	go func() {
		for {
			select {
			case <- ticker.C:
				fmt.Print("\033[H\033[2J")
				fmt.Println("Explored nodes : ", len(closed))
				fmt.Println("Opens nodes : ", open_set.Len(), len(open_hash), open_set.Len() ==  len(open_hash))
				fmt.Println(len(closed) - old_closed_count, " nodes / s" )
				fmt.Println(open_set.Len() - old_open_count, " nodes / s" )
				//old_closed_count = len(solver.closed_set2)
				old_closed_count = len(closed)
				old_open_count = open_set.Len()
			}
		}
	}()
	
	for open_set.Len() != 0 {

		node = heap.Pop(&open_set).(*Item).node
		strb := fmt.Sprint(node.board)

		delete(open_hash, strb)

		//closed_set = append(closed_set, node)
		closed[strb] = node

		if Compare(node.board, common.goal) {
			fmt.Println("Found goal :D")
			fmt.Println(node.parent_count)
			fmt.Println(node.parent)
			
				time_elapsed := time.Since(time_start)
				fmt.Printf("> %-18s : %6.3fs\n", "Time taken", time_elapsed.Seconds())
			break
		}
		GetNextMoves(common, &open_set, open_hash, closed_set, closed, node, UP)
		GetNextMoves(common, &open_set, open_hash, closed_set, closed, node, DOWN)
		GetNextMoves(common, &open_set, open_hash, closed_set, closed, node, LEFT)
		GetNextMoves(common, &open_set, open_hash, closed_set, closed, node, RIGHT)

		//fmt.Println("Closed : ", len(closed))
		//fmt.Println("Open : ",open_set.Len())
		//fmt.Println()
	}

	_ = open_set
	_ = open_hash
	_ = closed
}