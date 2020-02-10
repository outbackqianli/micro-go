package main

import (
	"fmt"
)

var (
	input chan int
	done  chan bool
)

func addTo(i int) {
	input <- i
}
func getTo() {
	for {
		if a, ok := <-input; !ok {
			done <- true
			break
		} else {
			fmt.Println("a is :", a)
		}
	}
}

func main() {
	input = make(chan int)
	output = make(chan int)
	prints = make([]int, 0)
	done = make(chan bool)
	//var wg sync.WaitGroup

	go func() {
		for i := 1; i < 10; i++ {
			addTo(i)
		}
		close(input)
	}()
	go getTo()
	<-done
}
