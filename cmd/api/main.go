package main

import "fmt"

type Board [25]uint8

type Node struct {
	board Board
	path  []Board
}

func loadBoard() [25]uint8 {
	board := Board{
		6, 4, 7, 2, 5,
		8, 0, 3, 1, 9,
		10, 11, 12, 13, 14,
		15, 16, 17, 18, 19,
		20, 21, 22, 23, 24,
	}
	return board
}

func goalBoard() [25]uint8 {
	board := Board{
		1, 2, 3, 4, 5,
		6, 7, 8, 9, 10,
		11, 12, 13, 14, 15,
		16, 17, 18, 19, 20,
		21, 22, 23, 24, 0,
	}
	return board
}

func swap(b Board, i, j int) Board {
	b[i], b[j] = b[j], b[i]
	return b
}

func findBlank(b Board) int {
	for i, v := range b {
		if v == 0 {
			return i
		}
	}
	return -1
}

func successors(b Board) []Board {
	blank := findBlank(b)
	row := blank / 5
	col := blank % 5

	var next []Board

	if row > 0 {
		next = append(next, swap(b, blank, blank-3))
	}
	if row < 2 {
		next = append(next, swap(b, blank, blank+3))
	}
	if col > 0 {
		next = append(next, swap(b, blank, blank-1))
	}
	if col < 2 {
		next = append(next, swap(b, blank, blank+1))
	}

	return next
}

func bfs(start Board, goal Board) ([]Board, bool) {
	queue := []Node{
		{
			board: start,
			path:  []Board{start},
		},
	}

	visited := make(map[Board]bool)
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.board == goal {
			return current.path, true
		}

		for _, next := range successors(current.board) {
			if !visited[next] {
				visited[next] = true

				newPath := make([]Board, len(current.path))
				copy(newPath, current.path)
				newPath = append(newPath, next)

				queue = append(queue, Node{
					board: next,
					path:  newPath,
				})
			}
		}
	}

	return nil, false
}

func main() {
	board := loadBoard()
	goal := goalBoard()

	path, found := bfs(board, goal)
	if found {
		fmt.Printf("Solution Found! Path length: %d\n", len(path))
		for i, step := range path {
			fmt.Printf("Step %d:\n", i)
			for r := 0; r < 3; r++ {
				fmt.Printf("%d %d %d\n", step[r*3], step[r*3+1], step[r*3+2])
			}
			fmt.Println()
		}
	} else {
		println("No solution found.")
	}

}
