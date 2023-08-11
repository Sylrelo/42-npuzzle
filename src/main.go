package main

import "fmt"

func bToMb(b uint64) int {
	return int(b / 1024 / 1024)
}

func main() {
	var common Common
	var initial_board []int
	//var wg 						sync.WaitGroup

	common.size, initial_board = Parse()
	common.goal = GenerateSnail(common.size)

	PrintBoard(initial_board, common.size)

	new_astar(&common, initial_board)
	fmt.Println(initial_board)
	// IDA(&common, initial_board)

}
