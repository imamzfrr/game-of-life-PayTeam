package main

import (
	"fmt"
	"time"
)

const (
	width  = 3 // Ukuran grid (bisa diubah sesuai kebutuhan)
	height = 3
)

func newGrid() [][]bool {
	return [][]bool{
		{true, false, true},
		{false, true, false},
		{true, false, true},
	}
}

func printGrid(grid [][]bool) {
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func nextGeneration(grid [][]bool) [][]bool {
	newGrid := newGrid()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			liveNeighbors := countLiveNeighbors(grid, x, y)

			if grid[y][x] {
				newGrid[y][x] = liveNeighbors == 2 || liveNeighbors == 3
			} else {
				newGrid[y][x] = liveNeighbors == 3
			}
		}
	}

	return newGrid
}

func countLiveNeighbors(grid [][]bool, x, y int) int {
	neighbors := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	count := 0

	for _, n := range neighbors {
		nx, ny := x+n[0], y+n[1]
		if nx >= 0 && nx < width && ny >= 0 && ny < height && grid[ny][nx] {
			count++
		}
	}

	return count
}

func main() {
	grid := newGrid()

	for i := 0; i < 3; i++ { // Simulasi 10 generasi
		printGrid(grid)
		grid = nextGeneration(grid)
		time.Sleep(time.Second)
	}
}
