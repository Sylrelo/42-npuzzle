package main

import (
	"container/heap"
	"fmt"
	"math"
	"sync"
	"time"
)

func NewNode(common *Common, solver *Solver, current_node Node, new_board []int, zindex int, direction int) {
	if Same(solver.closed_set, new_board) {
		//fmt.Println("\033[33m~ Board already explored. Skipping.\033[0m")
		return
	}

	priority := current_node.parent_count + solver.heuristic(new_board, common.goal, common.size)
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
	new_board := make([]int, len(current_node.board))
	copy(new_board, current_node.board)

	switch direction {
		case UP:
			if current_node.zindex - common.size >= 0 {
				new_board[current_node.zindex] = new_board[current_node.zindex - common.size]
				new_board[current_node.zindex - common.size] = 0
				NewNode(common, solver, current_node, new_board, current_node.zindex - common.size, direction)
			}
		case DOWN:
			if current_node.zindex + common.size < common.size * common.size {
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
			if current_node.zindex % common.size <= 1 {
				new_board[current_node.zindex] = new_board[current_node.zindex + 1]
				new_board[current_node.zindex + 1] = 0
				NewNode(common, solver, current_node, new_board, current_node.zindex + 1, direction)
			}
		default:
			fmt.Print("--")
	}
}

func GenerateHistory(node Node) {
	var nodes			[]Node
	var reversed_nodes	[]Node

	tmp := node
	for {
		//PrintBoard(tmp.board, Size{9, 3})
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
	_ = common
	_ = solver
	_ = node
	return false, Node{}
}

func Astar(common Common, solver *Solver, node Node) (bool, Node) {
	for {
		if solver.open_set.Len() == 0 {
			fmt.Println("\033[1;31mEmpty queue, break.\033[0m")
			break
		}

		solver.complexity_in_size 	= math.Max(float64(solver.complexity_in_size), float64(solver.open_set.Len()))
		node 						:= heap.Pop(&solver.open_set).(*Item).node
		solver.closed_set 			= append(solver.closed_set, node.board)

		if Compare(node.board, common.goal) {
			return true, node
		}

		Move(&common, solver, node, UP)
		Move(&common, solver, node, DOWN)
		Move(&common, solver, node, LEFT)
		Move(&common, solver, node, RIGHT)
	}
	return false, Node{}
}

func Solve(wg *sync.WaitGroup, common Common, intial_board []int, algo int, heuristic int) {
	defer wg.Done()

	//var complexity_in_size	float64
	var node 				Node
	var	solver				Solver
	var solution			bool
	var solution_node		Node

	solver.open_set 			= make(PriorityQueue, 0)
	solver.closed_set 			= make([][]int, 0)
	solver.complexity_in_size	= 0

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


func main() {
	var common					Common
	var initial_board 			[]int
	var wg 						sync.WaitGroup

	common.size, initial_board 	= Parse()
	common.goal					= GenerateSnail(common.size)

	PrintBoard(initial_board, common.size)

	wg.Add(3)
	go Solve(&wg, common, initial_board, ASTAR, HAMMING)
	go Solve(&wg, common, initial_board, ASTAR, MANHATTAN)
	go Solve(&wg, common, initial_board, ASTAR, LINEAR_CONFLICT)
	wg.Wait()
}
