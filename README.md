## 42 N-PUZZLE

Objective : Resolve a n-puzzle

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