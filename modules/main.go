package main

import (
	"errors"
	"fmt"
	"log"
	"modules/dao"
)


func assert(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty string")
	}
	return s, nil
}

func main() {
	dao.DoSomething()
	fmt.Println(dao.DAO_INT)
	// fmt.Println(dao.dao_xx) // error
	_, err := assert("")
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
}