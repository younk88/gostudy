package main

import "fmt"

func got_err() {
	err := recover()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("good")
	}
}

func main()  {
	var e int
	e = 10
	fmt.Println(e / 2)
	e -= 10
	defer got_err()
	fmt.Println(2 / e)
	// got_err() // not work
}