package main

import "fmt"

func main() {
	// Initial matrix
	matrix := [][]bool{
		{true, false, true},
		{false, true, false},
		{false, false, true},
	}
	printGame(matrix)

	fmt.Println()

	// Create a hard copy of the matrix for new state
	var newMatrix = make([][]bool, len(matrix))
	for i := range matrix {
		newMatrix[i] = make([]bool, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
	}

	for rowIdx, row := range matrix {
		for colIdx := range row {
			liveNeighbourCount := analyzeNeighbours(rowIdx, colIdx, matrix)

			if matrix[rowIdx][colIdx] {
				newMatrix[rowIdx][colIdx] = liveNeighbourCount == 2 || liveNeighbourCount == 3
			} else {
				newMatrix[rowIdx][colIdx] = liveNeighbourCount == 3
			}
		}
	}

	// Print the new matrix after applying the rules
	fmt.Println("New Matrix (After One Generation):")
	printGame(newMatrix)
}

// Function to print the matrix
func printGame(matrix [][]bool) {
	for _, row := range matrix {
		for _, col := range row {
			if col {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}

func analyzeNeighbours(rowIdx, colIdx int, matrix [][]bool) int {
	var liveCount int
	rows := len(matrix)
	cols := len(matrix[0])

	for i := rowIdx - 1; i <= rowIdx+1; i++ {
		for j := colIdx - 1; j <= colIdx+1; j++ {
			if i >= 0 && i < rows && j >= 0 && j < cols {
				if !(i == rowIdx && j == colIdx) {
					if matrix[i][j] {
						liveCount++
					}
				}
			}
		}
	}

	return liveCount
}
