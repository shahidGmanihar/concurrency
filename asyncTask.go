package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan string)
	now := time.Now()
	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)
	<-done
	<-done
	<-done
	<-done

	//time.Sleep(time.Second)
	fmt.Println("Elapsed", time.Since(now), done)

}

func task1(done chan string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task1")
	done <- "Done Here"

}

func task2(done chan string) {

	time.Sleep(200 * time.Millisecond)
	fmt.Println("task2")
	done <- "Done Here"

}

func task3(done chan string) {
	//time.Sleep(100*time.Millisecond)
	fmt.Println("task3")
	done <- "Done Here"

}

func task4(done chan string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task4")
	done <- "Done Here"

}
