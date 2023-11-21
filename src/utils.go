package main

import (
	"fmt"
)

var heurMap = map[string]HeuriFunc{
	"MANHATTAN":      ManhattanDistance,
	"MISPLACED":      HammingDistance,
	"LINEARCONFLICT": LinearConflict,
	"EUCLIDIAN":      EuclideanDistance,
}

func FindIndex(haystack []int8, needle int8) int {
	for i, n := range haystack {
		if needle == n {
			return i
		}
	}
	return -1
}

func Compare(a []int8, b []int8) bool {
	count := 0
	for i, n := range a {
		if b[i] == n {
			count++
		}
	}
	return count == len(a)
}

func PrintBoardOnliner(board []int, size int) {
	for _, n := range board {
		if n == 0 {
			fmt.Print("\033[0;31m")
		}
		fmt.Printf("%3d \033[0m", n)
	}
	fmt.Println("")

}
func PrintBoard(board []int8, size int) {
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			if board[size*j+i] == 0 {
				fmt.Print("\033[0;31m")
			}
			fmt.Printf("%4d\033[0m ", board[size*j+i])
		}
		fmt.Println("")
	}
	fmt.Println("\n ")
}

func GenerateSnail(size int) []int8 {
	snail := make([]int8, 0)
	cur := int8(1)
	x := 0
	ix := 1
	y := 0
	iy := 0

	for i := 0; i < size*size; i++ {
		snail = append(snail, -1)
	}

	for {
		snail[x+y*size] = cur
		if cur == 0 {
			break
		}
		cur += 1
		if x+ix == size || x+ix < 0 || (ix != 0 && snail[x+ix+y*size] != -1) {
			iy = ix
			ix = 0
		} else if y+iy == size || y+iy < 0 || (iy != 0 && snail[x+(iy+y)*size] != -1) {
			ix = -iy
			iy = 0
		}
		x += ix
		y += iy

		if int(cur) == size*size {
			cur = 0
		}
	}

	return snail
}
