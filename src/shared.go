package main

const (
	NONE  = iota
	UP    = iota
	DOWN  = iota
	LEFT  = iota
	RIGHT = iota
)

const (
	NSIZE = 9
	NCOL  = 3
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
type Node struct {
	parent       *Node
	board        []int
	move         int
	cost         int
	distance     int
	parent_count int
	zindex       int
}

type Common struct {
	closed_set 	[][]int
	open_set	PriorityQueue
	goal		[]int
	size		int
}


