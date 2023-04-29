package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cell struct {
	visited bool
	walls   [4]bool // top, right, bottom, left
}

type maze struct {
	rows, cols int
	cells      [][]cell
}

func (m *maze) init() {
	for i := 0; i < m.rows; i++ {
		row := make([]cell, m.cols)
		for j := 0; j < m.cols; j++ {
			row[j] = cell{visited: false, walls: [4]bool{true, true, true, true}}
		}
		m.cells = append(m.cells, row)
	}
}

func (m *maze) dfs(row, col int) {
	m.cells[row][col].visited = true

	// Shuffle the neighbors
	neighbors := []struct{ r, c int }{
		{row - 1, col}, // Top
		{row, col + 1}, // Right
		{row + 1, col}, // Bottom
		{row, col - 1}, // Left
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(neighbors), func(i, j int) {
		neighbors[i], neighbors[j] = neighbors[j], neighbors[i]
	})

	// Visit each neighbor
	for _, n := range neighbors {
		if n.r >= 0 && n.r < m.rows && n.c >= 0 && n.c < m.cols {
			if !m.cells[n.r][n.c].visited {
				// Remove the wall between the current cell and the neighbor
				if n.r < row {
					m.cells[row][col].walls[0] = false // Top
					m.cells[n.r][n.c].walls[2] = false // Bottom
				} else if n.c > col {
					m.cells[row][col].walls[1] = false // Right
					m.cells[n.r][n.c].walls[3] = false // Left
				} else if n.r > row {
					m.cells[row][col].walls[2] = false // Bottom
					m.cells[n.r][n.c].walls[0] = false // Top
				} else {
					m.cells[row][col].walls[3] = false // Left
					m.cells[n.r][n.c].walls[1] = false // Right
				}
				m.dfs(n.r, n.c)
			}
		}
	}
}

func main() {
	var rows, cols int

	fmt.Println("Informe quantas linhas deseja")
	fmt.Scan(&rows)
	fmt.Println("Informe quantas colunas deseja")
	fmt.Scan(&cols)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	m := maze{rows: rows, cols: cols}
	m.init()
	m.dfs(0, 0)

	// Print the maze
	fmt.Print(" ")
	for j := 0; j < cols; j++ {
		fmt.Printf("_ ")
	}
	fmt.Println()
	for i := 0; i < rows; i++ {
		fmt.Print("|")
		for j := 0; j < cols; j++ {
			if m.cells[i][j].walls[2] {
				fmt.Print("_")
			} else {
				fmt.Print(" ")
			}
			if m.cells[i][j].walls[1] {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
