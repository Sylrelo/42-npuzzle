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
			fmt.Print(board[size * j + i], " ")
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

// func GenerateSnail(col int, row int) []int {
// 	result 		:= make([]int, col * row)
// 	//direction 	:= RIGHT
// 	irow		:= 0
// 	crow		:= 0
// 	dir			:= 1
// 	//icol		:= 0
// 	//istart		:= 0

// 	for i := col * row - 1; i >= 0; i-- {

// 		irow = 0
// 		for irow < (col * row) {

// 			crow += dir
// 			irow++

// 			dir *= -1
// 		}
// 		// result[i - 1] = 0
// 		// if direction == RIGHT && icol < col {
// 		// 	result[i - 1] = i
// 		// 	icol++
// 		// }
// 		// if icol == 3 && direction == RIGHT {
// 		// 	direction = DOWN
// 		// 	irow = 0
// 		// 	istart = row + icol
// 		// }
// 		// if direction == DOWN && irow <= 1 {
// 		// 	result[istart + (3 * irow) - 1] = 9
// 		// 	irow++
// 		// }
// 		// _ = istart
		
// 	}

// 	_ = irow
// 	return result
// }
