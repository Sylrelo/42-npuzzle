package main

import (
	"fmt"
	"os"
)

const (
	NOTHING		= iota
	ERROR		= iota
	FOUND 		= iota
	NOT_FOUND 	= iota
)

func IDA(common *Common, board []int) {
	node := Node{
		board:        board,
		move:         NONE,
		cost:         0,
		parent_count: 0,
		parent:       nil,
		zindex:       FindIndex(board, 0)}


	hash 		:= make(map[string]struct{})
	bound		:= LinearConflict(board, common.goal, common.size)
	path		:= New()

	path.Push(node)
	hash[fmt.Sprint(board)] = struct{}{}

	for {
		result, nbound := IDA_Start(common, path, hash, 0, bound, 0)

		if result == ERROR {
			fmt.Println("ERROR")
			//os.Exit(1)
		}
		if result == FOUND {
			fmt.Println("C'est trouvé !")
			fmt.Println(result)
			os.Exit(1)
		}
		if result == NOTHING && nbound == int(^uint(0) >> 1) {
			fmt.Println("C'est un échec !")
			os.Exit(1)
		}
		
		fmt.Println("Updating bound ", bound, nbound)
		
		bound = nbound
	}
}

func IDA_NextMoves(common *Common, node Node) []Node {
	nodes := make([]Node, 0)

	for dir := UP; dir <= RIGHT; dir++ {
		if moved, board, zindex := GenerateNextMoves(common, node, dir); moved {
			nodes = append(nodes, Node{
				board:        board,
				move:         dir,
				cost:         0,
				parent_count: node.parent_count + 1,
				parent:       &node,
				zindex:       zindex})
		}
	}
	return nodes
}

func IDA_Start(common *Common, stack *LIFO, hash map[string]struct{}, cost int, bound int, oy int) (int, int) {


	//if stack.Len() == 0 {
	//	return ERROR, cost
	//}

	node	:= stack.Last()
	strb := fmt.Sprint(node.board)
	delete(hash, strb)
	//fmt.Println(node.board)
	//fmt.Println("IDA_Start")

	ncost	:= node.parent_count + LinearConflict(node.board, common.goal, common.size)
	min		:= int(^uint(0) >> 1)
	//fmt.Println("IDA_Start - ", oy, min)

	if ncost > bound {
		return NOTHING, ncost
	}

	if Compare(common.goal, node.board) {
		fmt.Println(node.parent_count, node.board)
		return FOUND, 0
	}
	//fmt.Println("IDA_Start - 1")

	for _, successor := range IDA_NextMoves(common, node) {
		nstrb := fmt.Sprint(successor.board)
		_, exists := hash[nstrb]
		if exists {
			fmt.Println("==>>>>>>>>>>>>> Exists")
		}
		if !exists {
			//fmt.Println(exists, nstrb)
			stack.Push(successor)
			hash[nstrb] = struct{}{}
			//fmt.Println(i, " Push => ")
			result, nbound := IDA_Start(common, stack, hash, node.parent_count + LinearConflict(node.board, successor.board, common.size), bound, oy + 1)
			//if result == ERROR {
			//	continue 
			//}
			if result == FOUND {
				return FOUND, 0
			}
			if nbound < min {
				min = nbound
				//fmt.Println("EDIT ", nbound, min)
			}
			node = stack.Pop()
			delete(hash, fmt.Sprint(node.board))
			//fmt.Println("Pop <= ")

		}
	}
	return NOTHING, min
}
