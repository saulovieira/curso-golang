package main

import (
	"fmt"
	"time"
)

func loopInfinito() {
	count := 0
	for {
		fmt.Print(count)
		time.Sleep(time.Second)
		if count == 5 {
			fmt.Print("\n")
			break
		}
		count++
	}
}

func loopNormal() {
	for i := 0; i <= 5; i++ {
		fmt.Print(i)
		time.Sleep(time.Second)
	}
}

func main() {
	loopInfinito()
	loopNormal()
}
