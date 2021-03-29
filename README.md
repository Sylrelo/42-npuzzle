# 42 N-PUZZLE

The goal of this project is to solve the N-puzzle ("taquin" in French) game using the A* search algorithm or one of its variants.

### Contraints and general rules

- Manage various puzzle sizes (3, 4, 5, 17, etc...). The higher the program can go without dying a horrible, horrible death, the better
- The cost associated with each transition is always 1
- The user must be able to choose between at LEAST 3 (relevant) heuristic functions
- **The Manhattan-distance heuristic is mandatory**
- At the end of the search, the program has to provide the following values :
	- Complexity in time (total number of states ever selected in the "opened" set)
	- Complexity in size (Maximum number of states ever represented in memory at the same time during the search)
	- Number of moves required to transition from the initial state to the final state
	- The ordered sequence of states that make up the solution
- If the puzzle is unsolvable, the user must be informed and the program must exit properly


### Usage
```bash
go run ./src -f [FILE]
```
or
```bash
go run ./src <<< "$(python res_npuzzle-gen.py -s 3)"
```

### N-Puzzle file format
```
// Puzzle size, must be >= 3
3
// Followed by puzzle definition
1 2 3
0 8 7
6 5 4
```