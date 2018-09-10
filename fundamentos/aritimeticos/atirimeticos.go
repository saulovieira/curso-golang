package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("soma => ", a+b)
	fmt.Println("Subtracao => ", a-b)
	fmt.Println("multiplicacao => ", a*b)
	fmt.Println("divicao => ", a/b)
	fmt.Println("modulo => ", a%b)

	//bitwise
	fmt.Println("E => ", a&b)   // (3)11 & (2)10 = (2)10
	fmt.Println("OU => ", a|b)  // (3)11 & (2)10 = (3)11
	fmt.Println("XOR => ", a^b) // (3)11 & (2)10 = (1)01 | quando dois uns então o resultado da soma é 0

	fmt.Println("Maior => ", math.Max(float64(a), float64(b)))

	c := 2.0
	d := 3.0

	fmt.Println("Menor => ", math.Min(c, d))

	fmt.Println("Exponenciação => ", math.Pow(c, d))

	fmt.Println("Exponenciação => ", math.Sqrt(math.Pow(d, c)))

}
