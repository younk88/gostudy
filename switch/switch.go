package main

import "fmt"

func main()  {
	var it interface{}
	x := 1
	it = x
	switch t := it.(type) {
	case string:
		fmt.Println("x is string")
	case int, uint:
		fmt.Println("x is int")
	default:
		fmt.Println("unkown ", t)
	}

	switch x {
	case 1:
		fallthrough
	case 2:
		fmt.Println("x is below 3")
	case 3:
		fmt.Println("x is equal 3")
		
	}
}