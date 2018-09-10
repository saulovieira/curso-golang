package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("String :", "Banana" == "Banana")
	fmt.Println("!= :", 3 != 2)
	fmt.Println("< :", 3 < 2)
	fmt.Println("> :", 3 > 2)
	fmt.Println("<= :", 3 <= 2)
	fmt.Println(">= :", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("datas :", d1 == d2)
	fmt.Println("datas :", d1.Equal(d2))

	type pessoa struct {
		Nome string
	}

	p1 := pessoa{"Saulo"}
	p2 := pessoa{"João"}
	fmt.Println("struct :", p1 == p2)

	p2 = pessoa{"Saulo"}
	fmt.Println("struct :", p1 == p2)

}
