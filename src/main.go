package main


func bToMb(b uint64) int {
	return int(b / 1024 / 1024)
}

func main() {
	var common Common
	var initial_board []int

	common.size, initial_board, common.heuristic, common.verbose = Parse()
	common.goal = GenerateSnail(common.size)

	new_astar(&common, initial_board)
	// IDA(&common, initial_board)

}
