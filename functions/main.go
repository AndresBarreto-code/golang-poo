package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	y := func() int {
		return x * x
	}()
	fmt.Println(x)
	fmt.Println(y)
	c := make(chan int)
	go func() {
		fmt.Println("Starting function...")
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("End function.")
		c <- 1
	}()
	<-c
	fmt.Println("End main.")
}
