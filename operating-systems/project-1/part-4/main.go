package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	isMultiThread = false
	size          = 10000
)

func main() {
	matrix1 := generateMatrix(size)
	matrix2 := generateMatrix(size)

	var sumMatrix [][]int

	start := time.Now()

	if isMultiThread {
		sumMatrix = multiThread(matrix1, matrix2, size)
	} else {
		sumMatrix = singleThread(matrix1, matrix2, size)
	}

	// print summation calculation time
	elapsed := time.Since(start)
	fmt.Printf("time took: %s\n", elapsed)

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

func singleThread(matrix1, matrix2 [][]int, size int) [][]int {
	// initialize sumMatrix
	sumMatrix := make([][]int, size)

	for row := 0; row < size; row++ {
		sumMatrix[row] = sumVectors(matrix1[row], matrix2[row])
	}

	return sumMatrix
}

func multiThread(matrix1, matrix2 [][]int, size int) [][]int {
	// initialize sumMatrix
	sumMatrix := make([][]int, size)

	// WaitGroup are ways to synchronize green threads
	var wg sync.WaitGroup
	wg.Add(size)

	for row := 0; row < size; row++ {
		go func(row int) {
			defer wg.Done()
			sumMatrix[row] = sumVectors(matrix1[row], matrix2[row])
		}(row)
	}

	wg.Wait()

	return sumMatrix
}

func sumVectors(vec1, vec2 []int) []int {
	length := len(vec1)
	sum := make([]int, 0, length)

	for index := 0; index < length; index++ {
		sum = append(sum, vec1[index]+vec2[index])
	}

	return sum
}
