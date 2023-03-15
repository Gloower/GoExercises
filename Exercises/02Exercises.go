package main

import (
	"fmt"
)

func main() {

	var vetor [11]int

	for i := 1; i < 11; i++ {
		vetor[i] = (i)
	}

	for i := 1; i < len(vetor); i++ {
		fmt.Println(vetor[i])
	}
}
