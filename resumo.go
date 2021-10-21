package main

//import "fmt" //-> É possivel importar de duas formas: individualmente ou em conjunto

import (
	//f "fmt" //-> alias para chamar o pacote com um "apelido"
	"fmt"
)

func numeros(n1 float32, n2 float32) {
	soma := n1 + n2 // := possibilite a instanciação de varáveis sem usar o "var" no início
	fmt.Println(soma)
}

func ponteiros() {
	var a int = 7
	var b *int
	b = &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b)
	*b++
	fmt.Println(a)
}

func variaveis() {
	// Podemos instanciar várias variáveis dessa forma
	var (
		peso   float64 = 75.0
		idade  int     = 18
		altura float64 = 1.71
	)
	fmt.Println("Olá")
	fmt.Println("Sua idade é igual a: ", idade)
	fmt.Printf("Vc apresentou um peso de %f quilos e uma altura de %f metros \n", peso, altura)
}

func entrada() {
	idade, peso, altura := 0, 0.0, 0.0
	fmt.Print("Informe sua idade: ")
	fmt.Scanf("%v", &idade)
	fmt.Print("Informe o peso: ")
	fmt.Scanf("%v", &peso)
	fmt.Print("Informe a altura: ")
	fmt.Scanf("%v", &altura)
	fmt.Printf("Olá!\nSua idade é igual a: %v \n", idade)
	fmt.Printf("Vc apresentou um peso de %g quilos e uma altura de %g metros", peso, altura)
}

func constantes() {
	var num1 int = 10
	const num2 int = 10
	num1 = 20
	// num2 = 20 // consts não permitem que seja atribuido outros valores a variável

	fmt.Println(num1)
	fmt.Println(num2)
}

func condicionais() {
	idade := -1

	if idade <= 18 {
		fmt.Println("Menor de idade")
	} else {
		fmt.Println("Maior de idade")
	}

	switch {
	case idade > 1 && idade <= 12:
		fmt.Println("Pré adolescente")
	case idade > 12 && idade <= 18:
		fmt.Println("Adolescente")
	case idade > 18:
		fmt.Println("Adulto")
	default:
		fmt.Println("Feto kkk")
	}
}

func repeticao() {
	// for
	fmt.Println("Número impares:")
	for i := 100; i > 0; i-- {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	fmt.Print("------------------------------")

	// while
	j := 0
	for j <= 10 {
		fmt.Println(j)
		j++
	}
}

func arrays() {
	var frutas [5]string
	index := 1
	for i := 0; i < len(frutas); i++ {
		fmt.Printf("Informe o nome da %dº fruta: ", index)
		fmt.Scanf("%s", &frutas[i])
		index += 1
	}
	fmt.Println(frutas)
}

func main() {
	//numeros(2.3, 1.2)
	//ponteiros()
	//variaveis()
	//entrada()
	//constantes()
	//condicionais()
	//repeticao()
	arrays()

}
