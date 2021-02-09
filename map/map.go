package main

import "fmt"

func print_map(data map[string]int)  {
	fmt.Printf("\n===============\n")
	for k, v := range data {
		fmt.Printf("%s => %d\n", k, v)
	}
}

func main()  {
	var ndata map[string]int
	// ndata["pen"] = 5 // error
	v, ok := ndata["pen"]
	if (ok) {
		fmt.Printf("pen : %d\n", v)
	} else {
		fmt.Printf("pen is not in ndata\n")
	}
	fmt.Println(ndata)

	data := make(map[string]int)

	data["pen"] = 5
	data["book"] = 3
	d_ref := data
	fmt.Println(data)

	print_map(data)
	print_map(d_ref)
	delete(data, "pen")
	d_ref["bag"] = 20
	print_map(data)
	print_map(d_ref) // same as data
}