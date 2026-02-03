package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Board []uint8

type DFSNode struct {
	board Board
	depth int
}

func loadBoard(input string) (Board, int) {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data struct {
		TableSize    string  `json:"tableSize"`
		InitialBoard [][]int `json:"initialBoard"`
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}

	bs, err := strconv.Atoi(data.TableSize)
	if err != nil {
		panic(fmt.Sprintf("Invalid tableSize: %v", err))
	}

	board := make(Board, bs*bs)
	for r := 0; r < bs; r++ {
		for c := 0; c < bs; c++ {
			board[r*bs+c] = uint8(data.InitialBoard[r][c])
		}
	}
	return board, bs
}

func loadSampleBoard(bs int) Board {
	switch bs {
	case 5:
		return Board{
			9, 24, 3, 5, 17,
			6, 0, 13, 19, 10,
			11, 21, 22, 1, 20,
			16, 4, 14, 12, 15,
			8, 18, 23, 2, 7,
		}
	case 4:
		return Board{
			1, 2, 7, 4,
			5, 15, 0, 8,
			9, 10, 3, 11,
			13, 14, 6, 12,
		}
	case 3:
		return Board{
			1, 2, 3,
			8, 0, 5,
			7, 6, 4,
		}
	default:
		panic("Unsupported board size")
	}
}

func goalBoard(bs int) Board {
	size := bs * bs
	board := make(Board, size)
	for i := 0; i < size-1; i++ {
		board[i] = uint8(i + 1)
	}
	board[size-1] = 0
	return board
}

func swap(b Board, i, j int) Board {
	nb := make(Board, len(b))
	copy(nb, b)
	nb[i], nb[j] = nb[j], nb[i]
	return nb
}

func boardKey(b Board) string {
	return string(b)
}

func findBlank(b Board) int {
	for i, v := range b {
		if v == 0 {
			return i
		}
	}
	return -1
}

func successors(b Board, bs int) []Board {
	blank := findBlank(b)
	row := blank / bs
	col := blank % bs

	var next []Board

	if row > 0 {
		next = append(next, swap(b, blank, blank-bs))
	}
	if row < bs-1 {
		next = append(next, swap(b, blank, blank+bs))
	}
	if col > 0 {
		next = append(next, swap(b, blank, blank-1))
	}
	if col < bs-1 {
		next = append(next, swap(b, blank, blank+1))
	}

	return next
}

func bfs(start Board, goal Board, bs int) ([]Board, bool) {
	queue := []Board{start}
	visited := make(map[string]bool)
	parent := make(map[string]Board)
	nodesExplored := 0
	visited[boardKey(start)] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		nodesExplored++

		if boardKey(current) == boardKey(goal) {
			var path []Board
			for boardKey(current) != boardKey(start) {
				path = append([]Board{current}, path...)
				current = parent[boardKey(current)]
			}
			path = append([]Board{start}, path...)
			return path, true
		}

		for _, next := range successors(current, bs) {
			key := boardKey(next)
			if !visited[key] {
				visited[key] = true
				parent[key] = current
				queue = append(queue, next)
			}
			if nodesExplored%1000000 == 0 {
				fmt.Printf("Nodes explored: %d\r", nodesExplored)
			}
		}
	}

	fmt.Printf("Total nodes explored: %d\n", nodesExplored)
	return nil, false
}

func dfs(start Board, goal Board, bs int, maxDepth int) ([]Board, bool) {
	stack := []DFSNode{{board: start, depth: 0}}
	visited := make(map[string]bool)
	parent := make(map[string]Board)

	visited[boardKey(start)] = true

	nodesExplored := 0

	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		current := n.board
		depth := n.depth
		nodesExplored++

		if boardKey(current) == boardKey(goal) {
			// reconstruct path
			var path []Board
			for boardKey(current) != boardKey(start) {
				path = append([]Board{current}, path...)
				current = parent[boardKey(current)]
			}
			path = append([]Board{start}, path...)
			fmt.Printf("DFS total nodes explored: %d\n", nodesExplored)
			return path, true
		}

		if depth >= maxDepth {
			continue
		}

		for _, next := range successors(current, bs) {
			key := boardKey(next)
			if !visited[key] {
				visited[key] = true
				parent[key] = current
				stack = append(stack, DFSNode{
					board: next,
					depth: depth + 1,
				})
				if nodesExplored%1000000 == 0 {
					fmt.Printf("Nodes explored: %d\r", nodesExplored)
				}
			}
		}
	}

	fmt.Printf("DFS total nodes explored: %d\n", nodesExplored)
	return nil, false
}

func main() {
	boardSize := flag.Int("s", 3, "Size of the board: valid options are 3, 4, or 5")
	searchMethod := flag.String("m", "bfs", "Search method: vaild options are 'bfs' or 'dfs'")
	dfsDepth := flag.Int("d", 30, "Max depth for DFS")
	uboard := flag.String("i", "", "Path to input board file (JSON)")
	var board Board
	flag.Parse()
	startTime := time.Now()
	if *uboard != "" {
		var bs int
		board, bs = loadBoard(*uboard)
		boardSize = &bs
		println("Using user-provided board from", *uboard)
	} else {
		fmt.Printf("Using sample %d x %d board.\n", *boardSize, *boardSize)
		board = loadSampleBoard(*boardSize)
	}
	goal := goalBoard(*boardSize)
	switch *searchMethod {
	case "bfs":
		path, found := bfs(board, goal, *boardSize)
		if found {
			fmt.Printf("BFS Solution Found! Path length: %d\n", len(path))
			fmt.Printf("Time taken: %s\n", time.Since(startTime))
			fmt.Printf("Time Started: %s\nTime Ended: %s\n", startTime.Format(time.RFC3339), time.Now().Format(time.RFC3339))
			for i, step := range path {
				fmt.Printf("Step %d:\n", i)
				for r := 0; r < *boardSize; r++ {
					for c := 0; c < *boardSize; c++ {
						fmt.Printf("%d ", step[r**boardSize+c])
					}
					fmt.Println()
				}
				fmt.Println()
			}
		}
	case "dfs":
		path, found := dfs(board, goal, *boardSize, *dfsDepth)

		if found {
			fmt.Printf("DFS Solution Found! Path length: %d\n", len(path))
			fmt.Printf("Time taken: %s\n", time.Since(startTime))
			fmt.Printf("Time Started: %s\nTime Ended: %s\n", startTime.Format(time.RFC3339), time.Now().Format(time.RFC3339))
			for i, step := range path {
				fmt.Printf("Step %d:\n", i)
				for r := 0; r < *boardSize; r++ {
					for c := 0; c < *boardSize; c++ {
						fmt.Printf("%d ", step[r**boardSize+c])
					}
					fmt.Println()
				}
				fmt.Println()
			}
		} else {
			fmt.Println("DFS: No solution within depth limit")
		}
	default:
		println("Invalid search method. Use 'bfs' or 'dfs'.")
	}
}
