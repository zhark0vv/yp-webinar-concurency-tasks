package main

import (
	"fmt"
	"sync"
)

type Job struct {
	Value int
	// Priority int - опционально
}

// worker обрабатывает задачи из канала jobs и отправляет результаты в канал results
func worker(wg *sync.WaitGroup, f func(n int) int, jobs <-chan Job, results chan<- int) {
	// Твой код здесь...
}

func main() {
	const (
		numJobs    = 10
		numWorkers = 3
	)

	// Создаем каналы
	jobs := make(chan Job, numJobs)
	results := make(chan int, numJobs)
	wg := sync.WaitGroup{}

	mul := func(n int) int {
		return n * 2
	}

	// Твой код здесь...

	results := worker(&wg, mul, jobs, results)

	// Вывод результатов
	fmt.Println("Results:")
	for result := range results {
		fmt.Println(result)
	}
}
