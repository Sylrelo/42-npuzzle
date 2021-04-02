package main

import (
	"container/heap"
)

/*
** PriorityQueue Heap
 */

type Item struct {
	node     Node
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] 	= pq[j], pq[i]
	pq[i].index 	= i
	pq[j].index 	= j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n 			:= len(*pq)
	item 		:= x.(*Item)
	item.index 	= n
	*pq 		= append(*pq, item)

}

func (pq *PriorityQueue) Pop() interface{} {
	old 		:= *pq
	n 			:= len(old)
	item 		:= old[n - 1]
	old[n - 1] 	= nil
	item.index 	= -1
	*pq 		= old[0 : n - 1]
	return item
}

func (pq *PriorityQueue) update(item *Item, node Node, priority int) *Item {
	item.node = node
	item.priority = priority
	heap.Fix(pq, item.index)
	return item
}

/*
** LIFO Stack
*/

type LIFO struct {
	nodes []Node
	count int
}

func New() *LIFO {
	return &LIFO{}
}

func (s *LIFO) Push(n Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *LIFO) Pop() Node {
	if s.count == 0 {
		return Node{} 
	}
	s.count--
	return s.nodes[s.count]
}

func (s *LIFO) Len() int {
	return s.count
}

func (s *LIFO) Last() Node {
	if s.count == 0 {
		return Node{} 
	}
	return s.nodes[s.count - 1]
}
