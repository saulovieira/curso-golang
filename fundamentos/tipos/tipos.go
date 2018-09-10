package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literar inteiro é", reflect.TypeOf(32000))

	var b byte = 255
	fmt.Println("O byte  é", reflect.TypeOf(b))

	i1 := math.MaxInt64

	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	//o tipo rune foi omitido devidor o compilador já inferir pelo uso das aspas simples no valors
	var i2 = 'a'

	fmt.Println("O tipo de i2 é", reflect.TypeOf(i2))
	fmt.Println("O valor de i2 é", 'a', "valor inferido", i2)

	var x float32 = 49.99
	var y = 49.69 //o tipo float64 foi omitido devidor o compilador já inferir por ser o default para ponto flutuante

	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O valor de x é", x)
	fmt.Println("O tipo de y é", reflect.TypeOf(y))
	fmt.Println("O valor de y é", y)
	fmt.Println("O cast de float64 para float32:", reflect.TypeOf(float32(y)))

	bo := true

	fmt.Println("O tipo de b é", reflect.TypeOf(bo))
	fmt.Printf("O valor de b é %t e sua negação %t \n", bo, !bo)

	s1 := "Isto é uma String"

	fmt.Println("O tipo de s1 é", reflect.TypeOf(s1))
	fmt.Println("O valor de s1 é", s1+"!!(concatenado)")
	fmt.Println("O tamanho da string é", len(s1))


	s2 := `Isto é uma String 
	com multiplas 
	linhas`

	fmt.Println("O tipo de s1 é", reflect.TypeOf(s2))
	fmt.Println("O valor de s1 é", s2+"!!(concatenado)")
	fmt.Println("O tamanho da string é", len(s2))

}
