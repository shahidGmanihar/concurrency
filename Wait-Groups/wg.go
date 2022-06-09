package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	now := time.Now()
	go func() {
		A(&wg)
		B(&wg)

	}()
	wg.Wait()

	fmt.Println("Elapsed", time.Since(now))

	fmt.Println("Main function Done")

}

func A(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(250 * time.Millisecond)
	fmt.Println("I am fn A")
}

func B(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("I am fn B")

}
