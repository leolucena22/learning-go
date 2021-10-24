package main

import "fmt"

func main() {
	var dPercorrida int
	var tCombGasto, cMedio float64
	fmt.Scanf("%d", &dPercorrida)
	fmt.Scanf("%f", &tCombGasto)
	cMedio = float64(dPercorrida) / tCombGasto
	fmt.Printf("%.3f km/l\n", cMedio)
}
