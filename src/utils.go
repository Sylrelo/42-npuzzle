package main

import (
	"fmt"
)

func FindIndex(haystack []int, needle int) int {
	for i, n := range haystack {
		if needle == n {
			return i
		}
	}
	return -1
}

func Same(closed [][]int, board []int, size int) bool {
	for _, set := range closed {
		count := 0
		for i, n := range set {
			if board[i] == n {
				count++
			} else {
				break
			}
		}
		if count == size {
			return true
		}
	_ = set
}
	return false
}

func Compare(a []int, b []int) bool {
	count := 0
	for i, n := range a {
		if b[i] == n {
			count++
		}
	}
	return count == len(a)
}

func CompareRocol(a []int, b []int, size int, index int) bool {
	count := 0

	for row := index + 1; row < size; row++ {
		//fmt.Print(a[row + index * size], "  ")
		if a[row + index * size] == b[row + index * size] {
			count++
		}
	}

	for col := index; col < size; col++ {
		//fmt.Print(a[index + col * size], "  ")
		if a[index + col * size] == b[index + col * size] {
			count++
		}
	}
	return count == size * 2
}


func PrintBoardOnliner(board []int, size int) {
	for _, n := range board {
		if n == 0 {
			fmt.Print("\033[0;31m")
		}
		fmt.Printf("%3d \033[0m" , n)
	}
	fmt.Println("")

}
func PrintBoard(board []int, size int) {
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			if board[size * j + i] == 0 {
				fmt.Print("\033[0;31m")
			}
			fmt.Printf("%4d\033[0m ", board[size * j + i])
		}
		fmt.Println("")
	}

	// for i := 0; i < size; i++ {
	// 	fmt.Print(board[(size)+i], " ")
	// }
	// fmt.Println("")
	// for i := 0; i < size; i++ {
	// 	fmt.Print(board[(size*2)+i], " ")
	// }
	fmt.Println("\n ")
}

func GenerateSnail(size int) []int{
	snail := make([]int, 0)
	cur := 1
	x 	:= 0
	ix 	:= 1
	y 	:= 0
	iy 	:= 0

	for i := 0; i < size * size; i++ {
		snail = append(snail, -1)
	}

	for {
		snail[x + y * size] = cur
		if cur == 0 {
			break 
		}
		cur += 1
		if x + ix == size || x + ix < 0 || (ix != 0 && snail[x + ix + y * size] != -1) {
			iy = ix
			ix = 0
		} else if y + iy == size || y + iy < 0 || (iy != 0 && snail[x + (iy + y) * size] != -1) {
			ix = -iy
			iy = 0
		}
		x += ix
		y += iy

		if cur == size * size {
			cur = 0
		}
	}

	return snail
}
