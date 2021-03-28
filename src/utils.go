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

func Same(a []int, b []int) bool {
	for i, n := range a {
		if b[i] != n {
			return false
		}
	}
	return true
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
