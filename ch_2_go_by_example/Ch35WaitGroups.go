package main

import (
	"fmt"
	"sync"
	"time"
)

func worker2(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Wroker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	//Note that a WaitGroup must be passed to functions by pointer
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker2(i, &wg)
	}

	wg.Wait()
}
