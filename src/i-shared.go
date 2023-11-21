package main

import "time"

const (
	NONE            = iota
	UP              = iota
	DOWN            = iota
	LEFT            = iota
	RIGHT           = iota
	MANHATTAN       = iota
	LINEAR_CONFLICT = iota
	HAMMING         = iota
	ASTAR           = iota
)

/*
** Linear Conflict Heuristic optimisation structures
 */

type vec2i struct {
	x int
	y int
}
type LinearConflictHelper struct {
	board []vec2i
	goal  []vec2i
}

/*
**
 */

type HeuristicFunc func(board []int8, goal []int8, size int) int

/*
**
 */

type Node struct {
	parent       *Node
	board        []int8
	move         int
	cost         int
	parent_count int
	zindex       int
}

type HeuriFunc func([]int8, []int8, int) int

type Common struct {
	goal      []int8
	size      int
	heuristic string
	verbose   bool
}

type Result struct {
	node               Node
	time_start         time.Time
	complexity_in_size int
	complexity_in_time int
}
