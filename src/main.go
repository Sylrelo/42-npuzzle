package main

func bToMb(b uint64) int {
    return int(b / 1024 / 1024)
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
	
	//new_astar(&common, initial_board)
	//IDA(&common, initial_board)

	DisjoinctDatabaseGenerator(&common, initial_board)
	//var m runtime.MemStats
    //    runtime.ReadMemStats(&m)
    //    fmt.Printf("\tAlloc = %v MiB\n", bToMb(m.Alloc))
    //    fmt.Printf("\tTotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
    //    fmt.Printf("\tSys = %v MiB\n", bToMb(m.Sys))
    //    fmt.Printf("\tNumGC = %v\n", m.NumGC)
		
	
}
