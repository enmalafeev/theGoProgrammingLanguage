package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := range 100000 {
			ch <- i + 1
		}
	}()

	go func() {
		for v := range ch {
			fmt.Println("v = ", v, "worker1")
		}
	}()

	go func() {
		for v := range ch {
			fmt.Println("v = ", v, "worker2")
		}
	}()

	for v := range ch {
		fmt.Println("v = ", v, "worker3")
	}
}
