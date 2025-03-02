package channels

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex sync.Mutex
)

func Increment(wg *sync.WaitGroup) {
	mutex.Lock()
	defer mutex.Unlock()
	defer wg.Done()

	counter++
	fmt.Println("Counter: ", counter)
}

func CallMutexDemo() {
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go Increment(&wg)
	}
	wg.Wait()
}
