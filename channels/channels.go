package channels

import (
	"fmt"
	"time"
)

func Producer(ch chan<- int) {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Produced %d\n", i)
		ch <- i
	}
	close(ch)
}

func Consumer(ch <- chan int) {
	for value := range ch {
		fmt.Printf("Consumed %d\n", value)
		time.Sleep(3 * time.Second)
		fmt.Println("Consumer: Ready for new value")
	}
}

func CallMe() {
	ch := make(chan int, 3)
	go Producer(ch)
	Consumer(ch)
}
