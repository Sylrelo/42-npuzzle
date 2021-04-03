package main

import (
	"container/list"
	"fmt"
	"time"
)

var pattern_663 = [][]int{
							{	0, 1, 1, 0,
								1, 1, 0, 0,
								1, 1, 0, 0,
								0, 0, 0, 0,
							},
							{	0, 0, 0, 1,
								0, 0, 1, 1,
								0, 0, 1, 1,
								0, 0, 0, 1,
							},
							{	0, 0, 0, 0,
								0, 0, 0, 0,
								0, 0, 0, 0,
								1, 1, 1, 0,
							},
						}


// func DisjointMoves(common *Common, current_node Node, pattern []int, direction int) (bool, []int, int) {
// 	new_board 	:= make([]int, common.size * common.size)
// 	zindex		:= current_node.zindex
// 	copy(new_board, current_node.board)

// 	switch direction {
// 		case UP:
// 			if zindex - common.size >= 0 {
// 				new_board[zindex] = new_board[zindex - common.size]
// 				new_board[zindex - common.size] = 0
// 				zindex -= common.size
// 			}
// 		case DOWN:
// 			if zindex + common.size < common.size * common.size  {
// 				new_board[zindex] = new_board[zindex + common.size]
// 				new_board[zindex + common.size] = 0
// 				zindex += common.size
// 			}
// 		case LEFT:
// 			if zindex % common.size >= 1 {
// 				new_board[zindex] = new_board[zindex - 1]
// 				new_board[zindex - 1] = 0
// 				zindex -= 1
// 			}
// 		case RIGHT:
// 			if zindex % common.size < common.size - 1 {
// 				new_board[zindex] = new_board[zindex + 1]
// 				new_board[zindex + 1] = 0
// 				zindex += 1
// 			}
// 	}

// 	return zindex != current_node.zindex, new_board, zindex
// }

func DisjoinctGoal(pattern []int, board []int, common *Common) bool {
	var count 			int	= 0
	var count_pattern 	int = 0

	for key, val := range board {
		if pattern[key] == 1 {
			count_pattern++
			if val == common.goal[key] {
				count++
			}
		}
	}
	return count == count_pattern
}
 
func DisjoinctBFS(common *Common, pattern []int, board []int) {
	var queue *list.List = list.New()
	var node Node

	explored 	:= 0

	ticker 		:= time.NewTicker(1 * time.Second)

	var old_open_count int
	var old_explored int
	go func() {
		for {
			select {
			case <- ticker.C:
				fmt.Println("To explore : ", queue.Len())
				fmt.Println("Explored : ", explored)
				fmt.Println("+ Nodes : ", queue.Len() - old_open_count)
				fmt.Println("- Nodes : ", explored - old_explored)
				old_open_count 	= queue.Len()
				old_explored 	= explored
			}
		}
	}()


	move_cost 	:= make(map[string]int)

	first_node := Node{
		board:        board,
		move:         NONE,
		cost:         0,
		parent_count: 0,
		parent:       nil,
		zindex:       FindIndex(board, 0)}

	queue.PushBack(first_node)
	for queue.Len() != 0 {

		q := queue.Front()
		node = q.Value.(Node)
		queue.Remove(q)	
		move_cost[fmt.Sprint(node.board)] = node.parent_count
		explored++
		if DisjoinctGoal(pattern, node.board, common) {
			fmt.Println("GOAL FOR PATTERN FOUND")
			fmt.Println(pattern)
			fmt.Println(node.board)
			fmt.Println("====")
			break 
		}

		for direction := UP; direction <= RIGHT; direction++ {
			if moved, board, zindex := GenerateNextMoves(common, node, direction); moved {

				if _, exists := move_cost[fmt.Sprint(board)]; !exists {
					pcount := node.parent_count
					if pattern[zindex] == 1 && pattern[zindex] == pattern[node.zindex] {
						pcount += 1
					}
					queue.PushBack(Node {
						board:        board,
						move:         direction,
						cost:         0,
						parent_count: pcount,
						parent:       nil,
						zindex:       zindex,
					})

					move_cost[fmt.Sprint(board)] = pcount
				}

			}
		}

		//fmt.Println(node.board)
	}
}

func DisjoinctDatabaseGenerator(common *Common, board []int) {


	for i := 0; i < common.size * common.size; i++ {

		if pattern_663[0][i] == 1 {
			fmt.Printf("%3d ", common.goal[i])
		} else {
			fmt.Print("  x ")
		}
		if (i + 1) % common.size == 0 {
			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Println()

	for i := 0; i < common.size; i++ {
		for j := 0; j < common.size; j++ {

			index := j + i * common.size

			if pattern_663[0][index] == 1 || pattern_663[1][index] == 1 || pattern_663[2][index] == 1 {
				fmt.Printf("%3d ", index)
			} else {
				fmt.Print("  x ")
			}
		}
		println()
	}

	fmt.Println()

	DisjoinctBFS(common, pattern_663[0], board)
	fmt.Println()
}