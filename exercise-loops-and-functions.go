package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	return x
}
func main() {
	fmt.Println(
		Sqrt(16),
		math.Sqrt(16),
	)
}
