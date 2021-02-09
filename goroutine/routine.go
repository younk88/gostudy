package main

import "fmt"
import "time"

var _c int = 0

func generate_int(c chan int) {
	for i := 0; i < 20; i++ {
		c <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(c)
}

func print_int(c chan int, st chan int)  {
	for {
		v, ok := <- c
		if ok {
			fmt.Printf("got value: %d %v\n", v, ok)
		} else {
			break
		}
	}
	//st <- 1
}

func greeting(c1, c2 chan string) {
	x := ""
	for {
		select {
		case x = <- c1:
			fmt.Printf("Hello %s\n", x)
		case x = <- c1:
			fmt.Printf("Bye %s\n", x)
		case <- c2:
			fmt.Println("end greeting")
			break
		}
	}
}

func walk(data *[]int, count int, ch chan int) {
	for i := 0; i < count; i++ {
		ch <- (*data)[i]
	}
	close(ch)
}

func match(data *[]int, count int, ch chan int, cs chan bool) {
	var v int
	var ok bool
	i := 0
	for {
		v, ok = <- ch
		if ok && i < count {
			if v != (*data)[i] {
				cs <- false
				break
			}
			i++
		} else {
			cs <- true
			break
		}
	}
}

var _sc int = 0

func main()  {
	status := make(chan int)
	//ct := make(chan int, 5)
	//go generate_int(ct)
	//go print_int(ct, status)

	go (func (st chan int) {
		for i := 0; i < 20; i++ {
			_sc++
			fmt.Printf("task1 %d\n", _sc)
		}
		st <- 1
	})(status)

	go (func (st chan int) {
		for i := 0; i < 20; i++ {
			_sc++
			fmt.Printf("task2 %d\n", _sc)
		}
		st <- 1
	})(status)

	gc := make(chan string)
	gs := make(chan string)
	go greeting(gc, gs)

	gc <- "Tom"
	gc <- "Hellen"
	gc <- "Lily"
	gc <- "Miya"
	gs <- "end"

	step := 0
	for {
		_ = <- status
		step += 1
		if step >= 2 {
			break
		}
	}
	fmt.Printf("end\n")
	
	d1 := []int{1, 3, 5, 7, 9, 11}
	d2 := []int{1, 3, 5, 6, 9, 11}
	vc := make(chan int)
	rc := make(chan bool)
	go walk(&d1, len(d1), vc)
	go match(&d2, len(d2), vc, rc)
	is_same := <- rc
	fmt.Printf("d1 compare with d2 result: %v\n", is_same)
}