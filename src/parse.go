package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseContent(content string, heuristicName string, verbose bool) (int, []int, HeuriFunc, bool) {
	var size int
	var result []int

	heurMap := map[string]HeuriFunc{
		"MANHATTAN": ManhattanDistance,
		"MISPLACED": HammingDistance,
		"LINEARCONFLICT": LinearConflict,
		"EUCLIDIAN": EuclideanDistance,
	}

	size = -1
	content_splitted := strings.Split(content, "\n")

	for line_number, line := range content_splitted {
		if len(line) < 1 || line[0] == '#' {
			continue
		}
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
					result = append(result, atoi)
				}
			}
		}
	}
	return size, result, heurMap[heuristicName], verbose
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

func Parse() (int, []int, HeuriFunc, bool) {
	// args := os.Args[1:]

	filepath := flag.String("file", "", "Path to the puzzle file")
	verbose := flag.Bool("v", false, "Enables board overview through resolution")
	// greedy 		:= flag.Bool("g", false, "This flag enable the greedy-search")
	// uniform 		:= flag.Bool("u", false, "This flag enable the uniform-cost search")
	// algo 		:= flag.String("a", "", "Run the specified algorithm only [ ASTAR, IDASTAR, BFS, DSF ]")
	heuristic := flag.String("h", "MANHATTAN", "Run the specified heuristic. Only [ MANHATTAN, MISPLACED, LINEARCONFLICT ]")
	flag.Parse()
	data, err := os.ReadFile(*filepath)
	if err != nil {
		fmt.Println("Error reading file", *filepath)
		os.Exit(1)
	}
	if !CheckHeuristicName(*heuristic) {
		fmt.Println("Error in heuristic choice : must be one of [ MANHATTAN, MISPLACED, LINEARCONFLICT ] and not " + *heuristic)
		os.Exit(1)
	}
	return ParseContent(string(data), *heuristic, *verbose)

	// if len(args) > 0 {
	// 	if args[0] == "-f" {
	// 		if len(args) > 1 {
	// 			data, err := os.ReadFile(args[1])
	// 			if err != nil {
	// 				fmt.Println("Error reading file", args[1])
	// 				os.Exit(1)
	// 			}
	// 			return ParseContent(string(data))
	// 		} else {
	// 			fmt.Println("You must specify a file")
	// 			os.Exit(1)
	// 		}
	// 	}
	// } else {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	content, _ := reader.ReadString(0)
	// 	return ParseContent(content)
	// }

	fmt.Println("Parsing failed for unknown reason.")
	os.Exit(1)
	return -1, nil, nil, false
}
