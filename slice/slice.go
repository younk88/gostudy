package main

import "fmt"

func print_slice_info(s []int)  {
	fmt.Printf("slice len: %d, cap: %d slice: %v\n", len(s), cap(s), s)
}

func print_arr_info(s [2]int)  {
	sum := 0
	for _, n := range s {
		sum += n
	}
	fmt.Printf("arr len: %d, cap: %d sum: %d, slice: %v\n", len(s), cap(s), sum, s)
}

func change_slice(s []int)  {
	s[0] = 0
}

func change_arr(s [2]int)  {
	s[0] = 0
}

func change_arr_ptr(s *[2]int) {
	s[0] = 0
}

func append_element(s *[]int) {
	*s = append(*s, 1)
}

func main()  {
	data := []int {1,3,5,7,9,11,13}
	print_slice_info(data)

	s1 := data[1:3]
	s2 := data[1:]
	s3 := data[:3]
	s4 := data[3:]
	print_slice_info(s1)
	print_slice_info(s2)
	print_slice_info(s3)
	print_slice_info(s4)

	s4 = append(s4, 8)
	print_slice_info(s4)
	s4 = append(s4, 8)
	s4 = append(s4, 8)
	s4 = append(s4, 8)
	s4 = append(s4, 8)
	print_slice_info(s4)
	// 1. cap的扩容是double 原cap？

	s5 := []int{1, 2}
	print_slice_info(s5)
	s5 = append(s5, 3, 4, 5, 4, 7, 6, 2)
	print_slice_info(s5)
	// 回答1：是的，每次append如果容量不足时，扩容策略是每次double调用append时的cap，多次执行该扩容方案，直到能满足容纳添加的元素

	change_slice(s5)
	print_slice_info(s5)

	c := make([]int, len(s5), 9)
	copy(c, s5)
	print_slice_info(c)

	arr := [2]int{1,2}
	print_arr_info(arr)
	change_arr(arr)
	print_arr_info(arr) // arr not change
	// arr = append(arr, 3) // error
	arr_copy := arr
	arr_ptr := &arr
	fmt.Println(arr_copy)
	fmt.Println(arr_ptr)
	fmt.Println(*arr_ptr)
	arr_copy[1] = 3 // only change copy, not change arr
	change_arr_ptr(arr_ptr) // arr changed
	print_arr_info(arr)
	fmt.Println(arr_copy)
	fmt.Println(arr_ptr)

	d := [6]int{}
	ds1 := d[1: 5]
	ds2 := d[:4]
	print_slice_info(ds1)
	print_slice_info(ds2)
	ds1[0] = 1
	d[2] = 2
	ds2[3] = 3
	ds1 = append(ds1, 4)
	ds1[3] = 14
	fmt.Println(d)
	print_slice_info(ds1)
	ds1 = append(ds1, 6) // 切片超出源数组长度后，完全指向新的数组，不在和源数组有关联
	ds1[3] = 16
	ds11 := ds1[3:]
	ds2[0] = 20
	ds11[0] = 88
	fmt.Println(d)
	print_slice_info(ds1)
	print_slice_info(ds2)
	print_slice_info(ds11)

	var ns []int
	print_slice_info(ns)
	ns = append(ns, 2)
	print_slice_info(ns)

	as := []int{}
	append_element(&as)
	append_element(&as)
	print_slice_info(as)
}