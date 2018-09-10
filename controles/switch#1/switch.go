package main

import (
	"fmt"
)

func obterConceito(nota float64) string {
	//ainda tenho que entender o uso do fallthrough
	switch nota {
	case 10, 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3, 2:
		return "D"
	case 1, 0:
		return "E"
	default:
		return "Indefinido"
	}
}

func main() {
	fmt.Println(obterConceito(10))
	fmt.Println(obterConceito(9))
	fmt.Println(obterConceito(8))
	fmt.Println(obterConceito(7))
	fmt.Println(obterConceito(6))
	fmt.Println(obterConceito(5))
	fmt.Println(obterConceito(4))
	fmt.Println(obterConceito(3))
	fmt.Println(obterConceito(2))
	fmt.Println(obterConceito(1))
	fmt.Println(obterConceito(0))
	fmt.Println(obterConceito(11))
}
