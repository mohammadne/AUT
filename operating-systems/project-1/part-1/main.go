package main

import (
	"fmt"
	"syscall"
)

func main() {
	foo := 4
	bar := 10

	id, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
		fmt.Println(err.Error())
	} else if id == 0 {
		foo++
		fmt.Println("In child:", id, foo, bar)
	} else {
		bar++
		fmt.Println("In parent:", id, foo, bar)
	}
}
