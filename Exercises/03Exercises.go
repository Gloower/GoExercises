package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	txt := bufio.NewReader(os.Stdin)
	fmt.Println("Digite o seu texto aqui: ")
	text, _ := txt.ReadString('\n')
	inverted := invert(text)

	fmt.Println("Seu texto invertido: " + inverted)
}

func invert(text string) string {
	txt := []rune(text)
	ret := []rune{}

	for i := len(txt) - 2; i >= 0; i-- {
		ret = append(ret, txt[i])
	}

	return string(ret)
}
