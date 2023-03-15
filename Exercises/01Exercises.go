package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var ola = "Olá"
	var bem = "tudo bem com você?"

	sc := bufio.NewReader(os.Stdin)

	fmt.Println("Qual o seu nome?")
	name, _ := sc.ReadString('\n')
	fmt.Println(ola, name, bem)
}
