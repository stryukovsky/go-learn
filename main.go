package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(id int, group *sync.WaitGroup) {
	defer group.Done()
	fmt.Printf("Worker #%d sleep for %d seconds\n", id, id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker #%d finished\n", id)
}

func worker(ch chan string) {
	fmt.Println("Worker")
	ch <- "Hello from worker"
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go printNumbers(1, &wg)
	wg.Add(1)
	go printNumbers(2, &wg)
	wg.Wait()

	channel := make(chan int, 3)
	channel <- 1
	channel <- 2
	channel <- 3
	close(channel)

	for value := range channel {
		fmt.Println(value)
	}
}
