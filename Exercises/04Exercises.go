package main

import (
	"fmt"
)

func main() {

	var num1, num2, num3 int
	var maiorNum = 0

	fmt.Println("Insira o primeiro número:")
	fmt.Scanf("%s", &num1)

	fmt.Println("Insira o segundo número:")
	fmt.Scanf("%s", &num2)

	fmt.Println("Insira o terceiro número:")
	fmt.Scanf("%s", &num3)

	if num1 > maiorNum {
		maiorNum = num1

	}

	if num2 > maiorNum {
		maiorNum = num2

	}

	if num3 > maiorNum {
		maiorNum = num3

	}

}
