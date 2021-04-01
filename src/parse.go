package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ParseContent(content string) (int, []int) {
	var size int
	var result []int

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
				fmt.Println("Error puzzle size, line", line_number + 1)
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
	return size, result
}

func Parse() (int, []int) {
	args := os.Args[1:]


	// filepath 	:= flag.String("file", "", "The input file for the puzzle")
	// greedy 		:= flag.Bool("g", false, "This flag enable the greedy-search")
	// uniform 		:= flag.Bool("u", false, "This flag enable the uniform-cost search")
	// algo 		:= flag.String("a", "", "Run the specified algorithm only [ ASTAR, IDASTAR, BFS, DSF ]")
	// heuristic 	:= flag.String("h", "", "Run the specified heuristic only [ MANHATTAN, MISPLACED, LINEARCONFLICT ]")
	// flag.Parse()

	if len(args) > 0 {
		if args[0] == "-f" {
			if len(args) > 1 {
				data, err := ioutil.ReadFile(args[1])
				if err != nil {
					fmt.Println("Error reading file", args[1])
					os.Exit(1)
				}
				return ParseContent(string(data))
			} else {
				fmt.Println("You must specify a file")
				os.Exit(1)
			}
		}
	} else {
		reader 		:= bufio.NewReader(os.Stdin)
		content, _ 	:= reader.ReadString(0)
		return ParseContent(content)
	}
	fmt.Println("Parsing failed for unknown reason.")
	os.Exit(1)
	return -1, nil
}