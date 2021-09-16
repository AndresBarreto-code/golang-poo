package main

import (
	"fmt"
)

func sum(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func printName(values ...string) {
	for _, v := range values {
		fmt.Println(v)
	}
}

func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7))
	printName("Name 0")
	printName("Name 0", "Name 1")
	d, t, q := getValues(5)
	fmt.Println(d, t, q)
}
