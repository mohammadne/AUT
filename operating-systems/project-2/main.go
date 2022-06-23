package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

type (
	car       bool
	direction bool
)

var (
	normal car = true
	heavy  car = true

	north direction = true
	south direction = false
)

const (
	queueSize  = 10000
	bridgeSize = 10
)

func main() {
	mutex := sync.Mutex{}

	cars := map[direction][]car{
		north: generateCars(queueSize),
		south: generateCars(queueSize),
	}

	semaphor := sync.WaitGroup{}
	semaphor.Add(2)

	for _, dir := range []direction{north, south} {
		go passCars(cars[dir], dir, &mutex, &semaphor)
	}

	semaphor.Wait()

	fmt.Println("all cars have been passed")
}

// generateCars will generate a queue of cars for the given size
//
// 0 indicates normal car (0.8 chance of generation)
//
// 1 indicate a heavy car (0.2 chance of generation)
func generateCars(size int) []car {
	cars := make([]car, size)

	for index := 0; index < size; index++ {
		cars[index] = rand.Intn(10) > 7
	}

	return cars
}

func passCars(cars []car, dir direction, mutex *sync.Mutex, semaphor *sync.WaitGroup) {
	pointer := 0

	for {
		if pointer >= queueSize-1 {
			semaphor.Done()
			break
		}

		mutex.Lock()

		to := min(pointer+bridgeSize, queueSize)

		for index := pointer; index < to; index++ {
			if cars[index] == heavy {
				if index == pointer {
					to = index + 1
				} else {
					to = index
				}
				break
			}
		}

		var stringDir string
		if dir == north {
			stringDir = "north"
		} else {
			stringDir = "south"
		}

		if cars[pointer] == heavy {
			fmt.Printf("Heavy car %d passed in %s direction\n", pointer, stringDir)
		} else if to-pointer <= 1 {
			fmt.Printf("Normal car %d passed in %s direction\n", pointer, stringDir)
		} else {
			fmt.Printf("Normal cars %d-%d passed in %s direction\n", pointer, to-1, stringDir)
		}

		pointer = to
		mutex.Unlock()
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
