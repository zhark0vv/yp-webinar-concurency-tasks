package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go sendValues(ch1, 1, 2, 3, 4, 5)

	ch2 := make(chan int)
	go sendValues(ch2, 6, 7, 8, 9, 10)

	ch3 := make(chan int)
	go sendValues(ch3, 11, 12, 13, 14, 15)

	out := mergeTwo(ch1, ch2)
	go printChan(out)

	out := mergeN(ch1, ch2, ch3)

	go printChan(out)

	time.Sleep(1 * time.Second)
}

func sendValues(ch chan int, values ...int) {
	for _, v := range values {
		ch <- v
	}
	close(ch)
}

func printChan(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func mergeTwo(ch1, ch2 chan int) chan int {
	// Твой код здесь...
}

func mergeN(chans ...chan int) chan int {
	// Твой код здесь...
}
