package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const size = 1000

type channel struct {
	row, column, value int
}

func main() {
	start1 := time.Now()

	matrix1 := generateMatrix(size)
	matrix2 := generateMatrix(size)

	elapsed1 := time.Since(start1)
	fmt.Printf("time took for generation: %s\n", elapsed1)

	// -----------------------------------------------------------------

	start2 := time.Now()

	ch := make(chan channel, size)
	go multiply(matrix1, matrix2, size, ch)

	// initialize result matrix
	result := make([][]int, size)
	for row := 0; row < size; row++ {
		result[row] = make([]int, size)
	}

	// do multiplication
	for index := 0; index < size*size; index++ {
		message := <-ch
		result[message.row][message.column] = message.value
	}

	elapsed2 := time.Since(start2)
	fmt.Printf("time took for multiplication: %s\n", elapsed2)
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

func multiply(matrix1, matrix2 [][]int, size int, ch chan channel) {
	var wg sync.WaitGroup
	wg.Add(size * size)

	for row := 0; row < size; row++ {
		for column := 0; column < size; column++ {
			go func(row, column int) {
				value := 0
				for iterator := 0; iterator < size; iterator++ {
					value += matrix1[row][iterator] * matrix2[iterator][column]
				}
				ch <- channel{row: row, column: column, value: value}
				wg.Done()
			}(row, column)
		}
	}

	wg.Wait()
}
