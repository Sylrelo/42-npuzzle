package main

import "fmt"

func FindIndex(haystack [9]int, needle int) int {
	for i, n := range haystack {
		if (needle == n) {
			return i
		}
	}
	return -1
}

func Same(a [9]int, b [9]int) bool {
	for i, n := range a {
		if (b[i] != n) {
			return false
		}
	}
	return true
}

func PrintBoard(board [9]int) {
	for i := 0; i < 3; i++ {
		fmt.Print(board[0 + i], " ")
	}
	fmt.Println("")
	for i := 0; i < 3; i++ {
		fmt.Print(board[3 + i], " ")
	}
	fmt.Println("")
	for i := 0; i < 3; i++ {
		fmt.Print(board[6 + i], " ")
	}
	fmt.Println("")
}