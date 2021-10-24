package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanf("%f %f %f", &a, &b, &c)
	maiorAB := (a + b + math.Abs(a-b)) / 2

	if maiorAB > c {
		fmt.Printf("%.f eh o maior\n", maiorAB)
	} else {
		fmt.Printf("%.f eh o maior\n", c)
	}
}
