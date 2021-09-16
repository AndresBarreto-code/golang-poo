package main

import (
	"fmt"

	hellomod "github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
)

func main() {
	fmt.Println("Hello World!")
	hellomod.SayHello()
	hellomod2.SayHello("Andres ")
}
