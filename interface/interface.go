package main

import (
	"fmt"
	"interface/sub"
	"strconv"
)

func print_obj(i sub.IInfo) {
	fmt.Printf("%v, %T\n", i, i)
}

type ITest interface{
	run()
}

func print_type(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("is int")
	case string:
		fmt.Println("is string")
	default:
		fmt.Printf("unkown type %T\n", v)
	}
}

type IPAddr [4]byte

func (ip IPAddr)String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	info := sub.CreateInfo("info1")
	info_r := info

	print_obj(info)
	print_obj(info_r)

	fmt.Println(info.Text())
	fmt.Println(info_r.Text())

	info.SetText("I'm here")
	fmt.Println(info.Text())
	fmt.Println(info_r.Text())

	t, ok := info.(sub.IInfo)
	fmt.Printf("info is sub.IInfo %v, %v\n", t, ok)

	x, xo := info.(ITest)
	fmt.Printf("info is ITest %v, %v\n", x, xo)

	print_type(1)
	print_type("1")
	print_type(true)
	print_type(1.2)
	var n uint = 16
	print_type(n)

	var locIp = IPAddr{127, 0, 0, 1}
	fmt.Println(locIp)
	fmt.Println(locIp.String())
	fmt.Printf("%v\n", locIp)
}