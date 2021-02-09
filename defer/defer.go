package main

import "fmt"

func dfunc(i int) {
	fmt.Printf("defer call %d\n", i)
}

func wrapf() {
	for i := 100; i < 105; i++ {
		defer dfunc(i)
	}
	fmt.Println("execute wrap task")
	fmt.Println("wrap end")
}

func aint(c int) int {
	defer (func ()  {
		c++
		fmt.Println("aint defer ", c)
	})()
	return c
}

func main()  {
	var c int = 0
	defer (func (p int)  {
		fmt.Printf("first defer p=%d c=%d\n", p, c)
	})(c)
	for i := 0; i < 10; i++ {
		c++
		defer dfunc(i)
	}
	wrapf()
	for i := 20; i < 30; i++ {
		defer dfunc(i)
	}

	x := aint(3)
	fmt.Println("aint return: ", x)

	fmt.Println("execute task")
	fmt.Println("end")
}