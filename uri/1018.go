package main

import "fmt"

func main() {
	var valor, troco, cem, cinquenta, vinte, dez, cinco, dois int
	fmt.Scanf("%d", &valor)
	troco = valor
	cem = troco / 100
	troco = troco - (cem * 100)
	cinquenta = troco / 50
	troco = troco - (cinquenta * 50)
	vinte = troco / 20
	troco = troco - (vinte * 20)
	dez = troco / 10
	troco = troco - (dez * 10)
	cinco = troco / 5
	troco = troco - (cinco * 5)
	dois = troco / 2
	troco = troco - (dois * 2)
	fmt.Printf("%d\n", valor)
	fmt.Printf("%d nota(s) de R$ 100,00\n", cem)
	fmt.Printf("%d notas(s) de R$ 50,00\n", cinquenta)
	fmt.Printf("%d notas(s) de R$ 20,00\n", vinte)
	fmt.Printf("%d notas(s) de R$ 10,00\n", dez)
	fmt.Printf("%d notas(s) de R$ 5,00\n", cinco)
	fmt.Printf("%d notas(s) de R$ 2,00\n", dois)
	fmt.Printf("%d notas(s) de R$ 1,00\n", troco)
}
