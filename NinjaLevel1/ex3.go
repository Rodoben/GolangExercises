package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	fmt.Println(x, y, z)
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}
