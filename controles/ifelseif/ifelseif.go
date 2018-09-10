package main

import "fmt"

func obterConceitoNota(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 8 && nota <= 9 {
		return "B"
	} else if nota >= 5 && nota < 8 {
		return "C"
	} else if nota >= 3 && nota <= 5 {
		return "D"
	} else if nota >= 1 && nota <= 3 {
		return "E"
	} else {
		return "F"
	}
}

func main() {
	fmt.Println(obterConceitoNota(0))
	fmt.Println(obterConceitoNota(1))
	fmt.Println(obterConceitoNota(2))
	fmt.Println(obterConceitoNota(3))
	fmt.Println(obterConceitoNota(4))
	fmt.Println(obterConceitoNota(5))
	fmt.Println(obterConceitoNota(6))
	fmt.Println(obterConceitoNota(7))
	fmt.Println(obterConceitoNota(8))
	fmt.Println(obterConceitoNota(9))
	fmt.Println(obterConceitoNota(10))
}
