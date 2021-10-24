package main

import "fmt"

func main() {
	var nFunc, qHora int
	var vHora, total float64
	fmt.Scanf("%d", &nFunc)
	fmt.Scanf("%d", &qHora)
	fmt.Scanf("%f", &vHora)
	total = float64(qHora) * vHora
	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", nFunc, total)
}
