package main

import (
	"container/list"
	_"container/heap"
)

const (
	NONE	= iota
	UP 		= iota
	DOWN 	= iota
	LEFT 	= iota
	RIGHT 	= iota
)
const (
	NSIZE = 9
	NCOL 	= 3
)

type Node struct {
	parent		*Node
	board		[9]int
	move		int
	cost		int
	distance	int
}

type Solver struct {
	open		*list.List
	closed		*list.List
	ncol		int
}