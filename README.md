# 42 N-PUZZLE

The goal of this project is to solve the N-puzzle ("taquin" in French) game using the A* search algorithm or one of its variants.


`∆ The project and the README are still WIP`

## 🟧 **`Contraints and general rules`**

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

## 🟧 **`Usage`**
```bash
go run ./src -f [FILE]
```
```bash
go run ./src <<< "$(python res_npuzzle-gen.py -s 3)"
```

`Puzzle file format`
```
// Puzzle size, must be >= 3
3
// Followed by puzzle definition
1 2 3
0 8 7
6 5 4
```

## 🟨 **`Informations`**
### &nbsp; &nbsp; ⬜️ **Implemented Algorithms**

&nbsp; &nbsp; &nbsp; &nbsp; ▫️ **A\***
>>> The A* algorithm is an algorithm for finding a path in a graph between an initial node and an end node. It uses a heuristic evaluation on each node to estimate the best path through it, and then visits the nodes in order of this heuristic evaluation, ignoring already visited nodes

&nbsp; &nbsp; &nbsp; &nbsp; ▫️ ~~**IDA\***~~
>>> The IDA* algorithm is basically the same than A*, except it concentrate on exploring the most promising node without ignoring already visited one

&nbsp; &nbsp; &nbsp; &nbsp; ▫️ ~~**Breadth-first search**~~

### &nbsp; &nbsp; ⬜️ **Heuristics**
&nbsp; &nbsp; &nbsp; &nbsp; ▫️ **Hamming Distance**
>>> The Hamming distance is the total number of misplaced tiles

&nbsp; &nbsp; &nbsp; &nbsp; ▫️ **Manhattan Distance**
>>> The distance between two points measured along axes at right angles

&nbsp; &nbsp; &nbsp; &nbsp; ▫️ **Linear Conflict**
>>> If two tiles are in the same row/column, and their goal positions are in the same row/colum, a linear conflict happens.
This heuristic is always combined with Manhattan Distance

&nbsp; &nbsp; &nbsp; &nbsp; ▫️ ~~**Gaschnig**~~
