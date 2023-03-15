package main

import "fmt"

func main() {

	var ola = "Olá"
	var bem = "tudo bem com você?"
	var name string

	fmt.Println("Qual o seu nome?")
	fmt.Scanf("%s", &name)
	fmt.Println(ola, name, bem)
}
