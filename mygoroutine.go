package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}

	time.Sleep(time.Millisecond)
}

func Add(x, y int) {
	z := x + y
	fmt.Println("z=", z)
}
