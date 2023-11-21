package main

import (
	"fmt"
	"math/big"
	"regexp"
	"runtime"
	"time"
)

func formatNumber(num int) string {
	str := fmt.Sprintf("%d", num)
	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != str; {
		n = str
		str = re.ReplaceAllString(str, "$1 $2")
	}
	return str
}

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}

func SolutionFound(common *Common, result Result, algo string) {
	var m runtime.MemStats
	time_elapsed := time.Since(result.time_start)

	runtime.ReadMemStats(&m)

	fmt.Print("\033[H\033[2J")
	fmt.Printf("\033[1;32m%s\033[0m\n", "√ Solution found")

	fmt.Printf("\033[1;33m%-18s : \033[0m%s\n", "Algorithm", algo)
	fmt.Printf("\033[1;33m%-18s : \033[0m%s\n\n", "Heuristic", common.heuristic)

	fmt.Printf("\033[1m%-18s : \033[0m%.4fs\n", "Time taken", time_elapsed.Seconds())
	fmt.Printf("\033[1m%-18s : \033[0m%s\n", "Complexity in size", formatNumber(result.complexity_in_size))
	fmt.Printf("\033[1m%-18s : \033[0m%s\n\n", "Complexity in time", formatNumber(result.complexity_in_time))

	fmt.Printf("\033[1m\x1b[38;2;40;177;249m%-18s : \033[0m%d\n", "Puzzle size", common.size*common.size)
	fmt.Printf("\033[1m\x1b[38;2;40;177;249m%-18s : \033[0m%s\n\n", "Possible states", new(big.Int).Div(factorial(big.NewInt(int64(common.size)*int64(common.size))), big.NewInt(2)))

	fmt.Printf("\033[1m\x1b[38;2;140;77;249m%-18s : \033[0m%6s Mb\n", "Allocated", formatNumber(bToMb(m.Alloc)))
	fmt.Printf("\033[1m\x1b[38;2;140;77;249m%-18s : \033[0m%6s Mb\n\n", "Total allocated", formatNumber(bToMb(m.TotalAlloc)))

	fmt.Printf("\033[1m\x1b[38;2;40;177;249m%-18s : \033[0m%d\n", "Number of moves", result.node.parent_count)
	fmt.Printf("\033[1m\x1b[38;2;40;177;249m%-18s:\033[0m \n", "Solution :")
}

func GenerateNextMoves(common *Common, current_node Node, direction int) (bool, []int8, int) {
	new_board := make([]int8, common.size*common.size)
	zindex := current_node.zindex
	copy(new_board, current_node.board)

	switch direction {
	case UP:
		if zindex-common.size >= 0 {
			new_board[zindex] = new_board[zindex-common.size]
			new_board[zindex-common.size] = 0
			zindex -= common.size
		}
	case DOWN:
		if zindex+common.size < common.size*common.size {
			new_board[zindex] = new_board[zindex+common.size]
			new_board[zindex+common.size] = 0
			zindex += common.size
		}
	case LEFT:
		if zindex%common.size >= 1 {
			new_board[zindex] = new_board[zindex-1]
			new_board[zindex-1] = 0
			zindex -= 1
		}
	case RIGHT:
		if zindex%common.size < common.size-1 {
			new_board[zindex] = new_board[zindex+1]
			new_board[zindex+1] = 0
			zindex += 1
		}
	}

	return zindex != current_node.zindex, new_board, zindex
}

func printMove(node Node, size int, verbose bool) {
	switch node.move {
	case NONE:
		fmt.Print("INITIAL")
	case UP:
		fmt.Print("BAS")
	case DOWN:
		fmt.Print("HAUT")
	case LEFT:
		fmt.Print("DROITE")
	case RIGHT:
		fmt.Print("GAUCHE")
	}
	if verbose {
		fmt.Println("")
		PrintBoard(node.board, size)
	} else {
		fmt.Print(" > ")
	}
}

func GenerateHistory(common *Common, node Node) {
	var nodes []Node
	var reversed_nodes []Node

	tmp := node
	for {
		nodes = append(nodes, tmp)
		if tmp.parent == nil {
			break
		}
		tmp = *tmp.parent
	}
	for i := range nodes {
		n := nodes[len(nodes)-1-i]
		reversed_nodes = append(reversed_nodes, n)
		if common.verbose {
			printMove(n, common.size, true)
		}
	}
	for _, n := range reversed_nodes {
		printMove(n, common.size, false)
	}
	fmt.Print("The End 👰‍♀️💍👰💍👰‍♂️  🍼👶👧🧒👦👩🧑🧑‍🦲🛝\n")
}
