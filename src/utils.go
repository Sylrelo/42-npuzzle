package main

import "fmt"

func FindIndexLinearConflict(goal []int, board []int, g int, b int) [4]int {
	ret := [4]int{0, 0, 0, 0}

	return ret
}

func FindIndex(haystack []int, needle int) int {
	for i, n := range haystack {
		if needle == n {
			return i
		}
	}
	return -1
}

func Same(a [][]int, b []int) bool {
	for _, set := range a {
		count := 0
		for i, n := range set {
			if b[i] == n {
				count++
			}
		}
		if count == 9 {
			return true
		}
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

func PrintBoard(board []int, size int) {
	for j := 0; j < size; j++ {
		for i := 0; i < size; i++ {
			fmt.Print(board[size*j+i], " ")
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
