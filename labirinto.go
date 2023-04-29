package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
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

func printMaze(m maze) {
	// Print the maze
	fmt.Print(" ")
	for j := 0; j < m.cols; j++ {
		fmt.Printf("_ ")
	}
	fmt.Println()
	for i := 0; i < m.rows; i++ {
		fmt.Print("|")
		for j := 0; j < m.cols; j++ {
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

func createImage(m maze, dir string) {
	rows, cols := m.rows, m.cols

	// Create a new RGBA image with a white background
	img := image.NewRGBA(image.Rect(0, 0, cols*10, rows*10))
	bg := color.RGBA{255, 255, 255, 255}
	for i := 0; i < img.Bounds().Dx(); i++ {
		for j := 0; j < img.Bounds().Dy(); j++ {
			img.SetRGBA(i, j, bg)
		}
	}

	// Draw the maze on the image
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if m.cells[i][j].walls[2] {
				for k := 0; k < 10; k++ {
					img.SetRGBA(j*10+k, (i+1)*10, color.RGBA{0, 0, 0, 255})
				}
			}
			if m.cells[i][j].walls[1] {
				for k := 0; k < 10; k++ {
					img.SetRGBA((j+1)*10, i*10+k, color.RGBA{0, 0, 0, 255})
				}
			}
		}
	}

	// Save the image to a file
	f, _ := os.Create(dir + ".png")
	defer f.Close()
	png.Encode(f, img)
	fmt.Println("Imagem salva com sucesso")
}

func createMaze(rows, cols int) maze {
	m := maze{rows: rows, cols: cols}
	m.init()
	m.dfs(0, 0)

	return m
}

func getInput() (int, int, string) {
	var rows, cols int
	var dir string

	fmt.Println("Informe quantas linhas deseja")
	fmt.Scan(&rows)
	fmt.Println("Informe quantas colunas deseja")
	fmt.Scan(&cols)

	fmt.Println("Informe o nome da imagem")
	fmt.Scan(&dir)

	return rows, cols, dir
}

func main() {
	rows, cols, dir := getInput()

	fmt.Println()
	fmt.Println()

	m := createMaze(rows, cols)
	printMaze(m)

	fmt.Println()
	fmt.Println()

	createImage(m, dir)
}