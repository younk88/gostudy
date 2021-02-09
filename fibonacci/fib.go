package main

import "fmt"

func fib1(n int) int {
	fmt.Printf("call fib1 %d\n", n)
	if n < 2 {
		return n;
	}
	return fib1(n - 1) + fib1(n - 2)
}

func fib2_inner(n int) (int, int) {
	fmt.Printf("call fib2_inner %d\n", n)
	if n < 2 {
		return 0, n;
	}
	x, y := fib2_inner(n - 1)
	return y, x + y
}

func fib2(n int) int {
	_, r := fib2_inner(n)
	return r
}

func main()  {
	fmt.Printf("fib1(8) == %d\n", fib1(8))
	
	fmt.Printf("fib2(8) == %d\n", fib2(8))
}