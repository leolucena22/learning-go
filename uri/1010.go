package main

import "fmt"

func main() {
	var codPeca1, codPeca2, qPeca1, qPeca2 int
	var vPeca1, vPeca2, total float64
	fmt.Scanf("%d %d %f", &codPeca1, &qPeca1, &vPeca1)
	fmt.Scanf("%d %d %f", &codPeca2, &qPeca2, &vPeca2)
	total = (float64(qPeca1) * vPeca1) + (float64(qPeca2) * vPeca2)
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
