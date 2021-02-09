package main

import "fmt"

type cb func(int)

func do_something(num int, f cb)  {
	f(num * 2)
}

func main()  {
	do_something(3, func (n int) {
		fmt.Println(n);
	})
}