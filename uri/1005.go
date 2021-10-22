package main

import "fmt"

func main() {
	var notaA, notaB, media float64
	fmt.Scanf("%f", &notaA)
	fmt.Scanf("%f", &notaB)
	media = ((notaA * 3.5) + (notaB * 7.5)) / 11
	fmt.Printf("MEDIA = %.5f\n", media)
}
