package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func main() {
	fmt.Println("rand num is: ",rand.Intn(10))
	fmt.Println("time now is ",time.Now())
	fmt.Println("now you have %g problems",math.Sqrt(7))
	fmt.Println("PI",math.Pi)
	fmt.Println("2+3=",add(2,3))
	a,b := swap("hello","world")
	fmt.Println("swap",a,b)
}
//
func add(x int,y int) int {
	return x+y
}

//
func swap(a ,b string)(string,string) {
	return b ,a
}
