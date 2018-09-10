package main

import (
	"fmt"
)

func compas(trab1, trab2 bool) (bool, bool, bool) {
	compraTv50 := trab1 && trab2
	compraTv32 := trab1 != trab2
	compraSorvete := trab1 || trab2

	return compraTv50, compraTv32, compraSorvete
}

func main() {
	traba1, trab2 := false, false
	tv50, tv32, tomaSorve := compas(traba1, trab2)

	fmt.Printf("Tv50: %v tv32: %v tomaSoverve: %v MaisSaudável:%v", tv50, tv32, tomaSorve, !tomaSorve)
}
