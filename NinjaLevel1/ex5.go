package main

import "fmt"

type hotdog int

var y int

func main() {

	var x hotdog
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)

	y = int(x)
	fmt.Printf("Y:%v\n", y)
}
