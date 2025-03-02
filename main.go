package main

import (
	"fmt"
	"sync"
	"time"
)

import (
	"github.com/stryukovsky/go-learn/channels"
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
	channels.CallMe()	
}
