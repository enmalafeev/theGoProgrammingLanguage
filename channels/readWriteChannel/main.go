package main

import (
	"fmt"
	"time"
)

func main() {
	reader(double(writer()))
}

func writer() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := range 10 {
			ch <- i + 1
		}
	}()

	return ch
}

func double(ch <-chan int) <-chan int {
	outCh := make(chan int)

	go func() {
		defer close(outCh)
		for v := range ch {
			time.Sleep(500 * time.Millisecond)
			outCh <- v * 2
		}
	}()

	return outCh
}

func reader(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
