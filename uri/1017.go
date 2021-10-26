package main

import "fmt"

func main() {
	var tempo, velocidade, distancia int
	var litros float64
	fmt.Scanf("%d", &tempo)
	fmt.Scanf("%d", &velocidade)
	distancia = tempo * velocidade
	litros = float64(distancia) / 12
	fmt.Printf("%.3f\n", litros)
}
