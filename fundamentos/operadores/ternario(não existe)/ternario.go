package main

import "fmt"

func obterResultadoProva(nota float64) string {
	// return nota >=60 ? "aprovado" : "reprovado"   Obs: não existe ternario em go
	if nota >= 60 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultadoProva(59))
	fmt.Println(obterResultadoProva(69))
}
