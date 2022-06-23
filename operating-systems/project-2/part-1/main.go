package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const size = 1000

func main() {
	start1 := time.Now()

	matrix1 := generateMatrix(size)
	matrix2 := generateMatrix(size)

	elapsed1 := time.Since(start1)
	fmt.Printf("time took for generation: %s\n", elapsed1)

	// -----------------------------------------------------------------

	start2 := time.Now()
	sumMatrix := multiply(matrix1, matrix2, size)

	elapsed2 := time.Since(start2)
	fmt.Printf("time took for multiplication: %s\n", elapsed2)

	_ = sumMatrix
}

func generateMatrix(size int) [][]int {
	matrix := make([][]int, size)

	//looping over each element of array and assigning it a random variable
	for row := 0; row < size; row++ {
		matrix[row] = make([]int, size)

		for column := 0; column < size; column++ {
			matrix[row][column] = rand.Intn(100)
		}
	}

	return matrix
}

func multiply(matrix1, matrix2 [][]int, size int) [][]int {
	// initialize result
	result := make([][]int, size)

	// WaitGroup are ways to synchronize green threads
	var wg sync.WaitGroup
	wg.Add(size * size)

	for row := 0; row < size; row++ {
		for column := 0; column < size; column++ {
			result[row] = make([]int, size)

			go func(row, column int) {
				result[row][column] = 0
				for iterator := 0; iterator < size; iterator++ {
					result[row][column] += matrix1[row][iterator] * matrix2[iterator][column]
				}
				wg.Done()
			}(row, column)
		}
	}

	wg.Wait()

	return result
}

// type channel struct {
// 	row, column, value int
// }
