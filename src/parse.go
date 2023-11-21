package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkNums(board []int8, size int) {
	boardCopy := make([]int, len(board))

	for i := 0; i < size*size; i++ {
		boardCopy[i] = int(board[i])
	}
	sort.Ints((boardCopy))
	for i := 0; i < size*size; i++ {
		if boardCopy[i] != i {
			fmt.Println("Error puzzle : Wrong numbers")
			os.Exit(0)
		}
	}
}

func ParseContent(content string, heuristicName string, verbose bool) (int, []int8, string, bool) {
	var size int
	var result []int8

	size = -1
	content_splitted := strings.Split(content, "\n")

	for line_number, line := range content_splitted {
		if len(line) < 1 || line[0] == '#' {
			continue
		}
		line = strings.Split(line, "#")[0]
		if size == -1 {
			atoi, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Error puzzle size, line", line_number+1)
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
						fmt.Println("Error puzzle, line", line_number+1)
						os.Exit(0)
					}
					result = append(result, int8(atoi))
				}
			}
		}
	}
	if len(result) != size*size {
		fmt.Println("Error puzzle: wrong size")
		os.Exit(0)
	}
	checkNums(result, size)
	return size, result, heuristicName, verbose
}

func CheckHeuristicName(heuristicName string) bool {
	possibleHeurNames := [4]string{"MANHATTAN", "MISPLACED", "LINEARCONFLICT", "EUCLIDIAN"}

	for _, possibleName := range possibleHeurNames {
		if heuristicName == possibleName {
			return true
		}
	}
	return false
}

func Parse() (int, []int8, string, bool) {

	filepath := flag.String("file", "", "Path to the puzzle file")
	verbose := flag.Bool("v", false, "Enables board overview through resolution")
	// greedy 		:= flag.Bool("g", false, "This flag enable the greedy-search")
	// uniform 		:= flag.Bool("u", false, "This flag enable the uniform-cost search")
	heuristic := flag.String("h", "MANHATTAN", "Run the specified heuristic. Only [ MANHATTAN, MISPLACED, LINEARCONFLICT EUCLIDIAN ]")
	flag.Parse()
	data, err := os.ReadFile(*filepath)
	if err != nil {
		fmt.Println("Error reading file", *filepath)
		os.Exit(1)
	}
	if !CheckHeuristicName(*heuristic) {
		fmt.Println("Error in heuristic choice : must be one of [ MANHATTAN, MISPLACED, LINEARCONFLICT EUCLIDIAN ] and not " + *heuristic)
		os.Exit(1)
	}
	return ParseContent(string(data), *heuristic, *verbose)
}
