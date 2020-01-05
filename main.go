package main

import (
	"fmt"
	"math"
	"time"
)

func add(x, y, z int) int {
	return x + y + z
}
func main() {

	fmt.Println("Hello, May")
	fmt.Println("The time is", time.Now())
	fmt.Println("totals", math.Max(7, 10))
	a := add(44, 12, 33)
	fmt.Println(a)

}
