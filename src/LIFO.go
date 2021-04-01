package main

type LIFO struct {
	nodes []*Node
	count int
}

func New() *LIFO {
	return &LIFO{}
}

func (s *LIFO) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *LIFO) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func (s *LIFO) Len() int {
	return s.count
}
