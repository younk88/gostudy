package main

import (
	"fmt"

	"rsc.io/quote"
)


func init()  {
	fmt.Println("init done")
}

func main() {
	fmt.Println("enter main")
	var a string = "ok"
	var b, c int = 2, 3
	var d = true
	e, f := 5, 9
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(&a)

	var r = &a
	fmt.Println(a, *r)
	a = "fail"
	fmt.Println(a, *r)
	*r = "of"
	fmt.Println(a, *r)

	const (
		A = iota
		B
		C
		D = "ha"
		E
		F = 99
		G
		H = iota
		I
	)
	fmt.Println(A, B, C, D, E, F, G, H, I)

	const (
		X = 3 * iota
		Y
		Z
	)
	fmt.Println(X, Y, Z)

	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [5]int{1,3,5}
	for i, n := range numbers {
		fmt.Println(i, n)
	}

	fmt.Println(quote.Go())
}