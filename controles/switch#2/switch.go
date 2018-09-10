package main

import (
	"fmt"
	"time"
)

func printHour() {
	t := time.Now()
	switch {
	case t.Hour() < 23:
		fmt.Printf("%v Boa noite\n", t.Hour())
	case t.Hour() < 22:
		fmt.Printf("%v Boa noite\n", t.Hour())
	case t.Hour() < 18:
		fmt.Printf("%v Boa noite\n", t.Hour())
	case t.Hour() < 12:
		fmt.Printf("%v Boa tarde\n", t.Hour())
	case t.Hour() < 6:
		fmt.Printf("%v Bom dia\n", t.Hour())
	default:
		fmt.Printf("%v Indefinido\n", t)
	}
}

func main() {
	printHour()

}
