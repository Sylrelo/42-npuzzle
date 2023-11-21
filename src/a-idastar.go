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

func IDA(common *Common, board []int8) {
	node := Node{
		board:        board,
		move:         NONE,
		cost:         0,
		parent_count: 0,
		parent:       nil,
		zindex:       FindIndex(board, 0)}


	threshold	:= heurMap[common.heuristic](board, common.goal, common.size)
	path		:= New()

	path.Push(node)

	for {
		result, nbound, node := IDA_Start(common, node, 0, threshold, 0)

		if result == ERROR {
			fmt.Println("ERROR")
		}
		if result == FOUND {
			fmt.Println("C'est trouvé !")
			GenerateHistory(common, node)
			os.Exit(1)
		}
		if result == NOTHING && nbound == int(^uint(0) >> 1) {
			fmt.Println("C'est un échec !")
			os.Exit(1)
		}
		
		fmt.Println("Updating bound ", threshold, nbound)
		
		threshold = nbound
	}
}

func IDA_Start(common *Common, node Node, cost int, threshold int, oy int) (int, int, Node) {

	
	ncost	:= node.parent_count + cost + heurMap[common.heuristic](node.board, common.goal, common.size)
	min		:= int(^uint(0) >> 1)

	if ncost > threshold {
		return NOTHING, ncost, node
	}

	if Compare(common.goal, node.board) {
		return FOUND, 0, node
	}
	for _, successor := range IDA_NextMoves(common, node) {
		successor_cost := cost + heurMap[common.heuristic](node.board, successor.board, common.size)
		result, nbound, node := IDA_Start(common, successor, successor_cost, threshold, oy + 1)
		if result == FOUND {
			return FOUND, 0, node
		}
		if nbound < min {
			min = nbound
		}
	}
	return NOTHING, min, node
}


