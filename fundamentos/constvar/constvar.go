package main

import (
	"fmt"
	m "math"
)

func main() {
	//float64 é o default mas poderia ser também float32
	const PI float64 = 3.1416
	var raio = 3.2

	//forma reduzida de criar variável, declara e atribui ao mesmo tempo
	area := PI * m.Pow(raio, 2)
	fmt.Println("A area da circunferencia é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 1
		d = 2
	)

	println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)

}
