package main

import (
	"fmt"
	"math"
)

func main() {
	var r, vol float64
	fmt.Scanf("%f", &r)
	vol = (4 / 3.0) * 3.14159 * math.Pow(r, 3.00)
	fmt.Printf("VOLUME = %.3f\n", vol)
}
