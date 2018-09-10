package main

import "fmt"

func obterResultadoProva(nota float64) {
	notaMinima := 60.0
	fmt.Println("Nota minima", notaMinima)
	if nota >= notaMinima {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	obterResultadoProva(59)
	obterResultadoProva(69)
}
