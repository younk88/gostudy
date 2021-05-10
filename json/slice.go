package main

import (
	"encoding/json"	
	"fmt"
)


func main() {
	s := "[[1, [1, 3]],[2, [27, 29]], [4, [42, 48]]]"
	data := make([][]interface{}, 0)
	err := json.Unmarshal([]byte(s), &data)
	fmt.Println(err)
	fmt.Println(data)
	for _, item := range data {
		g := item[0].(float64)
		v := item[1].([]interface{})
		fmt.Println(g, v)
	}

	var x uint64 = 0
	fmt.Printf("%d != 0 => %v\n", x, x != 0)
}