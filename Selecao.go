package Exercicios

import "fmt"

func Sel_1() {
	var a int
	var b int
	fmt.Printf("Insere dois numeros inteiros e mostra o maior \n")
	fmt.Printf("Insira o primeiro numero")
	fmt.Scanf("%d", &a)
	fmt.Printf("Insira o segundo numero")
	fmt.Scanf("%d", &b)

	fmt.Print("O maior numero é: ")
	if a > b {
		fmt.Printf("%d \n", a)
	} else if a < b {
		fmt.Printf("%d \n", b)
	} else if a == b {
		fmt.Println("...\nNENHUM\nSão iguais")
	}
	fmt.Println("FIM")
}

func Sel_2() {
	var a int
	var b int
	var c int
	fmt.Printf("Insere tres numeros inteiros e mostra o menor \n")
	fmt.Printf("Insira o primeiro numero")
	fmt.Scanf("%d", &a)
	fmt.Printf("Insira o segundo numero")
	fmt.Scanf("%d", &b)
	fmt.Printf("Insira o terceiro numero")
	fmt.Scanf("%d", &c)

	fmt.Print("O menor numero é: ")
	if a < b && a < c {
		fmt.Printf("%d \n", a)
	} else if b < a && b < c {
		fmt.Printf("%d \n", b)
	} else if c < a && c < b {
		fmt.Printf("%d \n", c)
	} else if a == b {
		fmt.Println("...\nNENHUM\nPelo menos dois são iguais\nTHERE CAN BE ONLY ONE")
	}
	fmt.Println("FIM")
}

func Sel_3() {
	var a int
	fmt.Printf("Insere um numero e mostra se ele é par ou impar \n")
	fmt.Printf("Insira o numero")
	fmt.Scanf("%d", &a)
	if a%2 == 1 {
		fmt.Println("IMPAR")
	} else {
		println("PAR")
	}
}

func Sel_4() {

}
