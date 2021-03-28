package main

import "fmt"

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

func PrintBoard(board []int, size Size) {
	for i := 0; i < size.ncol; i++ {
		fmt.Print(board[0+i], " ")
	}
	fmt.Println("")
	for i := 0; i < size.ncol; i++ {
		fmt.Print(board[(size.ncol)+i], " ")
	}
	fmt.Println("")
	for i := 0; i < size.ncol; i++ {
		fmt.Print(board[(size.ncol*2)+i], " ")
	}
	fmt.Println("\n ")
}

func GenerateSnail(col int, row int) []int {
	result 		:= make([]int, col * row)
	//direction 	:= RIGHT
	irow		:= 0
	crow		:= 0
	dir			:= 1
	//icol		:= 0
	//istart		:= 0

	for i := col * row - 1; i >= 0; i-- {

		irow = 0
		for irow < (col * row) {

			crow += dir
			irow++

			dir *= -1
		}
		// result[i - 1] = 0
		// if direction == RIGHT && icol < col {
		// 	result[i - 1] = i
		// 	icol++
		// }
		// if icol == 3 && direction == RIGHT {
		// 	direction = DOWN
		// 	irow = 0
		// 	istart = row + icol
		// }
		// if direction == DOWN && irow <= 1 {
		// 	result[istart + (3 * irow) - 1] = 9
		// 	irow++
		// }
		// _ = istart
		
	}

	_ = irow
	return result
}
