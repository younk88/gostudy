package main

import "fmt"

func xx() int {
	x := 1 + 2
	return x
}

func main() {
	x := xx()
	fmt.Println(x)
}
