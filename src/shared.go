package main

import (
	_ "container/heap"
	_ "container/list"
)

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

type Node struct {
	parent       *Node
	board        []int
	move         int
	cost         int
	distance     int
	parent_count int
	zindex       int
}

type Size struct {
	nsize int
	ncol  int
}
