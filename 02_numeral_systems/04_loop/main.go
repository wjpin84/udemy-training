package main

import "fmt"

func main() {
	for x := 60; x < 127; x++ {
		fmt.Printf("%d\t %b\t%x\t%q\n", x, x, x, x)
	}
}
