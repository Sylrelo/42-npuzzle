package main

const (
	NONE  			= iota
	UP    			= iota
	DOWN  			= iota
	LEFT  			= iota
	RIGHT 			= iota
	MANHATTAN		= iota
	LINEAR_CONFLICT = iota
	HAMMING			= iota
	ASTAR			= iota
)


/*
** Linear Conflict Heuristic optimisation structures
*/
type vec2i struct {
	x			int
	y			int
}
type LinearConflictHelper struct {
	board		[]vec2i
	goal		[]vec2i
}
/*
**
*/

type HeuristicFct func(board []int, goal []int, size int) int

/*
**
*/
type Node struct {
	parent       *Node
	board        []int
	move         int
	cost         int
	distance     int
	parent_count int
	zindex       int
}

type Solver struct {
	closed_set 	[][]int
	open_set	PriorityQueue
	heuristic	HeuristicFct
}
type Common struct {
	goal		[]int
	size		int
}


