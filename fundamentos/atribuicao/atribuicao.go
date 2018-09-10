package main

import (
	"fmt"
)

func main() {
	var b byte = 3
	i := b
	fmt.Println(i)
	i += 3
	fmt.Println(i)
	i -= 3
	fmt.Println(i)
	i /= 3
	fmt.Println(i)
	i *= 3
	fmt.Println(i)
	i %= 2
	fmt.Println(i)

	x, y, z := 1, 2, 3

	fmt.Println(x, y, z)

	x, y, z = y, z, x

	fmt.Println(x, y, z)
}
