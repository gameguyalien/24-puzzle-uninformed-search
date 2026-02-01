package main

import (
	"flag"
	"fmt"
	"time"
)

type Board []uint8

func loadBoard(bs int) Board {
	if bs == 5 {
		return Board{
			9, 24, 3, 5, 17,
			6, 0, 13, 19, 10,
			11, 21, 22, 1, 20,
			16, 4, 14, 12, 15,
			8, 18, 23, 2, 7,
		}
	} else if bs == 4 {
		return Board{
			1, 2, 7, 4,
			5, 15, 0, 8,
			9, 10, 3, 11,
			13, 14, 6, 12,
		}
	} else if bs == 3 {
		return Board{
			1, 2, 3,
			4, 0, 5,
			7, 8, 6,
		}
	} else {
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

func main() {
	boardSize := flag.Int("s", 3, "Size of the board")
	flag.Parse()
	startTime := time.Now()
	board := loadBoard(*boardSize)
	goal := goalBoard(*boardSize)
	path, found := bfs(board, goal, *boardSize)
	if found {
		fmt.Printf("Solution Found! Path length: %d\n", len(path))
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
		println("No solution found.")
	}

}
