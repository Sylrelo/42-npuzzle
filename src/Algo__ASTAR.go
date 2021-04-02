package main

import (
	"container/heap"
	"fmt"
	"time"
)

func GenerateNewNode(common *Common, open_set *PriorityQueue, open_hash map[string]*Item, closed map[string]Node, current_node Node, new_board []int, zindex int, direction int) {

	// informed		f = g + h	;		à l'air de fonctionner
	// uninformed 	f = h		; 		à l'air de fonctionner
	// greedy 		f = g		; 		est-il fonctionnel ?

	priority := current_node.parent_count + LinearConflict(new_board, common.goal, common.size)

	strb := fmt.Sprint(new_board)

	new_node := Node{
		board:        new_board,
		move:         direction,
		cost:         priority,
		parent_count: current_node.parent_count + 1,
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



func new_astar(common *Common, board []int) {

	var node Node

	complexity_in_size 	:= 0
	complexity_in_time 	:= 0
	open_set 			:= make(PriorityQueue, 0)
	open_hash 			:= make(map[string]*Item)
	closed 				:= make(map[string]Node)

	// switch heuristic {
	// 	case MANHATTAN:
	// 		solver.heuristic = ManhattanDistance
	// 	case HAMMING:
	// 		solver.heuristic = HammingDistance
	// 	case LINEAR_CONFLICT:
	// 		solver.heuristic = LinearConflict
	// 	default:
	// 		fmt.Println("Wrong heuristic")
	// 		return
	// }


	node = Node{
		board:        board,
		move:         NONE,
		cost:         0,
		parent_count: 0,
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
			//PrintBoard(node.board, common.size)


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
		closed[strb] = node

		complexity_in_time++

		if len_open_set := open_set.Len(); len_open_set > complexity_in_size {
			complexity_in_size = len_open_set
		}

		_ = time_start
		if Compare(node.board, common.goal) {
			SolutionFound(common, Result{
				time_start: time_start,
				complexity_in_size: complexity_in_size,
				complexity_in_time: complexity_in_time,
			})
			break
		}

		for dir := UP; dir <= RIGHT; dir++ {
			if moved, board, zindex := GenerateNextMoves(common, node, dir); moved {
				GenerateNewNode(common, &open_set, open_hash, closed, node, board, zindex, dir)
			}
		}
	
	}

}