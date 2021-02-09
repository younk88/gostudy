package main

import "fmt"

func multiple(m int)  {
	for i := 1; i <= m; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i * j);
		}
		fmt.Println("")
	}
}

func plus(m int)  {
	for i := 1; i <= m; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d + %d = %d\t", j, i, i + j)
		}
		fmt.Println("")
	}
}

func plus_array(m int) [][]string  {
	ret := [][]string{}
	row := []string{}
	for i := 1; i <= m; i++ {
		row = []string{}
		for j := 1; j <= i; j++ {
			row = append(row, fmt.Sprintf("%d + %d = %d\t", j, i, i + j))
		}
		ret = append(ret, row)
	}
	return ret
}

func main()  {
	multiple(9)
	plus(9)
	arr := plus_array(6)
	fmt.Println(arr)
}