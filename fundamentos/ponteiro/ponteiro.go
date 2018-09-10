package main

import (
	"fmt"
)

func main() {
	i := 1

	fmt.Println("variavel i: ", i)
	//define que p é um ponteiro que aponta para endereço do tamnaho int e ele nesse momentos é nulo
	var p *int
	
	p = &i //paga o endereço da variável i
	
	*p++ //inclementa o valor de i já que o ponteiro aponta para i
	
	fmt.Println("variavel i: ", i)
	//Go não tem aritimetica de ponteiros
	// p++
	
	fmt.Printf("valor i: %v, end mem i: %v \nvalor p: %v, end mem p: %v \n", i, &i, *p, p)
	fmt.Println("p e i apontam mesmo endereço de memoria: ", p == &i)

}
