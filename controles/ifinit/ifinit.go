package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatoreo() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatoreo(); i > 5 {
		fmt.Printf("%v : Ganhou!!!", i)
	} else {
		fmt.Printf("%v : Perdeu!", i)
	}
}
