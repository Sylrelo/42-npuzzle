package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func NewNode(common *Common, current_node Node, new_board []int, zindex int, direction int) {
	if Same(common.closed_set, new_board) {
		//fmt.Println("\033[33m~ Board already explored. Skipping.\033[0m")
		return
	}

	priority := current_node.parent_count + ManhattanDistance(new_board, common.goal, common.size) + (2 * LinearConflict(new_board, common.goal, common.size))
	new_node := Node{
		board:        new_board,
		move:         direction,
		cost:         priority,
		parent_count: current_node.parent_count + 1,
		distance:     priority,
		parent:       &current_node,
		zindex:       zindex}

	heap.Push(&common.open_set, &Item{node: new_node, priority: priority})
	//fmt.Println("\033[1;36m+ Queue push\033[0m")
}

func Move(common *Common, current_node Node, direction int) {
	new_board := make([]int, len(current_node.board))
	copy(new_board, current_node.board)

	switch direction {
		case UP:
			if current_node.zindex - common.size >= 0 {
				new_board[current_node.zindex] = new_board[current_node.zindex - common.size]
				new_board[current_node.zindex - common.size] = 0
				NewNode(common, current_node, new_board, current_node.zindex - common.size, direction)
			}
		case DOWN:
			if current_node.zindex + common.size < NSIZE {
				new_board[current_node.zindex] = new_board[current_node.zindex + common.size]
				new_board[current_node.zindex + common.size] = 0
				NewNode(common, current_node, new_board, current_node.zindex + common.size, direction)
			}
		case LEFT:
			if current_node.zindex % common.size >= 1 {
				new_board[current_node.zindex] = new_board[current_node.zindex - 1]
				new_board[current_node.zindex - 1] = 0
				NewNode(common, current_node, new_board, current_node.zindex - 1, direction)
			}
		case RIGHT:
			if current_node.zindex % common.size <= 1 {
				new_board[current_node.zindex] = new_board[current_node.zindex + 1]
				new_board[current_node.zindex + 1] = 0
				NewNode(common, current_node, new_board, current_node.zindex + 1, direction)
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

func ParseContent(content string) (int, []int) {
	var size		int
	var result		[]int

	size				= -1
	content_splitted 	:= strings.Split(content, "\n")

	for line_number, line := range content_splitted {
		if len(line) < 1 || line[0] == '#' {
			continue
		}
		if size == -1 {
			atoi, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error puzzle size, line", line_number + 1)
				os.Exit(0)
			}
			size = atoi
		} else {
			split_values := strings.Split(line, " ")
			for _, nb := range split_values {
				tmp := strings.TrimSpace(nb)
				if len(tmp) > 0 {
					atoi, err := strconv.Atoi(tmp)
					if err != nil {
						fmt.Println("Error puzzle, line", line_number + 1)
						os.Exit(0)
					}
					result = append(result, atoi)
				}
			}
		}
	}
	return size, result
}

func Parse() (int, []int) {
	args := os.Args[1:]

	if len(args) > 0 {
		if (args[0] == "-f") {
			if len(args) > 1 {
				data, err := ioutil.ReadFile(args[1])
				if err != nil {
					fmt.Println("Error reading file", args[1])
					os.Exit(1)
				}
				return ParseContent(string(data))
			} else {
				fmt.Println("You must specify a file")
				os.Exit(1)
			}
		}
	} else {
		reader 		:= bufio.NewReader(os.Stdin)
		content, _ 	:= reader.ReadString(0)
		return ParseContent(content)
	}
	return -1, nil
}

func main() {
	var common				Common
	var base 				[]int
	var node 				Node
	var complexity_in_size	float64

	common.open_set 		= make(PriorityQueue, 0)
	common.size, base 		= Parse()
	common.goal				= GenerateSnail(common.size)

	PrintBoard(base, common.size)

	time_start := time.Now()
	
	node = Node{
		board:        base,
		move:         NONE,
		cost:         0,
		parent_count: 0,
		distance:     0,
		parent:       nil,
		zindex:       FindIndex(base, 0)}

	heap.Push(&common.open_set, &Item{node: node, priority: 0})

	max_iterations := 0
	for max_iterations < 375000 {

		if common.open_set.Len() == 0 {
			fmt.Println("\033[1;31mEmpty queue, break.\033[0m")
			break
		}
		//fmt.Println("\033[1;34m- Queue pop.\033[0m")
		complexity_in_size 	= math.Max(float64(complexity_in_size), float64(common.open_set.Len()))

		node 				:= heap.Pop(&common.open_set).(*Item).node
		common.closed_set 	= append(common.closed_set, node.board)

		if Compare(node.board, common.goal) {
			fmt.Println("\033[1;34mSolution Found !\033[0m")
			fmt.Println(node.parent.parent_count, " parents")
			fmt.Println(len(common.closed_set), " complexity in time (closed)")
			fmt.Println(complexity_in_size, " complexity in size (max open)")
			//GenerateHistory(node)
			break 
		}
		Move(&common, node, UP)
		Move(&common, node, DOWN)
		Move(&common, node, LEFT)
		Move(&common, node, RIGHT)
		max_iterations++
	}
	time_elapsed := time.Since(time_start)

	fmt.Println(len(common.closed_set), " complexity in time (closed)")
	fmt.Println(complexity_in_size, " complexity in size (max open)")

	fmt.Printf("Time taken %s \n", time_elapsed)
}
