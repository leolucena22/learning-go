package main

import "fmt"

func main() {
	var a, b, c, aTriangulo, aCirculo, aTrapezio, aQuadrado, aRetangulo float64
	fmt.Scanf("%f %f %f", &a, &b, &c)
	aTriangulo = (a * c) / 2
	aCirculo = 3.14159 * (c * c)
	aTrapezio = ((a + b) * c) / 2
	aQuadrado = b * b
	aRetangulo = a * b
	fmt.Printf("TRIANGULO: %.3f\n", aTriangulo)
	fmt.Printf("CIRCULO: %.3f\n", aCirculo)
	fmt.Printf("TRAPEZIO: %.3f\n", aTrapezio)
	fmt.Printf("QUADRADO: %.3f\n", aQuadrado)
	fmt.Printf("RETANGULO: %.3f\n", aRetangulo)
}
