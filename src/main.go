package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
	"time"
)

func NewNode(common *Common, solver *Solver, current_node Node, new_board []int, zindex int, direction int) {
	//if _, exists := solver.closed_set2[fmt.Sprint(new_board)]; exists {
	//	return
	//}
	if Same(solver.closed_set, new_board, common.size * common.size) {
		return
	}
	priority := current_node.parent_count + solver.heuristic(new_board, common.goal, common.size)
	//priority :=  current_node.cost 
	new_node := Node{
		board:        new_board,
		move:         direction,
		cost:         priority,
		parent_count: current_node.parent_count + 1,
		distance:     priority,
		parent:       &current_node,
		zindex:       zindex}

	heap.Push(&solver.open_set, &Item{node: new_node, priority: priority})
	//fmt.Println("\033[1;36m+ Queue push\033[0m")
}

func Move(common *Common, solver *Solver, current_node Node, direction int) {
	new_board := make([]int, common.size * common.size)
	copy(new_board, current_node.board)

	switch direction {
		case UP:
			if current_node.zindex - common.size >= 0 {
				//fmt.Println("UP " , current_node.zindex, current_node.zindex - common.size)
				new_board[current_node.zindex] = new_board[current_node.zindex - common.size]
				new_board[current_node.zindex - common.size] = 0
				NewNode(common, solver, current_node, new_board, current_node.zindex - common.size, direction)
			}
		case DOWN:
			if current_node.zindex + common.size < common.size * common.size  {
				//fmt.Println("DOWN ", current_node.zindex, current_node.zindex + 1)
				new_board[current_node.zindex] = new_board[current_node.zindex + common.size]
				new_board[current_node.zindex + common.size] = 0
				NewNode(common, solver, current_node, new_board, current_node.zindex + common.size, direction)
			}
		case LEFT:
			if current_node.zindex % common.size >= 1 {
				new_board[current_node.zindex] = new_board[current_node.zindex - 1]
				new_board[current_node.zindex - 1] = 0
				NewNode(common, solver, current_node, new_board, current_node.zindex - 1, direction)
			}
		case RIGHT:
			if current_node.zindex % common.size < common.size - 1 {
				new_board[current_node.zindex] = new_board[current_node.zindex + 1]
				new_board[current_node.zindex + 1] = 0
				NewNode(common, solver, current_node, new_board, current_node.zindex + 1, direction)
			}
		default:
			fmt.Print("--")
	}
}

func MoveUNIT(board []int, size int, direction int) {
	new_board := make([]int, size * size)
	copy(new_board, board)


	zindex := FindIndex(board, 0)

	fmt.Println(direction, zindex)

	switch direction {
		case UP:
			if zindex - size >= 0 {
				new_board[zindex] = new_board[zindex - size]
				new_board[zindex - size] = 0
			}
		case DOWN:
			if zindex + size < size * size  {
				new_board[zindex] = new_board[zindex + size]
				new_board[zindex + size] = 0
			}
		case LEFT:
			if zindex % size >= 1 {
				new_board[zindex] = new_board[zindex - 1]
				new_board[zindex - 1] = 0
			}
		case RIGHT:
			fmt.Println(zindex, size, zindex % size)
			if zindex % size < size - 1 {
				new_board[zindex] = new_board[zindex + 1]
				new_board[zindex + 1] = 0
			}
		default:
			fmt.Print("--")
	}
	PrintBoard(new_board, size)
	PrintBoardOnliner(board, size)
	PrintBoardOnliner(new_board, size)
	fmt.Println()
}

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

func BreadthFirstSearch(common Common, solver *Solver, node Node) (bool, Node) {
	_ = common
	_ = solver
	_ = node
	return false, Node{}
}

func IDAstar(common Common, solver *Solver, node Node) (bool, Node) {

	fmt.Println("Starting IDA*")
	//solver.closed_set
	solver.lifo_stack = New()

	//cost := ManhattanDistance(node.board, common.goal, common.size) + 2 * LinearConflict(node.board, common.goal, common.size)


	var current_node 	*Node
	var cost			int
	var bound			int

	iter_n 	:= 0
	//min		:= math.Inf()
	//var ret Node
	for {

		current_node = solver.lifo_stack.Pop()
		cost := solver.heuristic(current_node.board, common.goal, common.size)
		bound = cost

		if cost > bound {
			continue 
		}

		if Compare(current_node.board, common.goal) {
			break 
		}



		//if ret.cost > 150 {
		//	break
		//}
		fmt.Println(iter_n)
		iter_n++
	}
	os.Exit(1)

	_ = cost
	// for currnode := solver.lifo_stack.Pop(); currnode != nil; {

		
	// 	fmt.Println(currnode)
	// }

	_ = common
	_ = solver
	_ = node
	return false, Node{}
}

func Astar(common Common, solver *Solver, node Node) (bool, Node) {
	for {
	//	fmt.Print("\033[H\033[2J")
		//time.Sleep(1000000)
		if solver.open_set.Len() == 0 {
			fmt.Println("\033[1;31mEmpty queue, break.\033[0m")
			break
		}

		solver.complexity_in_size 	= math.Max(float64(solver.complexity_in_size), float64(solver.open_set.Len()))
		node 						:= heap.Pop(&solver.open_set).(*Item).node
		solver.closed_set 			= append(solver.closed_set, node.board)

		//fmt.Println(int64(node.board))

		//fmt.Println(fmt.Sprint(node.board))
		
		//solver.closed_set2[fmt.Sprint(node.board)] = struct{}{}


		if Compare(node.board, common.goal) {
			return true, node
		}
		//PrintBoardOnliner(node.board, common.size)

		//fmt.Println(solver.open_set.Len(), len(solver.closed_set))
		// if (common.size >= 4 && CompareRocol(node.board, common.goal, common.size, solver.index)) {
		// 	PrintBoard(node.board, common.size)
		// 	solver.index++
		// }

		Move(&common, solver, node, UP)
		Move(&common, solver, node, DOWN)
		Move(&common, solver, node, LEFT)
		Move(&common, solver, node, RIGHT)

		//os.Exit(1)
	}
	return false, Node{}
}

func Solve(wg *sync.WaitGroup, common Common, intial_board []int, algo int, heuristic int) {
	defer wg.Done()

	var node 					Node
	var	solver					Solver
	var solution				bool
	var solution_node			Node

	solver.open_set 			= make(PriorityQueue, 0)
	solver.closed_set 			= make([][]int, 0)
	solver.complexity_in_size	= 0
	solver.index				= 0
	solver.closed_set2			= make(map[string]struct{})

	// if (common.size > 3) {
	// 	solver.index = 1
	// }

	old_closed_count := 0
	old_open_count := 0

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <- ticker.C:
				fmt.Println("Explored nodes : ", len(solver.closed_set))
				fmt.Println("Opens nodes : ", solver.open_set.Len())
				fmt.Println(len(solver.closed_set) - old_closed_count, " nodes / s" )
				fmt.Println(solver.open_set.Len() - old_open_count, " nodes / s" )
				//old_closed_count = len(solver.closed_set2)
				old_closed_count = len(solver.closed_set)
				old_open_count = solver.open_set.Len()
			}
		}
	}()

	solution					= false

	switch heuristic {
		case MANHATTAN:
			solver.heuristic = ManhattanDistance
		case HAMMING:
			solver.heuristic = HammingDistance
		case LINEAR_CONFLICT:
			solver.heuristic = LinearConflict
		default:
			fmt.Println("Wrong heuristic")
			return 
	}

	node = Node{
		board:        intial_board,
		move:         NONE,
		cost:         0,
		parent_count: 0,
		distance:     0,
		parent:       nil,
		zindex:       FindIndex(intial_board, 0)}

	heap.Push(&solver.open_set, &Item{node: node, priority: 0})

	time_start 				:= time.Now()

	switch algo {
		case ASTAR:
			solution, solution_node = Astar(common, &solver, node)
		default:
			solution, solution_node = false, Node{}
			return
	}
	
	if solution {
		time_elapsed := time.Since(time_start)
		fmt.Printf("\033[1;32m%s\033[0m\n", GetHeuristicName(heuristic))
		fmt.Printf("> %-18s : %6d\n", "Parents", solution_node.parent.parent_count)
		fmt.Printf("> %-18s : %6d\n", "Complexity in time", len(solver.closed_set))
		fmt.Printf("> %-18s : %6d\n", "Complexity in size", int(solver.complexity_in_size))
		fmt.Printf("> %-18s : %6.3fs\n", "Time taken", time_elapsed.Seconds())
		time.Sleep(100000 * 1000)
	}
}


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
	//testBoard := []int{11, 11, 11, 11,
	//					11, 0, 11, 11,
	//					11, 11, 11, 11,
	//					11, 11, 11, 11}
	// PrintBoard(testBoard, 4)

	// MoveUNIT(testBoard, 4, UP)
	// MoveUNIT(testBoard, 4, DOWN)
	// MoveUNIT(testBoard, 4, LEFT)
	// MoveUNIT(testBoard, 4, RIGHT)

	// os.Exit(1)
	//wg.Add(3)
	//go Solve(&wg, common, initial_board, ASTAR, HAMMING)
	//go Solve(&wg, common, initial_board, ASTAR, MANHATTAN)
	//go Solve(&wg, common, initial_board, ASTAR, LINEAR_CONFLICT)
	//wg.Wait()
//
	//openhash 	:= make(map[string]Node)
	//closedhash 	:= make(map[string]Node)
//
	//test2 	:= make([]Node, 0)
	//lala 	:= make(PriorityQueue, 0)

	for i := 0; i < 10000000; i++ {
		//openhash[fmt.Sprint(i * 1000000)] = Node{
		//	board: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		//	cost: i,
		//	distance: i,
		//	parent_count: i,
		//}

		//closedhash[fmt.Sprint(i * 1000000)] = Node{}
		//heap.Push(&lala, &Item{node: Node{}, priority: i})

	}	

	//_ = openhash
	//_ = closedhash
	//_ = lala
	//_ = test2

	new_astar(&common, initial_board)

	var m runtime.MemStats
        runtime.ReadMemStats(&m)
        // For info on each, see: https://golang.org/pkg/runtime/#MemStats
        fmt.Printf("\tAlloc = %v MiB\n", bToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
        fmt.Printf("\tSys = %v MiB\n", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
		
	os.Exit(1)

	node := Node{
		board:        initial_board,
		move:         NONE,
		cost:         0,
		parent_count: 0,
		distance:     0,
		parent:       nil,
		zindex:       FindIndex(initial_board, 0)}

	var solver Solver
	solver.heuristic = LinearConflict
	IDAstar(common, &solver, node)
}
