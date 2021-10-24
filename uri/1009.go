package main

import "fmt"

func main() {
	var nome string
	var salario, tVendas, total float64
	fmt.Scanf("%s", &nome)
	fmt.Scanf("%f", &salario)
	fmt.Scanf("%f", &tVendas)
	total = salario + (0.15 * tVendas)
	fmt.Printf("TOTAL = R$ %.2f\n", total)
}
