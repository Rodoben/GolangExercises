package main

import "fmt"

func main() {
	var a int = 2
	fmt.Printf("%d %b %#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d %b %#x\n", b, b, b)
}
