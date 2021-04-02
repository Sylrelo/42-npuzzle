package main

import (
	"fmt"
	"runtime"
)

/*
func GenerateHistory(common *Common, node Node) {
	var nodes			[]Node
	var reversed_nodes	[]Node

	tmp := node
	for {
		PrintBoard(tmp.board, common.size)
		nodes = append(nodes, tmp)
		if (tmp.parent == nil) {
			break
		}
		tmp = *tmp.parent
	}
	for i := range nodes {
		n := nodes[len(nodes) - 1 - i]
		reversed_nodes = append(reversed_nodes, n)
	}
	for _, n := range reversed_nodes {
		switch n.move {
			case NONE:
				fmt.Print("INITIAL")
			case UP:
				fmt.Print("UP")
			case DOWN:
				fmt.Print("DOWN")
			case LEFT:
				fmt.Print("LEFT")
			case RIGHT:
				fmt.Print("RIGHT")
		}
		fmt.Print(" > ")
	}
	fmt.Print("\n")
}

func GetHeuristicName(heuristic int) string {
	switch heuristic {
		case MANHATTAN:
			return "Manhattan Distance"
		case HAMMING:
			return "Hamming Distance (Missplaced Tiles)"
		case LINEAR_CONFLICT:
			return "Linear Conflict (+ Manhattan Distance)"
		default:
			return "---"
	}
}
*/
func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}

func main() {
	var common					Common
	var initial_board 			[]int
	//var wg 						sync.WaitGroup

	common.size, initial_board 	= Parse()
	common.goal					= GenerateSnail(common.size)

	PrintBoard(initial_board, common.size)

	//wg.Add(3)
	//go Solve(&wg, common, initial_board, ASTAR, HAMMING)
	//go Solve(&wg, common, initial_board, ASTAR, MANHATTAN)
	//go Solve(&wg, common, initial_board, ASTAR, LINEAR_CONFLICT)
	//wg.Wait()
	
	new_astar(&common, initial_board)

	var m runtime.MemStats
        runtime.ReadMemStats(&m)
        fmt.Printf("\tAlloc = %v MiB\n", bToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
        fmt.Printf("\tSys = %v MiB\n", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
		
	
}
