package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"
)

const (
	SQ = 3
	N  = 9
)

var (
	iterCnt int
)

type intmat [N][N]int

func solveSudoku(m intmat, row, col int) bool {

	if row == N-1 && col == N {
		return true
	}

	if col == N {
		row++
		col = 0
	}

	if m[row][col] != 0 {
		return solveSudoku(m, row, col+1)
	}

	for num := 1; num <= N; num++ {
		if isSafe(m, row, col, num) {
			m[row][col] = num

			if solveSudoku(m, row, col+1) {
				return true
			}
		}
		m[row][col] = 0
		iterCnt++
	}
	
	return false
}

func isSafe(m intmat, row, col, num int) bool {

	for a := 0; a < N; a++ {
		if m[row][a] == num {
			return false
		}
	}

	for a := 0; a < N; a++ {
		if m[a][col] == num {
			return false
		}
	}

	startRow := row - row%3
	startCol := col - col%3

	for i := 0; i < N/3; i++ {
		for j := 0; j < N/3; j++ {
			if m[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

func toGrid(s string) intmat {
	var (
		index  int
		err    error
		substr string
		m      intmat
	)
	m = intmat{}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			substr = s[index : index+1]
			fmt.Print(substr)
			if strings.Compare(substr, ".") == 0 {
				m[i][j] = 0
			} else {
				m[i][j], err = strconv.Atoi(substr)
			}

			if err != nil {
				log.Fatalf("Error in reading char at %d,%d.\n", i, j)
			}
			index++
		}
	}
	fmt.Println()
	return m
}

func printSudoku(m intmat) {
	var sqi, sqj int

	for i := 0; i < N; i++ {
		sqi = (i / SQ) % 2
		for j := 0; j < N; j++ {
			sqj = (j / SQ) % 2
			if (sqi == 0 && sqj == 1) || (sqi == 1 && sqj == 0) {
				color.LightBlue.Printf("%d ", m[i][j])
			} else {
				color.LightGreen.Printf("%d ", m[i][j])
			}
		}
		fmt.Println()
	}
}

func main() {

	//input := "3.65.84..52........87....31..3.1..8.9..863..5.5..9.6..13....25........74..52.63.."
	input := "...15....91..764..5.6.4.3........69.6..5.4..7.71........7.3.9.6..386..15....95..."

	if len(input) != N*N {
		log.Fatalf("Expected %d but got %d.\n", N*N, len(input))
	}

	mat := toGrid(input)
	printSudoku(mat)

	start := time.Now()
	solveSudoku(mat, 0, 0)
	elapsed := time.Since(start).Seconds()

	fmt.Println("\nFinal solution:")
	printSudoku(mat)
	log.Printf("Iterations: %d. Sudoku took %v sec\n", iterCnt, elapsed)
}
